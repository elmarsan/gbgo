package main

func read(addr uint16) uint8 {
	if addr < 0x8000 {
		return cartridge.Read(addr)
	}

	return 0
}

func write(addr uint16, val uint8) {}
