package main

var opcodes = [0x100]func(cpu *CPU){
	0x00: func(cpu *CPU) {
		// NOP
	},
	0x01: func(cpu *CPU) {
		// LD BC, d16
		val := cpu.nextPC16()
		cpu.bc.Set(val)
	},
	0x02: func(cpu *CPU) {
		// LD (BC), A
		val := cpu.af.Hi()
		cpu.bc.Set(uint16(val))
	},
	0x03: func(cpu *CPU) {
		// INC BC
		val := cpu.bc.val
		val += 1
		cpu.bc.Set(val)
	},
	0x04: func(cpu *CPU) {
		// INC B
		val := cpu.bc.Hi()
		val += 1
		cpu.bc.SetHi(val)
	},
	0x05: func(cpu *CPU) {
		// DEC B
		val := cpu.bc.Hi()
		val -= 1
		cpu.bc.SetHi(val)
	},
	0x06: func(cpu *CPU) {
		// LD B, d8
		val := cpu.bc.Hi()
		val += 1
		cpu.bc.SetHi(val)
	},
	0x07: func(cpu *CPU) {}, // TODO: RLCA
	0x08: func(cpu *CPU) {}, // TODO: LD (a16), SP
	0x09: func(cpu *CPU) {}, // TODO: ADD HL, BC
	0x0a: func(cpu *CPU) {}, // TODO: LD A, (BC)
	0x0b: func(cpu *CPU) {}, // TODO: DEC BC
	0x0c: func(cpu *CPU) {}, // TODO: INC C
	0x0d: func(cpu *CPU) {}, // TODO: DEC C
	0x0e: func(cpu *CPU) {}, // TODO: LD C, d8
	0x0f: func(cpu *CPU) {}, // TODO: RRCA

	0x10: func(cpu *CPU) {}, // TODO: STOP d8
	0x11: func(cpu *CPU) {}, // TODO: LD DE, d16
	0x12: func(cpu *CPU) {}, // TODO: LD (DE), A
	0x13: func(cpu *CPU) {}, // TODO: INC DE
	0x14: func(cpu *CPU) {}, // TODO: INC D
	0x15: func(cpu *CPU) {}, // TODO: DEC D
	0x16: func(cpu *CPU) {}, // TODO: LD D, d8
	0x17: func(cpu *CPU) {}, // TODO: RLA
	0x18: func(cpu *CPU) {}, // TODO: JR r8
	0x19: func(cpu *CPU) {}, // TODO: ADD HL, DE
	0x1a: func(cpu *CPU) {}, // TODO: LD A, (DE)
	0x1b: func(cpu *CPU) {}, // TODO: DEC DE
	0x1c: func(cpu *CPU) {}, // TODO: INC E
	0x1d: func(cpu *CPU) {}, // TODO: DEC E
	0x1e: func(cpu *CPU) {}, // TODO: LD E, d8
	0x1f: func(cpu *CPU) {}, // TODO: RRA

	0x20: func(cpu *CPU) {}, // TODO: JR NZ, r8
	0x21: func(cpu *CPU) {}, // TODO: LD HL, d16
	0x22: func(cpu *CPU) {}, // TODO: LD (HL+), A
	0x23: func(cpu *CPU) {}, // TODO: INC HL
	0x24: func(cpu *CPU) {}, // TODO: INC H
	0x25: func(cpu *CPU) {}, // TODO: DEC H
	0x26: func(cpu *CPU) {}, // TODO: LD H, d8
	0x27: func(cpu *CPU) {}, // TODO: DAA
	0x28: func(cpu *CPU) {}, // TODO: JR Z, r8
	0x29: func(cpu *CPU) {}, // TODO: ADD HL, HL
	0x2a: func(cpu *CPU) {}, // TODO: LD A, (HL+)
	0x2b: func(cpu *CPU) {}, // TODO: DEC HL
	0x2c: func(cpu *CPU) {}, // TODO: INC L
	0x2d: func(cpu *CPU) {}, // TODO: DEC L
	0x2e: func(cpu *CPU) {}, // TODO: LD L, d8
	0x2f: func(cpu *CPU) {}, // TODO: CPL

	0x30: func(cpu *CPU) {}, // TODO: JR NC, r8
	0x31: func(cpu *CPU) {}, // TODO: LD SP, d16
	0x32: func(cpu *CPU) {}, // TODO: LD (HL-), A
	0x33: func(cpu *CPU) {}, // TODO: INC SP
	0x34: func(cpu *CPU) {}, // TODO: INC (HL)
	0x35: func(cpu *CPU) {}, // TODO: DEC (HL)
	0x36: func(cpu *CPU) {}, // TODO: LD (HL), d8
	0x37: func(cpu *CPU) {}, // TODO: SCF
	0x38: func(cpu *CPU) {}, // TODO: JR C, r8
	0x39: func(cpu *CPU) {}, // TODO: ADD HL, SP
	0x3a: func(cpu *CPU) {}, // TODO: LD A, (HL-)
	0x3b: func(cpu *CPU) {}, // TODO: DEC SP
	0x3c: func(cpu *CPU) {}, // TODO: INC A
	0x3d: func(cpu *CPU) {}, // TODO: DEC A
	0x3e: func(cpu *CPU) {}, // TODO: LD A, d8
	0x3f: func(cpu *CPU) {}, // TODO: CCF
}
