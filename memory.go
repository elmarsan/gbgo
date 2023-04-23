package main

import (
	"fmt"
	"os"
)

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
	oam [0x100]uint8

	// io represents hardware register
	io [0x80]uint8

	// ie represents interrupt enable register
	ie uint8
}

func (m *Memory) init() {
	// Timer
	m.io[TIMA-IO_START] = 0xab
	m.write(TIMA, 0x00)
	m.write(TMA, 0x00)
	m.write(TAC, 0xf8)

	// Interrupt
	m.write(IF, 0xe1)
	m.write(IE, 0x00)

	// PPU
	m.write(LCDC, 0x91)
	m.write(STAT, 0x81)
	m.write(SCX, 0)
	m.write(SCY, 0)
	m.write(WX, 0)
	m.write(WY, 0)
	m.write(LY, 0x91)
	m.write(BGP, 0xfc)

	// DMA
	m.write(0xff46, 0xff)
}

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

	case addr <= 0xfeff:
		return 0

	case addr <= IO_END:
		return m.io[addr-IO_START]

	case addr <= HRAM_END:
		return m.hram[addr-HRAM_START]

	case addr == IE_REG:
		return m.ie

	default:
		fmt.Printf("Invalid memory address 0%x", addr)
		os.Exit(1)
		return 0
	}
}

// write writes value into memory address.
func (m *Memory) write(addr uint16, val uint8) {
	switch {
	case addr <= CARTRIDGE_END:
		return

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

	case addr <= 0xfeff:
		break

	case addr == DIV:
		m.io[addr-IO_START] = 0
		break

	// DMA transfer
	case addr == 0xff46:
		m.dmaTransfer(val)
		break

	case addr <= IO_END:
		m.io[addr-IO_START] = val
		break

	case addr <= HRAM_END:
		m.hram[addr-HRAM_START] = val
		break

	case addr == IE_REG:
		m.ie = val
		break

	default:
		fmt.Printf("Invalid memory address 0%x", addr)
		os.Exit(1)
	}
}

// dmaTransfer performs dma transfer
func (m *Memory) dmaTransfer(val uint8) {
	addr := uint16(val) * 0x100

	for i := uint16(0); i < 0xa0; i++ {
		m.write(OAM_START+i, m.read(addr+i))
	}
}
