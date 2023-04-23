package main

// PPU represents game boy pixel processing unit.
type PPU struct {
	ticks    int
	videoBuf [GB_W * GB_H]uint8
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

	LINE_TICKS  = 456
	FRAME_LINES = 154

	STAT_MODE_OAM_TICKS  = LINE_TICKS - 80
	STAT_MODE_DRAW_TICKS = STAT_MODE_OAM_TICKS - 172

	GB_W = 160 // Game boy screen width
	GB_H = 144 // Game boy screen height
)

const (
	STAT_MODE_HBLANK   = iota // MODE 0 HBLANK
	STAT_MODE_VBLANK          // MODE 1 VBLANK
	STAT_MODE_OAM_SCAN        // MODE 2 OAM SCAN
	STAT_MODE_DRAW            // MODE 3 DRAWING PIXELS
)

// Ticks emulates ppu ticks.
func (ppu *PPU) Tick(cpuTicks int) {
	ppu.tick()

	lcdc := memory.read(LCDC)
	if !isBitSet(lcdc, 7) {
		ppu.clearScreen()
		ppu.clearStatus()
		return
	}

	ppu.ticks -= cpuTicks

	if ppu.ticks <= 0 {
		currentLine := memory.read(LY)
		memory.write(LY, currentLine+1)

		currentLine = memory.read(LY)
		if currentLine > 153 {
			memory.write(LY, 0)
			ppu.clearScreen()
		}

		ppu.ticks += LINE_TICKS

		if currentLine == GB_H {
			gameboy.reqInterrupt(IT_VBLANK)
		}
	}
}

func (ppu *PPU) tick() {
	currentLine := memory.read(LY)
	status := memory.read(STAT)
	mode := status & 0x3

	switch {
	// STAT MODE VBLANK
	case currentLine >= 144:
		status = setBit(status, 0)
		status = clearBit(status, 1)

		if isBitSet(status, 4) && mode != STAT_MODE_VBLANK {
			gameboy.reqInterrupt(IT_LCD_STAT)
		}

		break

	// STAT MODE OAM
	case ppu.ticks >= STAT_MODE_OAM_TICKS:
		status = clearBit(status, 0)
		status = setBit(status, 1)

		if isBitSet(status, 5) && mode != STAT_MODE_OAM_SCAN {
			gameboy.reqInterrupt(IT_LCD_STAT)
		}

		break

	// STAT MODE DRAW
	case ppu.ticks >= STAT_MODE_DRAW_TICKS:
		status = setBit(status, 0)
		status = setBit(status, 1)

		if mode != STAT_MODE_DRAW {
			ppu.renderLY()
		}

		break

	// STAT MODE HBLANK
	default:
		status = clearBit(status, 0)
		status = clearBit(status, 1)

		if isBitSet(status, 3) && mode != STAT_MODE_HBLANK {
			gameboy.reqInterrupt(IT_LCD_STAT)
		}
	}

	ppu.compareLY()
}

// renderLY render current line into videobuf.
func (ppu *PPU) renderLY() {
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
	currentLine := memory.read(LY)

	tileData, tileMap := ppu.getTileDataAndTileMap()

	var (
		tileRow uint16
		yPos    uint8
	)

	window := false
	if isBitSet(lcdc, 5) && wy <= currentLine {
		window = true
		yPos = currentLine
		tileRow = (uint16(yPos) / 8) * 32
	} else {
		// Use background layer
		yPos = currentLine + scy
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
			var bufferX uint8

			if window {
				bufferX = uint8(int(mapOffsetX) + int(bit) + wx)
			} else {
				bufferX = uint8(mapOffsetX) + bit - scx
			}

			if bufferX >= GB_W {
				continue
			}

			var colorIndex uint8
			if (b1 & (0x1 << (7 - bit))) != 0 {
				colorIndex = 1
			}
			if (b2 & (0x1 << (7 - bit))) != 0 {
				colorIndex |= 2
			}

			lineWidth := uint(currentLine) * GB_W
			position := lineWidth + uint(bufferX)
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

// TODO
func (ppu *PPU) renderSprites() {}

// clearScreen sets all pixels of videoBuf to white palette index.
func (ppu *PPU) clearScreen() {
	ppu.videoBuf = [GB_W * GB_H]uint8{}

	for i := 0; i < len(ppu.videoBuf); i++ {
		ppu.videoBuf[i] = 0
	}
}

// clearStatus sets STAT default value.
func (ppu *PPU) clearStatus() {
	ppu.ticks = LINE_TICKS
	memory.write(LY, 0)
	memory.write(STAT, 0x81)
}

// compareLY compares LY and LYC
func (ppu *PPU) compareLY() {
	ly := memory.read(LY)
	lyc := memory.read(LYC)
	status := memory.read(STAT)

	if ly == lyc {
		memory.write(STAT, setBit(status, 2))

		if isBitSet(memory.read(STAT), 6) {
			gameboy.reqInterrupt(IT_LCD_STAT)
		}
	} else {
		memory.write(STAT, clearBit(status, 2))
	}
}
