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
		ifFlag := memory.read(IF)
		memory.write(IF, clearBit(ifFlag, IT_VBLANK))
		cpu.call(0x0040)
	},
	IT_LCD_STAT: func() {
		ifFlag := memory.read(IF)
		memory.write(IF, clearBit(ifFlag, IT_LCD_STAT))
		cpu.call(0x0048)
	},
	IT_TIMER: func() {
		ifFlag := memory.read(IF)
		memory.write(IF, clearBit(ifFlag, IT_TIMER))
		cpu.call(0x0050)
	},
	IT_SERIAL: func() {
		ifFlag := memory.read(IF)
		memory.write(IF, clearBit(ifFlag, IT_SERIAL))
		cpu.call(0x0058)
	},
	IT_JOYPAD: func() {
		ifFlag := memory.read(IF)
		memory.write(IF, clearBit(ifFlag, IT_JOYPAD))
		cpu.call(0x0060)
	},
}

func (g *Gameboy) interruptPending() bool {
	ifFlag := memory.read(IF)
	ieFlag := memory.read(IE)

	if ifFlag&ieFlag&0x1f > 0 {
		return true
	}

	return false
}

func (g *Gameboy) executeISR() {
	cpu.enablingIme = false
	cpu.ime = false

	ifFlag := memory.read(IF)

	var i uint8 = 0
	for ; i < 5; i++ {
		if isBitSet(ifFlag, i) {
			isrHandler[i]()
			return
		}
	}

	cpu.ticks += 20
}

func (g *Gameboy) reqInterrupt(it int) {
	ifFlag := memory.read(IF)

	switch it {
	case IT_VBLANK:
		memory.write(IF, setBit(ifFlag, IT_VBLANK))
		break
	case IT_LCD_STAT:
		memory.write(IF, setBit(ifFlag, IT_LCD_STAT))
		break
	case IT_TIMER:
		memory.write(IF, setBit(ifFlag, IT_TIMER))
		break
	case IT_SERIAL:
		memory.write(IF, setBit(ifFlag, IT_SERIAL))
		break
	case IT_JOYPAD:
		memory.write(IF, setBit(ifFlag, IT_JOYPAD))
		break
	}
}
