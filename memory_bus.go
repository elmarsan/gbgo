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

// MemoryBus represents Gameboy memory bus for r/w different memory banks.
type MemoryBus struct {
	// ram represents working ram.
	ram [0x2000]uint8

	// wram represents working ram.
	wram [0x2000]uint8

	// hram represents high ram.
	hram [0x80]uint8

	// vram represents video ram.
	vram [0x2000]uint8

	// oam represents object attribute memory.
	oam [0x100]uint8

	// io represents hardware register.
	io [0x80]uint8

	// ie represents interrupt enable register.
	ie uint8

	// cartridge represents the external piece of hardware that contains game data.
	cartridge *Cartridge

	// joypad represents physycal buttons of Gameboy.
	joypad *Joypad
}

// NewMemoryBus creates and returns a new MemoryBus instance.
// Some values are set to defaults of DMG boot sequence.
func NewMemoryBus(cartridge *Cartridge, joypad *Joypad) *MemoryBus {
	bus := &MemoryBus{
		cartridge: cartridge,
		joypad:    joypad,
	}

	// Timer
	bus.io[TIMA-IO_START] = 0xab
	bus.write(TIMA, 0x00)
	bus.write(TMA, 0x00)
	bus.write(TAC, 0xf8)

	// Interrupt
	bus.write(IF, 0xe1)
	bus.write(IE, 0x00)

	// PPU
	bus.write(LCDC, 0x91)
	bus.write(STAT, 0x81)
	bus.write(SCX, 0)
	bus.write(SCY, 0)
	bus.write(WX, 0)
	bus.write(WY, 0)
	bus.write(LY, 0x91)
	bus.write(BGP, 0xfc)

	// JOYPAD
	bus.write(JOYP, 0xff)

	return bus
}

// read reads from memory address.
func (bus *MemoryBus) read(addr uint16) uint8 {
	switch {
	case addr <= CARTRIDGE_END:
		return bus.cartridge.read(addr)

	case addr <= VRAM_END:
		return bus.vram[addr-VRAM_START]

	case addr <= RAM_END:
		return bus.ram[addr-RAM_START]

	case addr <= WRAM_END:
		return bus.wram[addr-WRAM_START]

		// prohibited area
	case addr <= ECHO_END:
		return 0

	case addr <= OAM_END:
		return bus.oam[addr-OAM_START]

	case addr <= 0xfeff:
		return 0

	case addr <= IO_END:
		if addr == JOYP {
			return bus.joypad.Get()
		}

		return bus.io[addr-IO_START]

	case addr <= HRAM_END:
		return bus.hram[addr-HRAM_START]

	case addr == IE_REG:
		return bus.ie

	default:
		fmt.Printf("Invalid memory address 0%x", addr)
		os.Exit(1)
		return 0
	}
}

// write writes value into memory address.
func (bus *MemoryBus) write(addr uint16, val uint8) {
	switch {
	case addr <= CARTRIDGE_END:
		return

	case addr <= VRAM_END:
		bus.vram[addr-VRAM_START] = val
		break

	case addr <= RAM_END:
		bus.ram[addr-RAM_START] = val
		break

	case addr <= WRAM_END:
		bus.wram[addr-WRAM_START] = val
		break

		// prohibited area
	case addr <= ECHO_END:
		break

	case addr <= OAM_END:
		bus.oam[addr-OAM_START] = val
		break

	case addr <= 0xfeff:
		break

	case addr == DIV:
		bus.io[addr-IO_START] = 0
		break

	// DMA transfer
	case addr == 0xff46:
		bus.dmaTransfer(val)
		bus.io[addr-IO_START] = val
		break

	case addr <= IO_END:
		if addr == JOYP {
			bus.joypad.Set(val)
		} else {
			bus.io[addr-IO_START] = val
		}

		break

	case addr <= HRAM_END:
		bus.hram[addr-HRAM_START] = val
		break

	case addr == IE_REG:
		bus.ie = val
		break

	default:
		fmt.Printf("Invalid memory address 0%x", addr)
		os.Exit(1)
	}
}

// dmaTransfer performs dma transfer
func (bus *MemoryBus) dmaTransfer(val uint8) {
	addr := uint16(val) * 0x100

	for i := uint16(0); i < 0xa0; i++ {
		bus.write(OAM_START+i, bus.read(addr+i))
	}
}
