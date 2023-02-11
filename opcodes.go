package main

var opcodes = [0x100]func(gb *Gameboy){
	0x00: func(gb *Gameboy) {
		// NOP
	},
	0x01: func(gb *Gameboy) {
		// LD BC, d16
		val := gb.cpu.nextPC16()
		gb.cpu.bc.Set(val)
	},
	0x02: func(gb *Gameboy) {
		// LD (BC), A
		val := gb.cpu.af.Hi()
		gb.cpu.bc.Set(uint16(val))
	},
	0x03: func(gb *Gameboy) {
		// INC BC
		val := gb.cpu.bc.val
		val += 1
		gb.cpu.bc.Set(val)
	},
	0x04: func(gb *Gameboy) {
		// INC B
		val := gb.cpu.bc.Hi()
		val += 1
		gb.cpu.bc.SetHi(val)
	},
	0x05: func(gb *Gameboy) {
		// DEC B
		val := gb.cpu.bc.Hi()
		val -= 1
		gb.cpu.bc.SetHi(val)
	},
	0x06: func(gb *Gameboy) {
		// LD B, d8
		val := gb.cpu.bc.Hi()
		val += 1
		gb.cpu.bc.SetHi(val)
	},
	0x07: func(gb *Gameboy) {}, // TODO: RLCA
	0x08: func(gb *Gameboy) {}, // TODO: LD (a16), SP
	0x09: func(gb *Gameboy) {}, // TODO: ADD HL, BC
	0x0a: func(gb *Gameboy) {}, // TODO: LD A, (BC)
	0x0b: func(gb *Gameboy) {}, // TODO: DEC BC
	0x0c: func(gb *Gameboy) {}, // TODO: INC C
	0x0d: func(gb *Gameboy) {}, // TODO: DEC C
	0x0e: func(gb *Gameboy) {}, // TODO: LD C, d8
	0x0f: func(gb *Gameboy) {}, // TODO: RRCA

	0x10: func(gb *Gameboy) {}, // TODO: STOP d8
	0x11: func(gb *Gameboy) {}, // TODO: LD DE, d16
	0x12: func(gb *Gameboy) {}, // TODO: LD (DE), A
	0x13: func(gb *Gameboy) {}, // TODO: INC DE
	0x14: func(gb *Gameboy) {}, // TODO: INC D
	0x15: func(gb *Gameboy) {}, // TODO: DEC D
	0x16: func(gb *Gameboy) {}, // TODO: LD D, d8
	0x17: func(gb *Gameboy) {}, // TODO: RLA
	0x18: func(gb *Gameboy) {}, // TODO: JR r8
	0x19: func(gb *Gameboy) {}, // TODO: ADD HL, DE
	0x1a: func(gb *Gameboy) {}, // TODO: LD A, (DE)
	0x1b: func(gb *Gameboy) {}, // TODO: DEC DE
	0x1c: func(gb *Gameboy) {}, // TODO: INC E
	0x1d: func(gb *Gameboy) {}, // TODO: DEC E
	0x1e: func(gb *Gameboy) {}, // TODO: LD E, d8
	0x1f: func(gb *Gameboy) {}, // TODO: RRA

	0x20: func(gb *Gameboy) {}, // TODO: JR NZ, r8
	0x21: func(gb *Gameboy) {}, // TODO: LD HL, d16
	0x22: func(gb *Gameboy) {}, // TODO: LD (HL+), A
	0x23: func(gb *Gameboy) {}, // TODO: INC HL
	0x24: func(gb *Gameboy) {}, // TODO: INC H
	0x25: func(gb *Gameboy) {}, // TODO: DEC H
	0x26: func(gb *Gameboy) {}, // TODO: LD H, d8
	0x27: func(gb *Gameboy) {}, // TODO: DAA
	0x28: func(gb *Gameboy) {}, // TODO: JR Z, r8
	0x29: func(gb *Gameboy) {}, // TODO: ADD HL, HL
	0x2a: func(gb *Gameboy) {}, // TODO: LD A, (HL+)
	0x2b: func(gb *Gameboy) {}, // TODO: DEC HL
	0x2c: func(gb *Gameboy) {}, // TODO: INC L
	0x2d: func(gb *Gameboy) {}, // TODO: DEC L
	0x2e: func(gb *Gameboy) {}, // TODO: LD L, d8
	0x2f: func(gb *Gameboy) {}, // TODO: CPL

	0x30: func(gb *Gameboy) {}, // TODO: JR NC, r8
	0x31: func(gb *Gameboy) {}, // TODO: LD SP, d16
	0x32: func(gb *Gameboy) {}, // TODO: LD (HL-), A
	0x33: func(gb *Gameboy) {}, // TODO: INC SP
	0x34: func(gb *Gameboy) {}, // TODO: INC (HL)
	0x35: func(gb *Gameboy) {}, // TODO: DEC (HL)
	0x36: func(gb *Gameboy) {}, // TODO: LD (HL), d8
	0x37: func(gb *Gameboy) {}, // TODO: SCF
	0x38: func(gb *Gameboy) {}, // TODO: JR C, r8
	0x39: func(gb *Gameboy) {}, // TODO: ADD HL, SP
	0x3a: func(gb *Gameboy) {}, // TODO: LD A, (HL-)
	0x3b: func(gb *Gameboy) {}, // TODO: DEC SP
	0x3c: func(gb *Gameboy) {}, // TODO: INC A
	0x3d: func(gb *Gameboy) {}, // TODO: DEC A
	0x3e: func(gb *Gameboy) {}, // TODO: LD A, d8
	0x3f: func(gb *Gameboy) {}, // TODO: CCF

	// LD B, C
	0x41: func(gb *Gameboy) {
		val := gb.cpu.bc.Low()
		gb.cpu.bc.SetHi(val)
	},
	// LD B, D
	0x42: func(gb *Gameboy) {
		val := gb.cpu.bc.Low()
		gb.cpu.de.SetHi(val)
	},
	// LD B, E
	0x43: func(gb *Gameboy) {
		val := gb.cpu.bc.Low()
		gb.cpu.de.SetLow(val)
	},
	// LD B, H
	0x44: func(gb *Gameboy) {
		val := gb.cpu.bc.Low()
		gb.cpu.hl.SetHi(val)
	},
	0x45: func(gb *Gameboy) {
		// LD B, L
		val := gb.cpu.bc.Low()
		gb.cpu.hl.SetLow(val)
	},
	0x46: func(gb *Gameboy) {
		// LD B, HL
		addr := gb.cpu.hl.val
		val, _ := gb.memory.read(addr)
		gb.cpu.bc.Set(val)
	},
	0x47: func(gb *Gameboy) {

	},
	0x48: func(gb *Gameboy) {

	},
	0x49: func(gb *Gameboy) {

	},
	0x4a: func(gb *Gameboy) {

	},
	0x4b: func(gb *Gameboy) {

	},
	0x4c: func(gb *Gameboy) {

	},
	0x4d: func(gb *Gameboy) {

	},
	0x4e: func(gb *Gameboy) {

	},
	0x4f: func(gb *Gameboy) {

	},
}
