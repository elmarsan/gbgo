package main

const (
	IT_VBLANK = iota
	IT_LCD_STAT
	IT_TIMER
	IT_SERIAL
	IT_JOYPAD

	// Interrupt flag
	IF = 0xff0f
	// Interrupt enable
	IE = 0xffff

	ISRClockCycles = 20
)

// InterruptBus represents the interruption system of game boy.
type InterruptBus struct {
	// bus represents memory bus used by Gameboy.
	memoryBus *MemoryBus

	// ime keeps a bool indicating whether interrupts are enabled or not.
	ime bool

	// enablingIme is set as previous step before enable/disable IME.
	enablingIme bool
}

// NewInterruptBus returns new instance of InterruptBus.
func NewInterruptBus(bus *MemoryBus) *InterruptBus {
	return &InterruptBus{
		memoryBus: bus,
	}
}

// Bit 0: VBlank Interrupt Request (INT $40)
// Bit 1: LCD STAT Interrupt Enable (INT $48)
// Bit 2: Timer Interrupt Request (INT $50)
// Bit 3: Serial Interrupt Request (INT $58)
// Bit 4: Joypad Interrupt Request (INT $60)
var isrHandler = map[uint8]func(ib *InterruptBus) uint16{
	IT_VBLANK: func(ib *InterruptBus) uint16 {
		ifFlag := ib.memoryBus.read(IF)
		ib.memoryBus.write(IF, clearBit(ifFlag, IT_VBLANK))
		return 0x0040
	},
	IT_LCD_STAT: func(ib *InterruptBus) uint16 {
		ifFlag := ib.memoryBus.read(IF)
		ib.memoryBus.write(IF, clearBit(ifFlag, IT_LCD_STAT))
		return 0x0048
	},
	IT_TIMER: func(ib *InterruptBus) uint16 {
		ifFlag := ib.memoryBus.read(IF)
		ib.memoryBus.write(IF, clearBit(ifFlag, IT_TIMER))
		return 0x0050
	},
	IT_SERIAL: func(ib *InterruptBus) uint16 {
		ifFlag := ib.memoryBus.read(IF)
		ib.memoryBus.write(IF, clearBit(ifFlag, IT_SERIAL))
		return 0x0058
	},
	IT_JOYPAD: func(ib *InterruptBus) uint16 {
		ifFlag := ib.memoryBus.read(IF)
		ib.memoryBus.write(IF, clearBit(ifFlag, IT_JOYPAD))
		return 0x0060
	},
}

// pending checks whether some interrupt is pending or not.
func (ib *InterruptBus) pending() bool {
	ifFlag := ib.memoryBus.read(IF)
	ieFlag := ib.memoryBus.read(IE)

	if ifFlag&ieFlag&0x1f > 0 {
		return true
	}

	return false
}

// locateISR returns the isr address in case some interrupt is pending.
func (ib *InterruptBus) locateISR() uint16 {
	// disable interrupts
	ib.enablingIme = false
	ib.ime = false

	ifFlag := ib.memoryBus.read(IF)

	// iterate over interrupt request flags in priority order and obtain interrupt service routine address
	var i uint8 = 0
	for ; i < 5; i++ {
		if isBitSet(ifFlag, i) {
			return isrHandler[i](ib)
		}
	}

	return 0
}

// request requests an interrupt by setting IF register.
func (ib *InterruptBus) request(it int) {
	ifFlag := ib.memoryBus.read(IF)

	switch it {
	case IT_VBLANK:
		ib.memoryBus.write(IF, setBit(ifFlag, IT_VBLANK))
		break
	case IT_LCD_STAT:
		ib.memoryBus.write(IF, setBit(ifFlag, IT_LCD_STAT))
		break
	case IT_TIMER:
		ib.memoryBus.write(IF, setBit(ifFlag, IT_TIMER))
		break
	case IT_SERIAL:
		ib.memoryBus.write(IF, setBit(ifFlag, IT_SERIAL))
		break
	case IT_JOYPAD:
		ib.memoryBus.write(IF, setBit(ifFlag, IT_JOYPAD))
		break
	}
}
