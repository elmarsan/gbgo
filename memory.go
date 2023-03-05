package main

// Memory Map
const (
	CARTRIDGE_BANK_0_START uint16 = 0xfff
	CARTRIDGE_BANK_0_END   uint16 = 0x3fff
	CARTRIDGE_BANK_1_START uint16 = 0x4000
	CARTRIDGE_BANK_1_END   uint16 = 0x7fff

	VRAM_START uint16 = 0x8000
	VRAM_END   uint16 = 0x9fff

	RAM_START uint16 = 0xa000
	RAM_END   uint16 = 0xbfff

	WRAM_BANK_0_START uint16 = 0xc000
	WRAM_BANK_0_END   uint16 = 0xcfff
	WRAM_BANK_1_START uint16 = 0xd000
	WRAM_BANK_1_END   uint16 = 0xdfff

	OAM_START uint16 = 0xfe00
	OAM_END   uint16 = 0xfe9f

	NOT_USABLE_START uint16 = 0xfea0
	NOT_USABLE_END   uint16 = 0xfeff

	IO_REG_START uint16 = 0xff00
	IO_REG_END   uint16 = 0xff7f

	HRAM_START uint16 = 0xff80
	HRAM_END   uint16 = 0xfffe

	IE_REG = 0xfff
)

type Memory struct{}

var memory = &Memory{}

func (m *Memory) read(addr uint16) uint8 {
	if addr < CARTRIDGE_BANK_1_END {
		return cartridge.read(addr)
	}

	if addr < VRAM_END {
		// TODO
		return 0
	}

	if addr < RAM_END {
		// TODO
		return 0
	}

	if addr < WRAM_BANK_1_END {
		return ram.readWram(addr)
	}

	if addr < OAM_END {
		// TODO
		return 0
	}

	if addr < NOT_USABLE_END {
		// TODO
		return 0
	}

	if addr < IO_REG_END {
		// TODO
		return 0
	}

	if addr < HRAM_END {
		return ram.readHram(addr)
	}

	if addr < IE_REG {
		// TODO
		return 0
	}

	return 0
}

func (m *Memory) write(addr uint16, val uint8) {
	if addr < CARTRIDGE_BANK_1_END {
		cartridge.write(addr, val)
	}

	if addr < VRAM_END {
		// TODO
	}

	if addr < RAM_END {
		// TODO
	}

	if addr < WRAM_BANK_1_END {
		ram.writeWram(addr, val)
	}

	if addr < OAM_END {
		// TODO
	}

	if addr < NOT_USABLE_END {
		// TODO
	}

	if addr < IO_REG_END {
		// TODO
	}

	if addr < HRAM_END {
		ram.writeHram(addr, val)
	}

	if addr < IE_REG {
		// TODO
	}
}
