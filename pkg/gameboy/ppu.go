package gameboy

import "github.com/elmarsan/gbgo/pkg/bit"

// ppu represents game boy pixel processing unit.
type ppu struct {
	// videoBuffer contains LCD pixels with palette index values.
	videoBuffer [ScreenWidth * ScreenHeight]uint8

	// vBlankClockCycles is used to track vBlank mode.
	vBlankClockCycles int

	// statusClockCycles is used to track frame clock cycles.
	statusClockCycles int

	// statusMode holds rendering mode
	// MODE 0 HBLANK
	// MODE 1 VBLANK
	// MODE 2 OAM SCAN
	// MODE 3 DRAWING PIXELS
	statusMode int

	// scanline keeps current horizontal line.
	// It might be about to be drawn, being drawn, or just been drawn.
	scanline uint8

	// vblankline keeps current vblankline. (0-10)
	vblankLine int

	// memoryBus represents memory memoryBus used by Gameboy.
	memoryBus *memoryBus

	// InterruptBus represents the interruption system of game boy.
	interruptBus *interruptBus
}

const (
	lcdc = 0xff40 // LCD control
	stat = 0xff41 // LCD status
	scy  = 0xff42 // Viewport Y position
	scx  = 0xff43 // Viewport X position
	wy   = 0xff4a // Window Y position
	wx   = 0xff4b // Window X position
	ly   = 0xff44 // LCD Y coordinate
	lyc  = 0xff45 // LY compare
	bgp  = 0xff47 // Background palette
	obp0 = 0xff48 // Object palette 0
	obp1 = 0xff49 // Object palette 0

	ScreenWidth  = 160 // Game boy screen width
	ScreenHeight = 144 // Game boy screen height
)

const (
	hblank   = iota // MODE 0 HBLANK
	vblank          // MODE 1 VBLANK
	oamScan         // MODE 2 OAM SCAN
	drawMode        // MODE 3 DRAWING PIXELS
)

// newPpu creates and returns a new ppu instance.
// Ppu members are set defaults of DMG boot sequence.
func newPpu(memoryBus *memoryBus, irBus *interruptBus) *ppu {
	videoBuf := [ScreenWidth * ScreenHeight]uint8{}

	// Set all pixels to palette index 0
	for i := 0; i < len(videoBuf); i++ {
		videoBuf[i] = 0
	}

	return &ppu{
		videoBuffer:       videoBuf,
		vBlankClockCycles: 0,
		statusClockCycles: 0,
		// ppu starts in VBLANK mode
		statusMode: 1,
		// First scanline line must be 144
		scanline:     144,
		vblankLine:   0,
		memoryBus:    memoryBus,
		interruptBus: irBus,
	}
}

// Ticks emulates ppu ticks.
func (ppu *ppu) tick(clockCycles int) {
	lcdc := ppu.memoryBus.read(lcdc)

	if !bit.IsSet(lcdc, 7) {
		return
	}

	ppu.statusClockCycles += clockCycles

	switch ppu.statusMode {

	case hblank:
		if ppu.statusClockCycles >= 204 {
			ppu.statusClockCycles -= 204
			ppu.updateStatusMode(oamScan)
			ppu.scanline++
			ppu.memoryBus.write(ly, ppu.scanline)
			ppu.compareLY()

			// Vblank mode starts
			if ppu.scanline == 144 {
				ppu.updateStatusMode(vblank)

				stat := ppu.memoryBus.read(stat)
				if bit.IsSet(stat, 4) {
					ppu.interruptBus.request(lcdStatInterrupt)
				}
			} else {
				stat := ppu.memoryBus.read(stat)
				if bit.IsSet(stat, 5) {
					ppu.interruptBus.request(lcdStatInterrupt)
				}
			}
		}

	case vblank:
		ppu.vBlankClockCycles += clockCycles

		if ppu.vBlankClockCycles >= 456 {
			ppu.vBlankClockCycles -= 456
			ppu.vblankLine++

			if ppu.vblankLine <= 9 {
				ppu.scanline++
				ppu.memoryBus.write(ly, ppu.scanline)
				ppu.compareLY()
			}
		}

		if (ppu.statusClockCycles >= 4104) && (ppu.vBlankClockCycles >= 4) && (ppu.scanline == 153) {
			ppu.scanline = 0
			ppu.memoryBus.write(ly, ppu.scanline)
			ppu.compareLY()
		}

		// Vblank mode ends
		if ppu.statusClockCycles >= 4560 {
			ppu.statusClockCycles -= 4560
			ppu.updateStatusMode(oamScan)

			stat := ppu.memoryBus.read(stat)
			if bit.IsSet(stat, 5) {
				ppu.interruptBus.request(lcdStatInterrupt)
			}
		}

	case oamScan:
		if ppu.statusClockCycles >= 80 {
			ppu.statusClockCycles -= 80
			ppu.updateStatusMode(drawMode)
		}

	case drawMode:
		if ppu.statusClockCycles >= 172 {
			ppu.statusClockCycles -= 172

			ppu.renderScanline()
			ppu.updateStatusMode(hblank)

			stat := ppu.memoryBus.read(stat)
			if bit.IsSet(stat, 3) {
				ppu.interruptBus.request(lcdStatInterrupt)
			}
		}
	}
}

// updateStatusMode updates status mode.
// If the mode is VBLANK_MODE, it sets the vblankLine to 0, sets the vBlankClockCycles to the current statusClockCycles,
// and requests a VBLANK interrupt.
func (ppu *ppu) updateStatusMode(mode int) {
	ppu.statusMode = mode
	statVal := ppu.memoryBus.read(stat)
	ppu.memoryBus.write(stat, uint8(statVal&0xfc)|uint8(ppu.statusMode&0x3))

	if mode == vblank {
		ppu.vblankLine = 0
		ppu.vBlankClockCycles = ppu.statusClockCycles
		ppu.interruptBus.request(lcdStatInterrupt)
	}
}

// renderScanline renders scanline into videobuf.
func (ppu *ppu) renderScanline() {
	lcdc := ppu.memoryBus.read(lcdc)

	if bit.IsSet(lcdc, 0) {
		ppu.renderTiles()
	}

	if bit.IsSet(lcdc, 1) {
		ppu.renderSprites()
	}
}

// renderTiles render tiles into videoBuf.
func (ppu *ppu) renderTiles() {
	lcdc := ppu.memoryBus.read(lcdc)
	scx := ppu.memoryBus.read(scx)
	scy := ppu.memoryBus.read(scy)
	wx := int(ppu.memoryBus.read(wx)) - 7
	wy := ppu.memoryBus.read(wy)
	palette := ppu.memoryBus.read(bgp)

	tileData, tileMap := ppu.getTileDataAndTileMap()

	var (
		tileRow uint16
		yPos    uint8
	)

	window := false
	if bit.IsSet(lcdc, 5) && wy <= ppu.scanline {
		window = true
		yPos = ppu.scanline
		tileRow = (uint16(yPos) / 8) * 32
	} else {
		// Use background layer
		yPos = ppu.scanline + scy
		tileRow = uint16(yPos/8) * 32
	}

	var line uint8 = (yPos % 8) * 2

	for x := uint16(0); x < 32; x++ {
		var tile uint8

		if tileData == 0x8800 {
			tile = uint8((int8(ppu.memoryBus.read(tileMap + uint16(tileRow) + x))))
			tile += 128
		} else {
			tile = ppu.memoryBus.read(tileMap + uint16(tileRow) + x)
		}

		mapOffsetX := x * 8
		tile16 := uint16(tile) * 16
		tileAddress := tileData + tile16 + uint16(line)

		b1 := ppu.memoryBus.read(tileAddress)
		b2 := ppu.memoryBus.read(tileAddress + 1)

		for bit := uint8(0); bit < 8; bit++ {
			var pixel uint8

			if window {
				pixel = uint8(int(mapOffsetX) + int(bit) + wx)
			} else {
				pixel = uint8(mapOffsetX) + bit - scx
			}

			if pixel >= ScreenWidth || ppu.scanline > 144 {
				continue
			}

			var colorIndex uint8
			if (b1 & (0x1 << (7 - bit))) != 0 {
				colorIndex = 1
			}
			if (b2 & (0x1 << (7 - bit))) != 0 {
				colorIndex |= 2
			}

			lineWidth := uint(ppu.scanline) * ScreenWidth
			position := lineWidth + uint(pixel)
			color := (palette >> (colorIndex * 2)) & 0x03
			ppu.videoBuffer[position] = color
		}
	}
}

// getTileDataAndTileMap returns vram tile data and vram tile map based on LCDC.
func (ppu *ppu) getTileDataAndTileMap() (uint16, uint16) {
	lcdc := ppu.memoryBus.read(lcdc)

	var (
		data uint16 = 0x8800
		addr uint16 = 0x9800
	)

	if bit.IsSet(lcdc, 4) {
		data = 0x8000
	}

	if bit.IsSet(lcdc, 3) {
		addr = 0x9C00
	}

	return data, addr
}

// renderSprites render sprites into videoBuf.
func (ppu *ppu) renderSprites() {
	var (
		lcdc     = ppu.memoryBus.read(lcdc)
		palette0 = ppu.memoryBus.read(obp0)
		palette1 = ppu.memoryBus.read(obp1)
	)

	// Get sprite height
	var ySize int32 = 8
	if bit.IsSet(lcdc, 2) {
		ySize = 16
	}

	for sprite := uint16(0); sprite < 40; sprite++ {
		index := sprite * 4

		yPos := int32(ppu.memoryBus.read(uint16(0xFE00+index))) - 16
		if int32(ppu.scanline) < yPos || int32(ppu.scanline) >= (yPos+ySize) {
			continue
		}

		xPos := int32(ppu.memoryBus.read(uint16(0xFE00+index+1))) - 8
		tileData := ppu.memoryBus.read(uint16(0xFE00 + index + 2))
		flags := ppu.memoryBus.read(uint16(0xFE00 + index + 3))

		// Sprite flags
		// TODO: check xFlip and priority
		palette := bit.IsSet(flags, 4)
		yFlip := bit.IsSet(flags, 6)

		line := int32(ppu.scanline) - yPos
		if yFlip {
			line = ySize - line - 1
		}

		spriteAddr := (uint16(tileData) * 16) + uint16(line*2)
		b1 := ppu.memoryBus.read(spriteAddr)
		b2 := ppu.memoryBus.read(spriteAddr + 1)

		for tilePixel := uint8(0); tilePixel < 8; tilePixel++ {
			pixel := int16(xPos) + int16(7-tilePixel)

			if pixel < 0 || pixel >= ScreenWidth {
				continue
			}

			var colorIndex uint8
			if (b1 & (0x1 << (7 - tilePixel))) != 0 {
				colorIndex = 1
			}
			if (b2 & (0x1 << (7 - tilePixel))) != 0 {
				colorIndex |= 2
			}

			if colorIndex == 0 {
				continue
			}

			var spritePalette = palette0
			if palette {
				spritePalette = palette1
			}

			lineWidth := uint(ppu.scanline) * ScreenWidth
			position := lineWidth + uint(pixel)
			color := (spritePalette >> (colorIndex * 2)) & 0x03
			ppu.videoBuffer[position] = color
		}
	}
}

// compareLY compares LY and LYC
func (ppu *ppu) compareLY() {
	lyc := ppu.memoryBus.read(lyc)
	status := ppu.memoryBus.read(stat)

	if ppu.scanline == lyc {
		ppu.memoryBus.write(stat, bit.Set(status, 2))

		if bit.IsSet(ppu.memoryBus.read(stat), 6) {
			ppu.interruptBus.request(lcdStatInterrupt)
		}
	} else {
		ppu.memoryBus.write(stat, bit.Clear(status, 2))
	}
}
