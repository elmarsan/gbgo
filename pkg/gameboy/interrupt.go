package gameboy

import "github.com/elmarsan/gbgo/pkg/bit"

const (
	vblankInterrupt = iota
	lcdStatInterrupt
	timerInterrupt
	serialInterrupt
	joypadInterrupt

	// Interrupt flag
	iFF = 0xff0f
	// Interrupt enable
	ie = 0xffff

	isrClockCycles = 20
)

// interruptBus represents the interruption system of game boy.
type interruptBus struct {
	// bus represents memory bus used by Gameboy.
	memoryBus *memoryBus

	// ime keeps a bool indicating whether interrupts are enabled or not.
	ime bool

	// enablingIme is set as previous step before enable/disable IME.
	enablingIme bool
}

// newInterruptBus returns new instance of InterruptBus.
func newInterruptBus(bus *memoryBus) *interruptBus {
	return &interruptBus{
		memoryBus: bus,
	}
}

// Bit 0: VBlank Interrupt Request (INT $40)
// Bit 1: LCD STAT Interrupt Enable (INT $48)
// Bit 2: Timer Interrupt Request (INT $50)
// Bit 3: Serial Interrupt Request (INT $58)
// Bit 4: Joypad Interrupt Request (INT $60)
var isrHandler = map[uint8]func(ib *interruptBus) uint16{
	vblankInterrupt: func(ib *interruptBus) uint16 {
		ifFlag := ib.memoryBus.read(iFF)
		ib.memoryBus.write(iFF, bit.Clear(ifFlag, vblankInterrupt))
		return 0x0040
	},
	lcdStatInterrupt: func(ib *interruptBus) uint16 {
		ifFlag := ib.memoryBus.read(iFF)
		ib.memoryBus.write(iFF, bit.Clear(ifFlag, lcdStatInterrupt))
		return 0x0048
	},
	timerInterrupt: func(ib *interruptBus) uint16 {
		ifFlag := ib.memoryBus.read(iFF)
		ib.memoryBus.write(iFF, bit.Clear(ifFlag, timerInterrupt))
		return 0x0050
	},
	serialInterrupt: func(ib *interruptBus) uint16 {
		ifFlag := ib.memoryBus.read(iFF)
		ib.memoryBus.write(iFF, bit.Clear(ifFlag, serialInterrupt))
		return 0x0058
	},
	joypadInterrupt: func(ib *interruptBus) uint16 {
		ifFlag := ib.memoryBus.read(iFF)
		ib.memoryBus.write(iFF, bit.Clear(ifFlag, joypadInterrupt))
		return 0x0060
	},
}

// pending checks whether some interrupt is pending or not.
func (ib *interruptBus) pending() bool {
	ifFlag := ib.memoryBus.read(iFF)
	ieFlag := ib.memoryBus.read(ie)

	if ifFlag&ieFlag&0x1f > 0 {
		return true
	}

	return false
}

// locateISR returns the isr address in case some interrupt is pending.
func (ib *interruptBus) locateISR() uint16 {
	// disable interrupts
	ib.enablingIme = false
	ib.ime = false

	ifFlag := ib.memoryBus.read(iFF)

	// iterate over interrupt request flags in priority order and obtain interrupt service routine address
	var i uint8 = 0
	for ; i < 5; i++ {
		if bit.IsSet(ifFlag, i) {
			return isrHandler[i](ib)
		}
	}

	return 0
}

// request requests an interrupt by setting IF register.
func (ib *interruptBus) request(it int) {
	ifFlag := ib.memoryBus.read(iFF)

	switch it {
	case vblankInterrupt:
		ib.memoryBus.write(iFF, bit.Set(ifFlag, vblankInterrupt))
		break
	case lcdStatInterrupt:
		ib.memoryBus.write(iFF, bit.Set(ifFlag, lcdStatInterrupt))
		break
	case timerInterrupt:
		ib.memoryBus.write(iFF, bit.Set(ifFlag, timerInterrupt))
		break
	case serialInterrupt:
		ib.memoryBus.write(iFF, bit.Set(ifFlag, serialInterrupt))
		break
	case joypadInterrupt:
		ib.memoryBus.write(iFF, bit.Set(ifFlag, joypadInterrupt))
		break
	}
}
