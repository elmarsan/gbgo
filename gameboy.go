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

	// memoryBus represents memory memoryBus used by Gameboy.
	memoryBus *MemoryBus

	// ppu represents pixel processing unit user by Gameboy.
	ppu *PPU

	// timer represents the timer of Gameboy.
	timer *Timer

	// cartridge represents the external piece of hardware that contains game data.
	cartridge *Cartridge

	// joypad represents physycal buttons of Gameboy.
	joypad *Joypad

	// InterruptBus represents the interruption system of game boy.
	interruptBus *InterruptBus
}

// NewGameboy creates and returns a new Gameboy instance.
func NewGameboy() *Gameboy {
	cartridge := &Cartridge{}
	joypad := &Joypad{}

	memoryBus := NewMemoryBus(cartridge, joypad)
	interruptBus := NewInterruptBus(memoryBus)

	cpu := NewCPU(memoryBus)
	ppu := NewPPU(memoryBus, interruptBus)

	timer := NewTimer(memoryBus, interruptBus)

	return &Gameboy{
		paused:       false,
		stopCh:       make(chan struct{}),
		cpu:          cpu,
		memoryBus:    memoryBus,
		ppu:          ppu,
		timer:        timer,
		cartridge:    cartridge,
		joypad:       joypad,
		interruptBus: interruptBus,
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
		gb.execute()
	} else {
		if gb.interruptBus.pending() {
		} else {
			gb.cpu.clockCycles += 4
		}
	}

	if gb.interruptBus.ime && gb.interruptBus.pending() {
		call(gb, gb.interruptBus.locateISR())
		gb.cpu.clockCycles += ISRClockCycles
	}

	if gb.interruptBus.enablingIme {
		gb.interruptBus.ime = true
		gb.interruptBus.enablingIme = false
	}

	gb.ppu.Tick(gb.cpu.clockCycles)
	gb.timer.Tick(gb.cpu.clockCycles)
}

// execute fetchs next opcode and executes the corresponding instruction.
func (gb *Gameboy) execute() {
	debug.logState(gb)

	gb.cpu.clockCycles = 0
	pc := gb.cpu.readPc()
	opcode := gb.memoryBus.read(pc)

	if opcode == 0xcb {
		pc := gb.cpu.readPc()
		opcode := gb.memoryBus.read(pc)
		prefixedInstructions[opcode](gb)
		gb.cpu.clockCycles += cbInstructionCycles[opcode] * 4
	} else {
		instructions[opcode](gb)
		gb.cpu.clockCycles += instructionCycles[opcode] * 4
	}
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
