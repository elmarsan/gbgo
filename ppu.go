package main

// PPU represents game boy pixel processing unit.
type PPU struct{}

var ppu = &PPU{}

// Sprite represents drawable object.
type Sprite struct {
	x         uint8
	y         uint8
	tileIndex uint8
	flags     uint8
}

// Palette number **Non CGB Mode Only** (0=OBP0, 1=OBP1)
func (s *Sprite) palette() bool {
	return readBit(s.flags, 4)
}

// X flip (0=Normal, 1=Horizontally mirrored)
func (s *Sprite) flipX() bool {
	return readBit(s.flags, 5)
}

// Y flip (0=Normal, 1=Vertically mirrored)
func (s *Sprite) flipY() bool {
	return readBit(s.flags, 6)
}

// BG and Window over OBJ (0=No, 1=BG and Window colors 1-3 over the OBJ)
func (s *Sprite) bg() bool {
	return readBit(s.flags, 7)
}
