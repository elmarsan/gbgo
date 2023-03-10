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

	ECHO_START uint16 = 0xe000
	ECHO_END   uint16 = 0xfdff

	OAM_START uint16 = 0xfe00
	OAM_END   uint16 = 0xfe9f

	IO_START uint16 = 0xff00
	IO_END   uint16 = 0xff7f

	HRAM_START uint16 = 0xff80
	HRAM_END   uint16 = 0xfffe

	IE_REG uint16 = 0xffff
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

	// oam represents object attribute memory.
	oam [40]uint8

	// io represents hardware register
	io [43]uint8

	// ie represents interrupt enable register
	ie uint8
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

		// prohibited area
	case addr <= ECHO_END:
		return 0

	case addr <= OAM_END:
		return m.oam[addr-OAM_START]

	case addr <= IO_END:
		return m.io[addr-IO_START]

	case addr <= HRAM_END:
		return m.hram[addr-HRAM_START]

	case addr == IE_REG:
		return m.ie

	default:
		log.Fatalf("Invalid memory address 0%x", addr)
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

		// prohibited area
	case addr <= ECHO_END:
		break

	case addr <= OAM_END:
		m.oam[addr-OAM_START] = val
		break

	case addr <= IO_END:
		m.hram[addr-IO_START] = val
		break

	case addr <= HRAM_END:
		m.hram[addr-HRAM_START] = val
		break

	case addr == IE_REG:
		m.ie = val
		break

	default:
		log.Fatalf("Invalid memory address 0%x", addr)
	}
}
