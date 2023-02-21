package main

type Memory struct{}

var memory = &Memory{}

func (m *Memory) read(addr uint16) uint8 {
	if addr < 0x8000 {
		return cartridge.Read(addr)
	}

	return 0
}

func (m *Memory) write(addr uint16, val uint8) {}
