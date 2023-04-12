package main

type Gameboy struct{}

func (g *Gameboy) Run() {
	for {
		if !cpu.halted {
			cpu.execute()
		} else {
			if g.interruptPending() {
				cpu.halted = false
			}
		}

		if cpu.ime {
			if g.interruptPending() {
				g.executeISR()
			}
		}

		if cpu.enablingIme {
			cpu.ime = true
			cpu.enablingIme = false
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

func (g *Gameboy) interruptPending() bool {
	ifFlag := memory.read(0xff0f)
	ieFlag := memory.read(0xffff)

	if ifFlag&ieFlag&0x1f > 0 {
		return true
	}

	return false
}

func (g *Gameboy) executeISR() {
	cpu.enablingIme = false
	cpu.ime = false

	ifFlag := memory.read(0xff0f)

	var i uint8 = 0
	for ; i < 5; i++ {
		if isBitSet(ifFlag, i) {
			isrHandler[i]()
			return
		}
	}
}

// Bit 0: VBlank Interrupt Request (INT $40)
// Bit 1: LCD STAT Interrupt Enable (INT $48)
// Bit 2: Timer Interrupt Request (INT $50)
// Bit 3: Serial Interrupt Request (INT $58)
// Bit 4: Joypad Interrupt Request (INT $60)
var isrHandler = map[uint8]func(){
	0: func() {
		ifFlag := memory.read(0xff0f)
		memory.write(0xff0f, clearBit(ifFlag, 0))
		cpu.call(uint16(memory.read(0x40)))
	},
	1: func() {
		ifFlag := memory.read(0xff0f)
		memory.write(0xff0f, clearBit(ifFlag, 1))
		cpu.call(uint16(memory.read(0x48)))
	},
	2: func() {
		ifFlag := memory.read(0xff0f)
		memory.write(0xff0f, clearBit(ifFlag, 2))
		cpu.call(uint16(memory.read(0x50)))
	},
	3: func() {
		ifFlag := memory.read(0xff0f)
		memory.write(0xff0f, clearBit(ifFlag, 3))
		cpu.call(uint16(memory.read(0x58)))
	},
	4: func() {
		ifFlag := memory.read(0xff0f)
		memory.write(0xff0f, clearBit(ifFlag, 4))
		cpu.call(uint16(memory.read(0x60)))
	},
}
