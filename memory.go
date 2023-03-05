package main

import "log"

// Memory map
const (
	CARTRIDGE_END uint16 = 0x7fff

	VRAM_START uint16 = 0x8000
	VRAM_END   uint16 = 0x9fff

	RAM_START uint16 = 0xa000
	RAM_END   uint16 = 0xbfff

	WRAM_START uint16 = 0xc000
	WRAM_END   uint16 = 0xdfff

	OAM_START uint16 = 0xfe00
	OAM_END   uint16 = 0xfe9f

	IO_START uint16 = 0xff00
	IO_END   uint16 = 0xff7f

	HRAM_START uint16 = 0xff80
	HRAM_END   uint16 = 0xfffe
)

// Memory represents gb memory.
type Memory struct {
	ram [0x2000]uint8

	// wram represents working ram.
	wram [0x2000]uint8

	// hram represents high ram.
	hram [0x80]uint8

	// vram represents video ram.
	vram [0x2000]uint8

	// oams represents object attribute memory.
	oams [40]uint8
}

var memory = &Memory{}

// read reads from memory address.
func (m *Memory) read(addr uint16) uint8 {
	switch {
	case addr <= CARTRIDGE_END:
		return cartridge.read(addr)

	case addr <= VRAM_END:
		return m.vram[addr-VRAM_START]

	case addr <= RAM_END:
		return m.ram[addr-RAM_START]

	case addr <= WRAM_END:
		return m.wram[addr-WRAM_START]

	case addr <= OAM_END:
		return m.oams[addr-OAM_START]

	case addr <= IO_END:
		// TODO
		return 0

	case addr <= HRAM_END:
		return m.hram[addr-HRAM_START]

	default:
		log.Fatalf("Invalid memory address 0%d", addr)
		return 0
	}
}

// write writes value into memory address.
func (m *Memory) write(addr uint16, val uint8) {
	switch {
	case addr <= CARTRIDGE_END:
		cartridge.write(addr, val)
		break

	case addr <= VRAM_END:
		m.vram[addr-VRAM_START] = val
		break

	case addr <= RAM_END:
		m.ram[addr-RAM_START] = val
		break

	case addr <= WRAM_END:
		m.wram[addr-WRAM_START] = val
		break

	case addr <= OAM_END:
		m.oams[addr-OAM_START] = val
		break

	case addr <= IO_END:
		// TODO
		break

	case addr <= HRAM_END:
		m.hram[addr-HRAM_START] = val
		break

	default:
		log.Fatalf("Invalid memory address 0%d", addr)
	}
}
