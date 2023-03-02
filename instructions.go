package main

var instructions = [0x100]func(){
	0x00: func() {}, // NOP
	0x01: func() {
		// LD BC, d16
		lsb := lo(cpu.readPc())
		msb := hi(cpu.readPc())
		cpu.set16Reg(REG_BC, joinUint8(msb, lsb))
	},
	0x02: func() {
		// LD (BC), A
		addr := cpu.read16Reg(REG_BC)
		memory.write(addr, cpu.read8Reg(REG_A))
	},
	0x03: func() {
		// INC BC
		cpu.inc16reg(REG_BC, false)
	},
	0x04: func() {
		// INC B
		cpu.inc8Reg(REG_B)
	},
	0x05: func() {
		// DEC B
		cpu.dec8Reg(REG_B)
	},
	0x06: func() {
		// LD B, d8
		val := memory.read(cpu.readPc())
		cpu.set8Reg(REG_B, val)
	},
	0x07: func() {}, // RLCA
	0x08: func() {
		// LD (a16), SP
		lsb := lo(cpu.readPc())
		msb := hi(cpu.readPc())
		addr := joinUint8(msb, lsb)

		memory.write(addr, lo(cpu.sp))
		memory.write(addr+1, hi(cpu.sp))
	},
	0x09: func() {
		// ADD HL, BC
		cpu.add16Reg(REG_HL, REG_BC)
	},
	0x0a: func() {
		// LD A, (BC)
		val := memory.read(cpu.read16Reg(REG_BC))
		cpu.set8Reg(REG_A, val)
	},
	0x0b: func() {
		// DEC BC
		cpu.dec16Reg(REG_BC)
	},
	0x0c: func() {
		// INC C
		cpu.inc8Reg(REG_C)
	},

	0x0d: func() {
		// DEC C
		cpu.dec8Reg(REG_D)
	},
	0x0e: func() {
		// LD C, d8
		cpu.set8Reg(REG_C, memory.read(cpu.readPc()))
	},
	0x0f: func() {}, // RRCA

	0x10: func() {
		// STOP d8
	},
	0x11: func() {}, // LD DE, d16
	0x12: func() {
		// LD (DE), A
		addr := cpu.read16Reg(REG_DE)
		val := cpu.read8Reg(REG_A)
		memory.write(addr, val)
	},
	0x13: func() {
		// INC DE
		cpu.inc16reg(REG_DE, false)
	},
	0x14: func() {
		// INC D
		cpu.inc8Reg(REG_D)
	},
	0x15: func() {
		// DEC D
		cpu.dec8Reg(REG_D)
	},
	0x16: func() {
		// LD D, d8
		cpu.set8Reg(REG_D, memory.read(cpu.readPc()))
	},
	0x17: func() {}, // RLA
	0x18: func() {}, // JR r8
	0x19: func() {
		// ADD HL, DE
		cpu.add16Reg(REG_HL, REG_DE)
	},
	0x1a: func() {}, // LD A, (DE)
	0x1b: func() {
		// DEC DE
		cpu.dec16Reg(REG_DE)
	},
	0x1c: func() {
		// INC E
		cpu.inc8Reg(REG_E)
	},
	0x1d: func() {
		// DEC E
		cpu.dec8Reg(REG_E)
	},
	0x1e: func() {}, // LD E, d8
	0x1f: func() {}, // RRA

	0x20: func() {}, // JR NZ, r8
	0x21: func() {}, // LD HL, d16
	0x22: func() {
		// LD (HL+), A
		addr := cpu.read16Reg(REG_HL)
		val := cpu.read8Reg(REG_A)
		memory.write(addr, val)
		cpu.set16Reg(REG_HL, addr+1)
	},
	0x23: func() {
		// INC HL
		cpu.inc16reg(REG_HL, false)
	},
	0x24: func() {
		// INC H
		cpu.inc8Reg(REG_H)
	},
	0x25: func() {
		// DEC H
		cpu.dec8Reg(REG_H)
	},
	0x26: func() {
		// LD H, d8
		cpu.set8Reg(REG_H, memory.read(cpu.readPc()))
	},
	0x27: func() {}, // DAA
	0x28: func() {}, // JR Z, r8
	0x29: func() {
		// ADD HL, HL
		cpu.add16Reg(REG_HL, REG_HL)
	},
	0x2a: func() {
		// LD A, (HL+)
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.set8Reg(REG_A, val)
		cpu.set16Reg(REG_HL, cpu.read16Reg(REG_HL)+1)
	},
	0x2b: func() {
		// DEC HL
		cpu.dec16Reg(REG_HL)
	},
	0x2c: func() {
		// INC L
		cpu.inc8Reg(REG_L)
	},
	0x2d: func() {
		// DEC L
		cpu.dec8Reg(REG_L)
	},
	0x2e: func() {}, // LD L, d8
	0x2f: func() {
		// CPL
		val := cpu.read8Reg(REG_A)
		cpu.set8Reg(REG_A, 0xff^val)
		cpu.nFlag = true
		cpu.hFlag = true
	},

	0x30: func() {}, // JR NC, r8
	0x31: func() {}, // LD SP, d16
	0x32: func() {
		// LD (HL-), A
		addr := cpu.read16Reg(REG_HL)
		val := cpu.read8Reg(REG_A)
		memory.write(addr, val)
		cpu.set16Reg(REG_HL, addr-1)
	},
	0x33: func() {
		// INC SP
		cpu.sp = cpu.sp + 1
	},
	0x34: func() {
		// INC (HL)
		cpu.inc16reg(REG_HL, true)
	},
	0x35: func() {
		// DEC (HL)
	},
	0x36: func() {
		// LD (HL), d8
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(cpu.readPc())
		memory.write(addr, val)
	},
	0x37: func() {}, // SCF
	0x38: func() {}, // JR C, r8
	0x39: func() {
		// ADD HL, SP
		cpu.add16Reg(REG_HL, REG_SP)
	},
	0x3a: func() {}, // LD A, (HL-)
	0x3b: func() {
		// DEC SP
		cpu.sp = cpu.sp - 1
	},
	0x3c: func() {
		// INC A
		cpu.inc8Reg(REG_A)
	},
	0x3d: func() {
		// DEC A
		cpu.dec8Reg(REG_A)
	},
	0x3e: func() {
		// LD A, d8
	},
	0x3f: func() {
		// CCF
		cpu.nFlag = false
		cpu.hFlag = false
		cpu.cFlag = !cpu.cFlag
	},

	0x40: func() {
		// LD B, B
		cpu.load8Reg(REG_B, REG_B)
	},
	0x41: func() {
		// LD B, C
		cpu.load8Reg(REG_B, REG_C)
	},
	0x42: func() {
		// LD B, D
		cpu.load8Reg(REG_B, REG_D)
	},
	0x43: func() {
		// LD B, E
		cpu.load8Reg(REG_B, REG_E)
	},
	0x44: func() {
		// LD B, H
		cpu.load8Reg(REG_B, REG_H)
	},
	0x45: func() {
		// LD B, L
		cpu.load8Reg(REG_B, REG_L)
	},
	0x46: func() {
		// LD B, HL
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.set8Reg(REG_B, val)
	},
	0x47: func() {
		// LD B, A
		cpu.load8Reg(REG_B, REG_A)
	},
	0x48: func() {
		// LD C, B
		cpu.load8Reg(REG_C, REG_B)
	},
	0x49: func() {
		// LD C, C
		cpu.load8Reg(REG_C, REG_C)
	},
	0x4a: func() {
		// LD C, D
		cpu.load8Reg(REG_C, REG_D)
	},
	0x4b: func() {
		// LD C, E
		cpu.load8Reg(REG_C, REG_E)
	},
	0x4c: func() {
		// LD C, H
		cpu.load8Reg(REG_C, REG_H)
	},
	0x4d: func() {
		// LD C, L
		cpu.load8Reg(REG_C, REG_L)
	},
	0x4e: func() {
		// LD C, (HL)
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.set8Reg(REG_C, val)
	},
	0x4f: func() {
		// LD C, A
		cpu.load8Reg(REG_C, REG_A)
	},

	0x50: func() {
		// LD D, B
		cpu.load8Reg(REG_D, REG_B)
	},
	0x51: func() {
		// LD D, C
		cpu.load8Reg(REG_D, REG_C)
	},
	0x52: func() {
		// LD D, D
		cpu.load8Reg(REG_D, REG_D)
	},
	0x53: func() {
		// LD D, E
		cpu.load8Reg(REG_D, REG_E)
	},
	0x54: func() {
		// LD D, H
		cpu.load8Reg(REG_D, REG_H)
	},
	0x55: func() {
		// LD D, L
		cpu.load8Reg(REG_D, REG_L)
	},
	0x56: func() {
		// LD D, HL
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.set8Reg(REG_D, val)
	},
	0x57: func() {
		// LD D, A
		cpu.load8Reg(REG_D, REG_A)
	},
	0x58: func() {
		// LD E, B
		cpu.load8Reg(REG_E, REG_B)
	},
	0x59: func() {
		// LD E, C
		cpu.load8Reg(REG_E, REG_C)
	},
	0x5a: func() {
		// LD E, D
		cpu.load8Reg(REG_E, REG_D)
	},
	0x5b: func() {
		// LD E, E
		cpu.load8Reg(REG_E, REG_E)
	},
	0x5c: func() {
		// LD E, H
		cpu.load8Reg(REG_E, REG_H)
	},
	0x5d: func() {
		// LD E, L
		cpu.load8Reg(REG_E, REG_L)
	},
	0x5e: func() {
		// LD E, HL
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.set8Reg(REG_E, val)
	},
	0x5f: func() {
		// LD E, A
		cpu.load8Reg(REG_E, REG_A)
	},

	0x60: func() {
		// LD H, B
		cpu.load8Reg(REG_H, REG_B)
	},
	0x61: func() {
		// LD H, C
		cpu.load8Reg(REG_H, REG_C)
	},
	0x62: func() {
		// LD H, D
		cpu.load8Reg(REG_H, REG_D)
	},
	0x63: func() {
		// LD H, E
		cpu.load8Reg(REG_H, REG_E)
	},
	0x64: func() {
		// LD H, H
		cpu.load8Reg(REG_H, REG_H)
	},
	0x65: func() {
		// LD H, L
		cpu.load8Reg(REG_H, REG_L)
	},
	0x66: func() {
		// LD H, (HL)
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.set8Reg(REG_H, val)
	},
	0x67: func() {
		// LD H, A
		cpu.load8Reg(REG_H, REG_A)
	},
	0x68: func() {
		// LD L, B
		cpu.load8Reg(REG_L, REG_B)
	},
	0x69: func() {
		// LD L, C
		cpu.load8Reg(REG_L, REG_C)
	},
	0x6a: func() {
		// LD L, D
		cpu.load8Reg(REG_L, REG_D)
	},
	0x6b: func() {
		// LD L, E
		cpu.load8Reg(REG_L, REG_E)
	},
	0x6c: func() {
		// LD L, H
		cpu.load8Reg(REG_L, REG_H)
	},
	0x6d: func() {
		// LD L, L
		cpu.load8Reg(REG_L, REG_L)
	},
	0x6e: func() {
		// LD L, (HL)
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.set8Reg(REG_L, val)
	},
	0x6f: func() {
		// LD L, A
		cpu.load8Reg(REG_L, REG_A)
	},

	0x70: func() {
		// LD (HL), B
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_B))
	},
	0x71: func() {
		// LD (HL), C
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_C))
	},
	0x72: func() {
		// LD (HL), D
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_D))
	},
	0x73: func() {
		// LD (HL), E
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_E))
	},
	0x74: func() {
		// LD (HL), H
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_H))
	},
	0x75: func() {
		// LD (HL), L
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_L))
	},
	0x76: func() {
		// HALT
	},
	0x77: func() {
		// LD (HL), A
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_A))
	},
	0x78: func() {
		// LD A, B
		cpu.load8Reg(REG_A, REG_B)
	},
	0x79: func() {
		// LD A, C
		cpu.load8Reg(REG_A, REG_C)
	},
	0x7a: func() {
		// LD A, D
		cpu.load8Reg(REG_A, REG_D)
	},
	0x7b: func() {
		// LD A, E
		cpu.load8Reg(REG_A, REG_E)
	},
	0x7c: func() {
		// LD A, H
		cpu.load8Reg(REG_A, REG_H)
	},
	0x7d: func() {
		// LD A, L
		cpu.load8Reg(REG_A, REG_L)
	},
	0x7e: func() {
		// LD A, (HL)
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.set8Reg(REG_A, val)
	},
	0x7f: func() {
		// LD A, A
		cpu.load8Reg(REG_A, REG_A)
	},

	0x80: func() {
		// ADD A, B
		cpu.add8Reg(REG_A, REG_B)
	},
	0x81: func() {
		// ADD A, C
		cpu.add8Reg(REG_A, REG_C)
	},
	0x82: func() {
		// ADD A, D
		cpu.add8Reg(REG_A, REG_D)
	},
	0x83: func() {
		// ADD A, E
		cpu.add8Reg(REG_A, REG_E)
	},
	0x84: func() {
		// ADD A, H
		cpu.add8Reg(REG_A, REG_H)
	},
	0x85: func() {
		// ADD A, L
		cpu.add8Reg(REG_A, REG_L)
	},
	0x86: func() {
		// ADD A, HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.add8RegD8(REG_A, val)
	},
	0x87: func() {
		// ADD A, A
		cpu.add8Reg(REG_A, REG_A)
	},
	0x88: func() {
		// ADC A, B
		cpu.adc8Reg(REG_A, REG_B)
	},
	0x89: func() {
		// ADC A, C
		cpu.adc8Reg(REG_A, REG_C)
	},
	0x8a: func() {
		// ADC A, D
		cpu.adc8Reg(REG_A, REG_D)
	},
	0x8b: func() {
		// ADC A, E
		cpu.adc8Reg(REG_A, REG_E)
	},
	0x8c: func() {
		// ADC A, H
		cpu.adc8Reg(REG_A, REG_H)
	},
	0x8d: func() {
		// ADC A, L
		cpu.adc8Reg(REG_A, REG_L)
	},
	0x8e: func() {
		// ADC A, HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.adc8RegD8(REG_A, val)
	},
	0x8f: func() {
		// ADC A, A
		cpu.adc8Reg(REG_A, REG_A)
	},

	0x90: func() {
		// SUB B
		cpu.sub8Reg(REG_A, REG_B)
	},
	0x91: func() {
		// SUB C
		cpu.sub8Reg(REG_A, REG_C)
	},
	0x92: func() {
		// SUB D
		cpu.sub8Reg(REG_A, REG_D)
	},
	0x93: func() {
		// SUB E
		cpu.sub8Reg(REG_A, REG_E)
	},
	0x94: func() {
		// SUB H
		cpu.sub8Reg(REG_A, REG_H)
	},
	0x95: func() {
		// SUB L
		cpu.sub8Reg(REG_A, REG_L)
	},
	0x96: func() {
		// SUB (HL)
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.sub8RegD8(REG_A, val)
	},
	0x97: func() {
		// SUB A
		cpu.sub8Reg(REG_A, REG_A)
	},
	0x98: func() {
		// SBC A, B
		cpu.sbc8Reg(REG_A, REG_B)
	},
	0x99: func() {
		// SBC A, C
		cpu.sbc8Reg(REG_A, REG_C)
	},
	0x9a: func() {
		// SBC A, D
		cpu.sbc8Reg(REG_A, REG_D)
	},
	0x9b: func() {
		// SBC A, E
		cpu.sbc8Reg(REG_A, REG_E)
	},
	0x9c: func() {
		// SBC A, H
		cpu.sbc8Reg(REG_A, REG_H)
	},
	0x9d: func() {
		// SBC A, L
		cpu.sbc8Reg(REG_A, REG_L)
	},
	0x9e: func() {
		// SBC A, (HL)
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.sbc8RegD8(REG_A, val)
	},
	0x9f: func() {
		// SBC A, A
		cpu.sbc8Reg(REG_A, REG_A)
	},

	0xa0: func() {
		// AND B
		cpu.and8Reg(REG_A, REG_B)
	},
	0xa1: func() {
		// AND C
		cpu.and8Reg(REG_A, REG_C)
	},
	0xa2: func() {
		// AND D
		cpu.and8Reg(REG_A, REG_D)
	},
	0xa3: func() {
		// AND E
		cpu.and8Reg(REG_A, REG_E)
	},
	0xa4: func() {
		// AND H
		cpu.and8Reg(REG_A, REG_H)
	},
	0xa5: func() {
		// AND L
		cpu.and8Reg(REG_A, REG_L)
	},
	0xa6: func() {
		// AND HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.and8RegD8(REG_A, val)
	},
	0xa7: func() {
		// AND A
		cpu.and8Reg(REG_A, REG_A)
	},
	0xa8: func() {
		// XOR B
		cpu.xor8Reg(REG_A, REG_B)
	},
	0xa9: func() {
		// XOR C
		cpu.xor8Reg(REG_A, REG_C)
	},
	0xaa: func() {
		// XOR D
		cpu.xor8Reg(REG_A, REG_D)
	},
	0xab: func() {
		// XOR E
		cpu.xor8Reg(REG_A, REG_E)
	},
	0xac: func() {
		// XOR H
		cpu.xor8Reg(REG_A, REG_H)
	},
	0xad: func() {
		// XOR L
		cpu.xor8Reg(REG_A, REG_L)
	},
	0xae: func() {
		// XOR HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.xor8RegD8(REG_A, val)
	},
	0xaf: func() {
		// XOR A
		cpu.xor8Reg(REG_A, REG_A)
	},

	0xb0: func() {
		// OR B
		cpu.or8Reg(REG_A, REG_B)
	},
	0xb1: func() {
		// OR C
		cpu.or8Reg(REG_A, REG_C)
	},
	0xb2: func() {
		// OR D
		cpu.or8Reg(REG_A, REG_D)
	},
	0xb3: func() {
		// OR E
		cpu.or8Reg(REG_A, REG_E)
	},
	0xb4: func() {
		// OR H
		cpu.or8Reg(REG_A, REG_H)
	},
	0xb5: func() {
		// OR L
		cpu.or8Reg(REG_A, REG_L)
	},
	0xb6: func() {
		// OR HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.or8RegD8(REG_A, val)
	},
	0xb7: func() {
		// OR A
		cpu.or8Reg(REG_A, REG_A)
	},
	0xb8: func() {
		// CP B
		cpu.cp8Reg(REG_A, REG_B)
	},
	0xb9: func() {
		// CP C
		cpu.cp8Reg(REG_A, REG_C)
	},
	0xba: func() {
		// CP D
		cpu.cp8Reg(REG_A, REG_D)
	},
	0xbb: func() {
		// CP E
		cpu.cp8Reg(REG_A, REG_E)
	},
	0xbc: func() {
		// CP H
		cpu.cp8Reg(REG_A, REG_H)
	},
	0xbd: func() {
		// CP L
		cpu.cp8Reg(REG_A, REG_L)
	},
	0xbe: func() {
		// CP (HL)
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.cp8RegD8(REG_A, val)
	},
	0xbf: func() {
		// CP A
		cpu.cp8Reg(REG_A, REG_A)
	},

	0xc0: func() {},
	0xc1: func() {},
	0xc2: func() {},
	0xc3: func() {},
	0xc4: func() {},
	0xc5: func() {},
	0xc6: func() {},
	0xc7: func() {},
	0xc8: func() {},
	0xc9: func() {},
	0xca: func() {},
	0xcb: func() {},
	0xcc: func() {},
	0xcd: func() {},
	0xce: func() {
		// ADC A, d8
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.adc8RegD8(REG_A, val)
	},
	0xcf: func() {},

	0xd0: func() {},
	0xd1: func() {},
	0xd2: func() {},
	0xd3: func() {},
	0xd4: func() {},
	0xd5: func() {},
	0xd6: func() {
		// SUB d8
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.sub8RegD8(REG_A, val)
	},
	0xd7: func() {},
	0xd8: func() {},
	0xd9: func() {},
	0xda: func() {},
	0xdb: func() {},
	0xdc: func() {},
	0xdd: func() {},
	0xde: func() {
		// SBC A, d8
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.sbc8RegD8(REG_A, val)
	},
	0xdf: func() {},

	0xe0: func() {
		// LDH (a8), A
		lsb := lo(cpu.readPc())
		addr := joinUint8(0xff, lsb)
		memory.write(addr, cpu.read8Reg(REG_A))
	},
	0xe1: func() {},
	0xe2: func() {},
	0xe3: func() {},
	0xe4: func() {},
	0xe5: func() {},
	0xe6: func() {
		// AND d8
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.and8RegD8(REG_A, val)
	},
	0xe7: func() {},
	0xe8: func() {
		// ADD SP, r8
		val := cpu.readPc()
		cpu.add16RegD8(REG_SP, uint8(val))
	},
	0xe9: func() {},
	0xea: func() {
		// LD (a16), A
		lsb := lo(cpu.readPc())
		msb := hi(cpu.readPc())
		addr := joinUint8(msb, lsb)
		memory.write(addr, cpu.read8Reg(REG_A))
	},
	0xeb: func() {},
	0xec: func() {},
	0xed: func() {},
	0xee: func() {
		// XOR d8
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.xor8RegD8(REG_A, val)
	},
	0xef: func() {},

	0xf0: func() {
		// LDH A, (a8)
		lsb := memory.read(cpu.readPc())
		addr := joinUint8(lsb, 0xff)
		cpu.set8Reg(REG_A, memory.read(addr))
	},
	0xf1: func() {},
	0xf2: func() {},
	0xf3: func() {},
	0xf4: func() {},
	0xf5: func() {},
	0xf6: func() {
		// OR d8
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.or8RegD8(REG_A, val)
	},
	0xf7: func() {},
	0xf8: func() {},
	0xf9: func() {},
	0xfa: func() {
		// LD A, (a16)
		lsb := lo(cpu.readPc())
		msb := hi(cpu.readPc())
		addr := joinUint8(msb, lsb)
		cpu.set8Reg(REG_A, memory.read(addr))
	},
	0xfb: func() {},
	0xfc: func() {},
	0xfd: func() {},
	0xfe: func() {
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.cp8RegD8(REG_A, val)
	},
	0xff: func() {},
}
