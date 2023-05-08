package main

import (
	"log"
)

// Gameboy represents an emulator of Gameboy console.
type Gameboy struct {
	// whether the console is paused.
	paused bool

	// channel for signaling the console to stop.
	stopCh chan struct{}

	// cpu represents the central processing unit used by Gameboy.
	cpu *CPU

	// bus represents memory bus used by Gameboy.
	bus *MemoryBus

	// ppu represents pixel processing unit user by Gameboy.
	ppu *PPU

	// timer represents the timer of Gameboy.
	timer *Timer

	// cartridge represents the external piece of hardware that contains game data.
	cartridge *Cartridge

	// joypad represents physycal buttons of Gameboy.
	joypad *Joypad
}

// NewGameboy creates and returns a new Gameboy instance.
func NewGameboy() *Gameboy {
	cartridge := &Cartridge{}
	joypad := &Joypad{}

	bus := NewMemoryBus(cartridge, joypad)

	cpu := NewCPU(bus)
	ppu := NewPPU(bus)

	return &Gameboy{
		paused:    false,
		stopCh:    make(chan struct{}),
		cpu:       cpu,
		bus:       bus,
		ppu:       ppu,
		timer:     &Timer{},
		cartridge: cartridge,
		joypad:    joypad,
	}
}

// Run runs the main loop of the Gameboy, which repeatedly executes instructions
// while the console is not paused, and waits for a signal on the stopCh channel
// to stop the console.
func (gb *Gameboy) Run() {
	for {
		select {
		case <-gb.stopCh:
			close(gb.stopCh)
			break

		default:
			if !gb.paused {
				gb.step()
			}
		}
	}
}

// Stop sends a signal on the stopCh channel to stop the Gameboy.
func (gb *Gameboy) Stop() {
	log.Println("Sending gameboy stop signal")
	gb.stopCh <- struct{}{}
}

// step executes the next instruction in the CPU, updates the PPU and timer,
// and handles interrupts
func (gb *Gameboy) step() {
	if !gb.cpu.halted {
		gb.cpu.execute()
	} else {
		if gb.interruptPending() {
			gb.cpu.halted = false
		} else {
			gb.cpu.clockCycles += 4
		}
	}

	if gb.cpu.ime {
		if gb.interruptPending() {
			gb.executeISR()
		}
	}

	if gb.cpu.enablingIme {
		gb.cpu.ime = true
		gb.cpu.enablingIme = false
	}

	gb.ppu.Tick(gb.cpu.clockCycles)
	gb.timer.Tick(gb.cpu.clockCycles)
}

// LoadRom loads a ROM file into the Gameboy's cartridge and returns an error if
// the file could not be loaded.
func (gb *Gameboy) LoadRom(rom string) error {
	err := gb.cartridge.load(rom)
	if err != nil {
		return err
	}

	return nil
}
