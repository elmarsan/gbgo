package gameboy

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
	cpu *cpu

	// MemoryBus represents memory MemoryBus used by Gameboy.
	memoryBus *memoryBus

	// ppu represents pixel processing unit user by Gameboy.
	ppu *ppu

	// timer represents the timer of Gameboy.
	timer *timer

	// cartridge represents the external piece of hardware that contains game data.
	cartridge *cartridge

	// Joypad represents physycal buttons of Gameboy.
	Joypad *Joypad

	// InterruptBus represents the interruption system of game boy.
	interruptBus *interruptBus
}

// New creates and returns a new Gameboy instance.
func New() *Gameboy {
	cartridge := &cartridge{}
	joypad := &Joypad{}

	memoryBus := newMemoryBus(cartridge, joypad)
	interruptBus := newInterruptBus(memoryBus)

	cpu := newCpu(memoryBus)
	ppu := newPpu(memoryBus, interruptBus)

	timer := newTimer(memoryBus, interruptBus)

	return &Gameboy{
		paused:       false,
		stopCh:       make(chan struct{}),
		cpu:          cpu,
		memoryBus:    memoryBus,
		ppu:          ppu,
		timer:        timer,
		cartridge:    cartridge,
		Joypad:       joypad,
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

// GetVideoBuffer returns ppu video buffer.
func (gb *Gameboy) GetVideoBuffer() []uint8 {
	return gb.ppu.videoBuffer[:]
}

// step executes the next instruction in the cpu, updates the ppu and timer,
// and handles interrupts
func (gb *Gameboy) step() {
	if !gb.cpu.halted {
		gb.execute()
	} else {
		if gb.interruptBus.pending() {
			gb.cpu.halted = false
		} else {
			gb.cpu.clockCycles += 4
		}
	}

	if gb.interruptBus.ime && gb.interruptBus.pending() {
		call(gb, gb.interruptBus.locateISR())
		gb.cpu.clockCycles += isrClockCycles
	}

	if gb.interruptBus.enablingIme {
		gb.interruptBus.ime = true
		gb.interruptBus.enablingIme = false
	}

	gb.ppu.tick(gb.cpu.clockCycles)
	gb.timer.tick(gb.cpu.clockCycles)
}

// execute fetchs next opcode and executes the corresponding instruction.
func (gb *Gameboy) execute() {
	gb.cpu.clockCycles = 0
	pc := gb.cpu.ReadPc()
	opcode := gb.memoryBus.read(pc)

	if opcode == 0xcb {
		pc := gb.cpu.ReadPc()
		cbCode := gb.memoryBus.read(pc)
		cbInstr[cbCode](gb)
		gb.cpu.clockCycles += cbCycle[cbCode] * 4
	} else {
		instr[opcode](gb)
		gb.cpu.clockCycles += instrCycle[opcode] * 4
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
