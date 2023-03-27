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

		g.handleInterrupts()
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
	if cpu.setIme {
		cpu.ime = true
		return
	}

	if !cpu.ime && cpu.halted {
		return
	}

	interruptFlag := memory.read(0xff0f)

	if interruptFlag > 0 && cpu.ime {
		for i := 0; i < 5; i++ {
			if isBitSet(interruptFlag, uint8(i)) {
				g.serveInterrupt(uint8(i))
				return
			}
		}
	}
}

// Bit 0: VBlank   Interrupt Request (INT $40)  (1=Request)
// Bit 1: LCD STAT Interrupt Request (INT $48)  (1=Request)
// Bit 2: Timer    Interrupt Request (INT $50)  (1=Request)
// Bit 3: Serial   Interrupt Request (INT $58)  (1=Request)
// Bit 4: Joypad   Interrupt Request (INT $60)  (1=Request)
func (g *Gameboy) serveInterrupt(bit uint8) {
	cpu.ime = false
	cpu.halted = false

	switch bit {
	case 0:
		cpu.call(0x40)
		break
	case 1:
		cpu.call(0x48)
		break
	case 2:
		cpu.call(0x50)
		break
	case 3:
		cpu.call(0x58)
		break
	case 4:
		cpu.call(0x60)
		break
	}
}
