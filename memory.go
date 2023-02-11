package main

import "fmt"

// gbaMemory represents some type of gameboy console memory.
type gbaMemory struct {
	data  []uint8
	start uint16
	end   uint16
}

// validAddr returns a bool indicating if a given address is valid.
// given address is valid if it's inside the limits of start and stop.
func (m gbaMemory) validAddr(addr uint16) bool {
	return addr >= m.start && addr <= m.end
}

// Memory represent gameboy console memory (ram, rom, vram and so on)
type Memory struct {
	romBank0   gbaMemory
	romBank1   gbaMemory
	vram       gbaMemory
	ram        gbaMemory
	wramBank0  gbaMemory
	wramBank1  gbaMemory
	ioRegister gbaMemory
	hram       gbaMemory
	ie         gbaMemory
}

// NewMemory returns Memory instance with all banks divided by type.
func NewMemory() *Memory {
	return &Memory{
		romBank0: gbaMemory{
			data:  make([]uint8, 16384),
			start: 0x000,
			end:   0x3fff,
		},
		romBank1: gbaMemory{
			data:  make([]uint8, 16384),
			start: 0x4000,
			end:   0x7fff,
		},
		vram: gbaMemory{
			data:  make([]uint8, 8192),
			start: 0x8000,
			end:   0x9fff,
		},
		ram: gbaMemory{
			data:  make([]uint8, 8192),
			start: 0xa000,
			end:   0xbfff,
		},
		wramBank0: gbaMemory{
			data:  make([]uint8, 4096),
			start: 0xc000,
			end:   0xcfff,
		},
		wramBank1: gbaMemory{
			data:  make([]uint8, 4096),
			start: 0xd000,
			end:   0xdfff,
		},
		ioRegister: gbaMemory{
			data:  make([]uint8, 127),
			start: 0xff00,
			end:   0xff7f,
		},
		hram: gbaMemory{
			data:  make([]uint8, 126),
			start: 0xff80,
			end:   0xfffe,
		},
		ie: gbaMemory{
			data:  make([]uint8, 1),
			start: 0xffff,
			end:   0xffff,
		},
	}
}

// read fetch the value of a given address.
func (m Memory) read(addr uint16) (uint16, error) {
	banks := []gbaMemory{
		m.romBank0, m.romBank1, m.vram, m.ram, m.wramBank0, m.wramBank1,
		m.ioRegister, m.hram, m.ie,
	}

	for _, b := range banks {
		if b.validAddr(addr) {
			b1 := byte(b.data[addr])
			b2 := byte(b.data[addr+1])
			val := uint16(b1)<<8 | uint16(b2)
			return val, nil
		}
	}

	return 0, fmt.Errorf("Unable to read (0x%x)", addr)
}
