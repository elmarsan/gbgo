package main

// Gameboy represents an emulation of gameboy console
type Gameboy struct {
	cpu    *CPU
	memory *Memory
}

// NewGameboy returns Gameboy instance.
func NewGameboy() *Gameboy {
	return &Gameboy{
		cpu:    &CPU{},
		memory: NewMemory(),
	}
}
