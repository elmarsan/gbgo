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
)

// Bit 0: VBlank Interrupt Request (INT $40)
// Bit 1: LCD STAT Interrupt Enable (INT $48)
// Bit 2: Timer Interrupt Request (INT $50)
// Bit 3: Serial Interrupt Request (INT $58)
// Bit 4: Joypad Interrupt Request (INT $60)
var isrHandler = map[uint8]func(){
	IT_VBLANK: func() {
		ifFlag := gb.bus.read(IF)
		gb.bus.write(IF, clearBit(ifFlag, IT_VBLANK))
		gb.cpu.call(0x0040)
	},
	IT_LCD_STAT: func() {
		ifFlag := gb.bus.read(IF)
		gb.bus.write(IF, clearBit(ifFlag, IT_LCD_STAT))
		gb.cpu.call(0x0048)
	},
	IT_TIMER: func() {
		ifFlag := gb.bus.read(IF)
		gb.bus.write(IF, clearBit(ifFlag, IT_TIMER))
		gb.cpu.call(0x0050)
	},
	IT_SERIAL: func() {
		ifFlag := gb.bus.read(IF)
		gb.bus.write(IF, clearBit(ifFlag, IT_SERIAL))
		gb.cpu.call(0x0058)
	},
	IT_JOYPAD: func() {
		ifFlag := gb.bus.read(IF)
		gb.bus.write(IF, clearBit(ifFlag, IT_JOYPAD))
		gb.cpu.call(0x0060)
	},
}

// interruptPending checks whether some interrupt is pending or not.
func (gb *Gameboy) interruptPending() bool {
	ifFlag := gb.bus.read(IF)
	ieFlag := gb.bus.read(IE)

	if ifFlag&ieFlag&0x1f > 0 {
		return true
	}

	return false
}

// executeISR executes interrupt service routine.
func (gb *Gameboy) executeISR() {
	// disable interrupts
	gb.cpu.enablingIme = false
	gb.cpu.ime = false

	ifFlag := gb.bus.read(IF)

	// iterate over interrupt request flags in priority order and execute the corresponding interrupt service routine
	var i uint8 = 0
	for ; i < 5; i++ {
		if isBitSet(ifFlag, i) {
			isrHandler[i]()
			gb.cpu.clockCycles += 20
			return
		}
	}
}

// reqInterrupt request and interrupt by setting IF register.
func (gb *Gameboy) reqInterrupt(it int) {
	ifFlag := gb.bus.read(IF)

	switch it {
	case IT_VBLANK:
		gb.bus.write(IF, setBit(ifFlag, IT_VBLANK))
		break
	case IT_LCD_STAT:
		gb.bus.write(IF, setBit(ifFlag, IT_LCD_STAT))
		break
	case IT_TIMER:
		gb.bus.write(IF, setBit(ifFlag, IT_TIMER))
		break
	case IT_SERIAL:
		gb.bus.write(IF, setBit(ifFlag, IT_SERIAL))
		break
	case IT_JOYPAD:
		gb.bus.write(IF, setBit(ifFlag, IT_JOYPAD))
		break
	}
}
