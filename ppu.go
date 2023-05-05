package main

// PPU represents game boy pixel processing unit.
type PPU struct {
	videoBuf       [GB_W * GB_H]uint8
	vBlankDots     int
	statusModeDots int
	statusMode     int
	scanline       uint8
	vblankLine     int
}

const (
	LCDC = 0xff40 // LCD control
	STAT = 0xff41 // LCD status
	SCY  = 0xff42 // Viewport Y position
	SCX  = 0xff43 // Viewport X position
	WY   = 0xff4a // Window Y position
	WX   = 0xff4b // Window X position
	LY   = 0xff44 // LCD Y coordinate
	LYC  = 0xff45 // LY compare
	BGP  = 0xff47 // Background palette
	OBP0 = 0xff48 // Object palette 0
	OBP1 = 0xff49 // Object palette 0

	GB_W = 160 // Game boy screen width
	GB_H = 144 // Game boy screen height
)

const (
	HBLANK_MODE   = iota // MODE 0 HBLANK
	VBLANK_MODE          // MODE 1 VBLANK
	OAM_SCAN_MODE        // MODE 2 OAM SCAN
	DRAW_MODE            // MODE 3 DRAWING PIXELS
)

func NewPPU() *PPU {
	videoBuf := [GB_W * GB_H]uint8{}

	for i := 0; i < len(videoBuf); i++ {
		videoBuf[i] = 0
	}

	return &PPU{
		videoBuf:       videoBuf,
		vBlankDots:     0,
		statusModeDots: 0,
		statusMode:     1,
		scanline:       144,
		vblankLine:     0,
	}
}

// Ticks emulates ppu ticks.
func (ppu *PPU) Tick(cycles int) {
	lcdc := memory.read(LCDC)

	if !isBitSet(lcdc, 7) {
		return
	}

	ppu.statusModeDots += cycles

	switch ppu.statusMode {

	case HBLANK_MODE:
		if ppu.statusModeDots >= 204 {
			ppu.statusModeDots -= 204
			ppu.updateStatusMode(OAM_SCAN_MODE)
			ppu.scanline++
			memory.write(LY, ppu.scanline)
			ppu.compareLY()

			// Vblank mode starts
			if ppu.scanline == 144 {
				ppu.updateStatusMode(VBLANK_MODE)

				stat := memory.read(STAT)
				if isBitSet(stat, 4) {
					gameboy.reqInterrupt(IT_LCD_STAT)
				}
			} else {
				stat := memory.read(STAT)
				if isBitSet(stat, 5) {
					gameboy.reqInterrupt(IT_LCD_STAT)
				}
			}
		}

	case VBLANK_MODE:
		ppu.vBlankDots += cycles

		if ppu.vBlankDots >= 456 {
			ppu.vBlankDots -= 456
			ppu.vblankLine++

			if ppu.vblankLine <= 9 {
				ppu.scanline++
				memory.write(LY, ppu.scanline)
				ppu.compareLY()
			}
		}

		if (ppu.statusModeDots >= 4104) && (ppu.vBlankDots >= 4) && (ppu.scanline == 153) {
			ppu.scanline = 0
			memory.write(LY, ppu.scanline)
			ppu.compareLY()
		}

		// Vblank mode ends
		if ppu.statusModeDots >= 4560 {
			ppu.statusModeDots -= 4560
			ppu.updateStatusMode(OAM_SCAN_MODE)

			stat := memory.read(STAT)
			if isBitSet(stat, 5) {
				gameboy.reqInterrupt(IT_LCD_STAT)
			}
		}

	case OAM_SCAN_MODE:
		if ppu.statusModeDots >= 80 {
			ppu.statusModeDots -= 80
			ppu.updateStatusMode(DRAW_MODE)
		}

	case DRAW_MODE:
		if ppu.statusModeDots >= 172 {
			ppu.statusModeDots -= 172

			ppu.renderScanline()
			ppu.updateStatusMode(HBLANK_MODE)

			stat := memory.read(STAT)
			if isBitSet(stat, 3) {
				gameboy.reqInterrupt(IT_LCD_STAT)
			}
		}
	}
}

func (ppu *PPU) updateStatusMode(mode int) {
	ppu.statusMode = mode
	stat := memory.read(STAT)
	memory.write(STAT, uint8(stat&0xfc)|uint8(ppu.statusMode&0x3))

	if mode == VBLANK_MODE {
		ppu.vblankLine = 0
		ppu.vBlankDots = ppu.statusModeDots
		gameboy.reqInterrupt(IT_VBLANK)
	}
}

// renderScanline renders scanline into videobuf.
func (ppu *PPU) renderScanline() {
	lcdc := memory.read(LCDC)

	if isBitSet(lcdc, 0) {
		ppu.renderTiles()
	}

	if isBitSet(lcdc, 1) {
		ppu.renderSprites()
	}
}

// renderTiles render tiles into videoBuf.
func (ppu *PPU) renderTiles() {
	lcdc := memory.read(LCDC)
	scx := memory.read(SCX)
	scy := memory.read(SCY)
	wx := int(memory.read(WX)) - 7
	wy := memory.read(WY)
	palette := memory.read(BGP)

	tileData, tileMap := ppu.getTileDataAndTileMap()

	var (
		tileRow uint16
		yPos    uint8
	)

	window := false
	if isBitSet(lcdc, 5) && wy <= ppu.scanline {
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
			tile = uint8((int8(memory.read(tileMap + uint16(tileRow) + x))))
			tile += 128
		} else {
			tile = memory.read(tileMap + uint16(tileRow) + x)
		}

		mapOffsetX := x * 8
		tile16 := uint16(tile) * 16
		tileAddress := tileData + tile16 + uint16(line)

		b1 := memory.read(tileAddress)
		b2 := memory.read(tileAddress + 1)

		for bit := uint8(0); bit < 8; bit++ {
			var pixel uint8

			if window {
				pixel = uint8(int(mapOffsetX) + int(bit) + wx)
			} else {
				pixel = uint8(mapOffsetX) + bit - scx
			}

			if pixel >= GB_W || ppu.scanline > 144 {
				continue
			}

			var colorIndex uint8
			if (b1 & (0x1 << (7 - bit))) != 0 {
				colorIndex = 1
			}
			if (b2 & (0x1 << (7 - bit))) != 0 {
				colorIndex |= 2
			}

			lineWidth := uint(ppu.scanline) * GB_W
			position := lineWidth + uint(pixel)
			color := (palette >> (colorIndex * 2)) & 0x03
			ppu.videoBuf[position] = color
		}
	}
}

// getTileDataAndTileMap returns vram tile data and vram tile map based on LCDC.
func (ppu *PPU) getTileDataAndTileMap() (uint16, uint16) {
	lcdc := memory.read(LCDC)

	var (
		data uint16 = 0x8800
		addr uint16 = 0x9800
	)

	if isBitSet(lcdc, 4) {
		data = 0x8000
	}

	if isBitSet(lcdc, 3) {
		addr = 0x9C00
	}

	return data, addr
}

// renderSprites render sprites into videoBuf.
func (ppu *PPU) renderSprites() {
	var (
		lcdc     = memory.read(LCDC)
		palette0 = memory.read(OBP0)
		palette1 = memory.read(OBP1)
	)

	// Get sprite height
	var ySize int32 = 8
	if isBitSet(lcdc, 2) {
		ySize = 16
	}

	for sprite := uint16(0); sprite < 40; sprite++ {
		index := sprite * 4

		yPos := int32(memory.read(uint16(0xFE00+index))) - 16
		if int32(ppu.scanline) < yPos || int32(ppu.scanline) >= (yPos+ySize) {
			continue
		}

		xPos := int32(memory.read(uint16(0xFE00+index+1))) - 8
		tileData := memory.read(uint16(0xFE00 + index + 2))
		flags := memory.read(uint16(0xFE00 + index + 3))

		// Sprite flags
		// TODO: check xFlip and priority
		palette := isBitSet(flags, 4)
		yFlip := isBitSet(flags, 6)

		line := int32(ppu.scanline) - yPos
		if yFlip {
			line = ySize - line - 1
		}

		spriteAddr := (uint16(tileData) * 16) + uint16(line*2)
		b1 := memory.read(spriteAddr)
		b2 := memory.read(spriteAddr + 1)

		for tilePixel := uint8(0); tilePixel < 8; tilePixel++ {
			pixel := int16(xPos) + int16(7-tilePixel)

			if pixel < 0 || pixel >= GB_W {
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

			lineWidth := uint(ppu.scanline) * GB_W
			position := lineWidth + uint(pixel)
			color := (spritePalette >> (colorIndex * 2)) & 0x03
			ppu.videoBuf[position] = color
		}
	}
}

// compareLY compares LY and LYC
func (ppu *PPU) compareLY() {
	lyc := memory.read(LYC)
	status := memory.read(STAT)

	if ppu.scanline == lyc {
		memory.write(STAT, setBit(status, 2))

		if isBitSet(memory.read(STAT), 6) {
			gameboy.reqInterrupt(IT_LCD_STAT)
		}
	} else {
		memory.write(STAT, clearBit(status, 2))
	}
}
