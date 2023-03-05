package main

import "log"

type Ram struct {
	wram [0x200]uint8
	hram [0x80]uint8
}

var ram = &Ram{}

// readWram reads from wram memory address $c000 - $dfff.
func (ram *Ram) readWram(addr uint16) uint8 {
	if addr < WRAM_BANK_0_START || addr > WRAM_BANK_1_END {
		log.Fatalf("Invalid wram address 0x%d", addr)
	}

	return ram.wram[addr-WRAM_BANK_0_START]
}

// writeWram writes in wram memory address $c000 - $dfff.
func (ram *Ram) writeWram(addr uint16, val uint8) {
	if addr < WRAM_BANK_0_START || addr > WRAM_BANK_1_END {
		log.Fatalf("Invalid wram address 0x%d", addr)
	}

	ram.wram[addr-WRAM_BANK_0_START] = val
}

// readHram reads from hram memory address $ff80 - $dfffe.
func (ram *Ram) readHram(addr uint16) uint8 {
	if addr < HRAM_START || addr > HRAM_END {
		log.Fatalf("Invalid hram address 0x%d", addr)
	}

	return ram.hram[addr-HRAM_START]
}

// writeHram writes in Hram memory address $ff80 - $dfffe.
func (ram *Ram) writeHram(addr uint16, val uint8) {
	if addr < HRAM_START || addr > HRAM_END {
		log.Fatalf("Invalid hram address 0x%d", addr)
	}

	ram.wram[addr-HRAM_START] = val
}
