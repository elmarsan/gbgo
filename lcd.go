package main

// 7    LCD and PPU enable	0=Off, 1=On
// 6	Window tile map area	0=9800-9BFF, 1=9C00-9FFF
// 5	Window enable	0=Off, 1=On
// 4	BG and Window tile data area	0=8800-97FF, 1=8000-8FFF
// 3	BG tile map area	0=9800-9BFF, 1=9C00-9FFF
// 2	OBJ size	0=8x8, 1=8x16
// 1	OBJ enable	0=Off, 1=On
// 0	BG and Window enable/priority	0=Off, 1=On

type LCD struct{}

var lcd = &LCD{}

// isEnabled returns bool indicating if lcd and ppu are enabled.
// bit 7 (0=Off, 1=On)
func (lcd *LCD) isEnabled() bool {
	reg := memory.read(LCDC)
	return readBit(reg, 7) == 1
}

// winTileMapArea returns addrRange of window tile map area.
// bit 6 of 0xff40 (0=9800-9BFF, 1=9C00-9FFF).
func (lcd *LCD) winTileMapArea() addrRange {
	reg := memory.read(LCDC)

	isSet := readBit(reg, 6) == 1
	if !isSet {
		return addrRange{
			start: 0x9800,
			end:   0x9bff,
		}
	}

	return addrRange{
		start: 0x9c00,
		end:   0x9fff,
	}
}

// isWinEnabled returns bool indicating if window is enabled.
// bit 5 (0=Off, 1=On)
func (lcd *LCD) isWinEnabled() bool {
	reg := memory.read(LCDC)
	return readBit(reg, 5) == 1
}

// bgWinTileDataArea returns addrRange of bg and window tile data area.
// bit 4 of 0xff40 (0=8800-97FF, 1=8000-8FFF).
func (lcd *LCD) bgWinTileDataArea() addrRange {
	reg := memory.read(LCDC)

	isSet := readBit(reg, 4) == 1
	if !isSet {
		return addrRange{
			start: 0x8800,
			end:   0x97ff,
		}
	}

	return addrRange{
		start: 0x8000,
		end:   0x8fff,
	}
}

// bgTileMapArea returns addrRange of bg tile data area.
// bit 3 of 0xff40 (0=9800-9BFF, 1=9C00-9FFF).
func (lcd *LCD) bgTileMapArea() addrRange {
	reg := memory.read(LCDC)

	isSet := readBit(reg, 3) == 1
	if !isSet {
		return addrRange{
			start: 0x9800,
			end:   0x9bff,
		}
	}

	return addrRange{
		start: 0x9c00,
		end:   0x9fff,
	}
}

// objSize returns type of obj size.
// bit 2 of 0xff40 (0=8x8, 1=8x16).
func objSize() uint8 {
	reg := memory.read(LCDC)
	return readBit(reg, 2)
}

// isObjEnabled returns bool indicating if obj is enabled.
// bit 1 of 0xff40 (0=Off, 1=On).
func isObjEnabled() bool {
	reg := memory.read(LCDC)
	return readBit(reg, 1) == 1
}

// isObjEnabled returns bool indicating if obj bg and window are enabled.
// bit 0 of 0xff40 (0=Off, 1=On).
func isBgAndWinEnabled() bool {
	reg := memory.read(LCDC)
	return readBit(reg, 0) == 1
}
