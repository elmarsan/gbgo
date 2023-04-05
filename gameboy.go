package main

type Gameboy struct {
	running bool
}

var gameboy = &Gameboy{}

func (g *Gameboy) Init() {
	gameboy.running = false
}

func (g *Gameboy) Run() {
	g.running = true

	cpu.init()
	cpu.halted = false

	for g.running {
		if !cpu.halted {
			cpu.execute()
		}

		if cpu.enableISR {
			g.handleInterrupts()
		}

	}
}

// LoadRom load rom file into cartridge.
func (g *Gameboy) LoadRom(rom string) error {
	err := cartridge.load(rom)
	if err != nil {
		return err
	}

	return nil
}

func (g *Gameboy) handleInterrupts() {
	if cpu.enableISR {
		cpu.enableISR = false
		cpu.interruptsOn = true
		return
	}

	if !cpu.interruptsOn && !cpu.halted {
		return
	}

	interruptFlag := memory.read(0xff0f)
	interruptEnable := memory.read(0xffff)

	// Clear IF and IE
	memory.write(0xff0f, 0)
	memory.write(0xffff, 0)

	// Bit 0: VBlank   Interrupt Request (INT $40)  (1=Request)
	// Bit 1: LCD STAT Interrupt Request (INT $48)  (1=Request)
	// Bit 2: Timer    Interrupt Request (INT $50)  (1=Request)
	// Bit 3: Serial   Interrupt Request (INT $58)  (1=Request)
	// Bit 4: Joypad   Interrupt Request (INT $60)  (1=Request)
	if interruptFlag > 0 {
		switch {
		case isBitSet(interruptFlag, 0) && isBitSet(interruptEnable, 0):
			cpu.call(0x40)
			break
		case isBitSet(interruptFlag, 1) && isBitSet(interruptEnable, 1):
			cpu.call(0x48)
			break
		case isBitSet(interruptFlag, 2) && isBitSet(interruptEnable, 2):
			cpu.call(0x50)
			break
		case isBitSet(interruptFlag, 3) && isBitSet(interruptEnable, 3):
			cpu.call(0x58)
			break
		case isBitSet(interruptFlag, 4) && isBitSet(interruptEnable, 4):
			cpu.call(0x60)
			break
		}
	}

	cpu.enableISR = false
	cpu.interruptsOn = false
}
