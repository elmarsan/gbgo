package main

type Gameboy struct{}

func (g *Gameboy) Run() {
	for {
		if !cpu.halted {
			cpu.execute()
		} else {
			if g.interruptPending() {
				cpu.halted = false
			} else {
				cpu.ticks += 4
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

		ppu.Tick(cpu.ticks)
		timer.update(cpu.ticks)
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
