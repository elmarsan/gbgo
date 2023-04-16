package main

// PPU represents game boy pixel processing unit.
type PPU struct {
	ticks int
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

	LINE_TICKS  = 456
	FRAME_LINES = 154

	STAT_MODE_HBLANK   = iota // MODE 0 HBLANK
	STAT_MODE_VBLANK          // MODE 1 VBLANK
	STAT_MODE_OAM_SCAN        // MODE 2 OAM SCAN
	STAT_MODE_DRAW            // MODE 3 DRAWING PIXELS

	STAT_MODE_OAM_TICKS  = LINE_TICKS - 80
	STAT_MODE_DRAW_TICKS = STAT_MODE_OAM_TICKS - 172
)

func (ppu *PPU) updatePixels() {
	if !ppu.isLCDEnabled() {
		ppu.clearLCD()
		return
	}

	status := memory.read(STAT)
	mode := status & 0x3
	ly := memory.read(LY)

	switch {
	// STAT MODE VBLANK
	case ly >= 144:
		status = setBit(status, 0)
		status = clearBit(status, 1)

		if isBitSet(status, 4) {
			gameboy.reqInterrupt(IT_LCD_STAT)
		}

	// STAT MODE OAM
	case ppu.ticks >= STAT_MODE_OAM_TICKS:
		status = clearBit(status, 0)
		status = setBit(status, 1)

		if isBitSet(status, 5) {
			gameboy.reqInterrupt(IT_LCD_STAT)
		}

	// STAT MODE DRAW
	case ppu.ticks >= STAT_MODE_DRAW_TICKS:
		status = setBit(status, 0)
		status = setBit(status, 1)

		if mode != STAT_MODE_DRAW {
			ppu.render()
		}

	// STAT MODE HBLANK
	default:
		status = clearBit(status, 0)
		status = clearBit(status, 1)

		if isBitSet(status, 3) {
			gameboy.reqInterrupt(IT_LCD_STAT)
		}
	}

	memory.write(STAT, status)
	ppu.compareLY()
}

func (ppu *PPU) render() {
	lcd := memory.read(LCDC)

	if isBitSet(lcd, 0) {
		ppu.renderTiles()
	}

	if isBitSet(lcd, 1) {
		ppu.renderSprites()
	}
}

func (ppu *PPU) renderTiles() {}

func (ppu *PPU) renderSprites() {}

func (ppu *PPU) isLCDEnabled() bool {
	return isBitSet(memory.read(LCDC), 7)
}

func (ppu *PPU) clearLCD() {
	ppu.ticks = LINE_TICKS
	memory.write(LY, 0)
	memory.write(LYC, 0)
	status := memory.read(STAT)
	status = clearBit(status, 0)
	status = clearBit(status, 1)
	memory.write(STAT, status)
}

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
