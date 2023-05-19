package gameboy

import (
	"fmt"
	"os"
)

// Memory map
const (
	cartridgeEnd = 0x7fff
	vramStart    = 0x8000
	vramEnd      = 0x9fff
	ramStart     = 0xa000
	ramEnd       = 0xbfff
	wramStart    = 0xc000
	wramEnd      = 0xdfff
	echoStart    = 0xe000
	echoEnd      = 0xfdff
	oamStart     = 0xfe00
	oamEnd       = 0xfe9f
	ioStart      = 0xff00
	ioEnd        = 0xff7f
	hramStart    = 0xff80
	hramEnd      = 0xfffe
	ieReg        = 0xffff
)

// memoryBus represents Gameboy memory bus for r/w different memory banks.
type memoryBus struct {
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
	cartridge *cartridge

	// joypad represents physycal buttons of Gameboy.
	joypad *Joypad
}

// newMemoryBus creates and returns a new memoryBus instance.
// Some values are set to defaults of DMG boot sequence.
func newMemoryBus(cartridge *cartridge, joypad *Joypad) *memoryBus {
	bus := &memoryBus{
		cartridge: cartridge,
		joypad:    joypad,
	}

	// Timer
	bus.io[tima-ioStart] = 0xab
	bus.write(tima, 0x00)
	bus.write(tma, 0x00)
	bus.write(tac, 0xf8)

	// Interrupt
	bus.write(iFF, 0xe1)
	bus.write(ie, 0x00)

	// PPU
	bus.write(lcdc, 0x91)
	bus.write(stat, 0x81)
	bus.write(scx, 0)
	bus.write(scy, 0)
	bus.write(wx, 0)
	bus.write(wy, 0)
	bus.write(ly, 0x91)
	bus.write(bgp, 0xfc)

	// JOYPAD
	bus.write(joyp, 0xff)

	return bus
}

// read reads from memory address.
func (bus *memoryBus) read(addr uint16) uint8 {
	switch {
	case addr <= cartridgeEnd:
		return bus.cartridge.read(addr)

	case addr <= vramEnd:
		return bus.vram[addr-vramStart]

	case addr <= ramEnd:
		return bus.ram[addr-ramStart]

	case addr <= wramEnd:
		return bus.wram[addr-wramStart]

		// prohibited area
	case addr <= echoEnd:
		return 0

	case addr <= oamEnd:
		return bus.oam[addr-oamStart]

	case addr <= 0xfeff:
		return 0

	case addr <= ioEnd:
		if addr == joyp {
			return bus.joypad.Get()
		}

		return bus.io[addr-ioStart]

	case addr <= hramEnd:
		return bus.hram[addr-hramStart]

	case addr == ieReg:
		return bus.ie

	default:
		fmt.Printf("Invalid memory address 0%x", addr)
		os.Exit(1)
		return 0
	}
}

// write writes value into memory address.
func (bus *memoryBus) write(addr uint16, val uint8) {
	switch {
	case addr <= cartridgeEnd:
		return

	case addr <= vramEnd:
		bus.vram[addr-vramStart] = val
		break

	case addr <= ramEnd:
		bus.ram[addr-ramStart] = val
		break

	case addr <= wramEnd:
		bus.wram[addr-wramStart] = val
		break

		// prohibited area
	case addr <= echoEnd:
		break

	case addr <= oamEnd:
		bus.oam[addr-oamStart] = val
		break

	case addr <= 0xfeff:
		break

	case addr == div:
		bus.io[addr-ioStart] = 0
		break

	// DMA transfer
	case addr == 0xff46:
		bus.dmaTransfer(val)
		bus.io[addr-ioStart] = val
		break

	case addr <= ioEnd:
		if addr == joyp {
			bus.joypad.Set(val)
		} else {
			bus.io[addr-ioStart] = val
		}

		break

	case addr <= hramEnd:
		bus.hram[addr-hramStart] = val
		break

	case addr == ieReg:
		bus.ie = val
		break

	default:
		fmt.Printf("Invalid memory address 0%x", addr)
		os.Exit(1)
	}
}

// dmaTransfer performs dma transfer
func (bus *memoryBus) dmaTransfer(val uint8) {
	addr := uint16(val) * 0x100

	for i := uint16(0); i < 0xa0; i++ {
		bus.write(oamStart+i, bus.read(addr+i))
	}
}
