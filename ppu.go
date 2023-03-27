package main

// PPU represents game boy pixel processing unit.
type PPU struct {
	pixels [160][144]uint32
}

var ppu = &PPU{}

const (
	LCDC = 0xff40 // LCD control
	STAT = 0xff41 // LCD status
	SCY  = 0xff42 // Viewport Y position
	SCX  = 0xff43 // Viewport X position
	WY   = 0xff4a // Window Y position
	WX   = 0xff4b // Window X position
	LY   = 0xff44 // LCD Y coordinate
	LYC  = 0xff45 // LY compare
)

var palette [4]uint32 = [4]uint32{0xffffffff, 0xffaaaaaa, 0xff555555, 0xff000000}

// update updates pixels of the lcd screen.
func (ppu *PPU) update() {
	if !ppu.isLdcEnabled() {
		return
	}

	ppu.renderTiles()
	// ppu.renderSprites()
}

// renderTiles render tiles in lcd screen.
func (ppu *PPU) renderTiles() {
	addressingMode := readBit(memory.read(LCDC), 7)

	// addressing mode 0 (0x8000 - 0x87ff)
	var startAddr uint16 = 0x8000
	// endAddr := 0x87ff

	// addresing mode 1 (0x8800 - 0x8fff)
	if addressingMode == 1 {
		startAddr = 0x8800
		// endAddr = 0x8fff
	}

	tileIndex := 0

	yPos, xPos := 0, 0

	for i := 0; i < 8; i++ {
		for x := 0; x < 16; x++ {
			for tileY := uint16(0); tileY < 16; tileY += 2 {
				b1Addr := startAddr + (uint16(tileIndex) * 16) + tileY
				b2Addr := startAddr + (uint16(tileIndex) * 16) + tileY + 1

				b1 := memory.read(b1Addr)
				b2 := memory.read(b2Addr)

				for bit := uint8(7); bit >= 0; bit-- {
					// color := readBit(b1, bit) | readBit(b2, bit)
					hi := uint8((b1&(1<<bit))>>bit) << 1
					lo := uint8((b2 & (1 << bit)) >> bit)
					color := hi | lo

					ppu.pixels[yPos][xPos] = palette[color]
				}
			}
		}

		yPos += 8
		xPos = 0
	}
}

func (ppu *PPU) renderSprites() {

}

// isLcdcEnabled returns bool flag indicating if LCDC bit 7 is set.
// 7 LCD and PPU enable	0=Off, 1=On
func (ppu *PPU) isLdcEnabled() bool {
	return isBitSet(memory.read(LCDC), 7)
}

// bgWindowPriority returns bool flag indicating if LCDC bit 0 is set.
// BG and Window enable/priority 0=Off, 1=On
func (ppu *PPU) bgWindowPriority() bool {
	return isBitSet(memory.read(LCDC), 0)
}

// Sprite represents drawable object.
type Sprite struct {
	x         uint8
	y         uint8
	tileIndex uint8
	flags     uint8
}

// Palette number **Non CGB Mode Only** (0=OBP0, 1=OBP1)
func (s *Sprite) palette() bool {
	return readBit(s.flags, 4) == 1
}

// X flip (0=Normal, 1=Horizontally mirrored)
func (s *Sprite) flipX() bool {
	return readBit(s.flags, 5) == 1
}

// Y flip (0=Normal, 1=Vertically mirrored)
func (s *Sprite) flipY() bool {
	return readBit(s.flags, 6) == 1
}

// BG and Window over OBJ (0=No, 1=BG and Window colors 1-3 over the OBJ)
func (s *Sprite) bg() bool {
	return readBit(s.flags, 7) == 1
}
