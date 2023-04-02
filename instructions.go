package main

var instructions = [0x100]func(){
	0x00: func() {}, // NOP
	0x01: func() {
		// LD BC, d16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		cpu.set16Reg(REG_BC, joinu8(msb, lsb))
	},
	0x02: func() {
		// LD (BC), A
		memory.write(cpu.read16Reg(REG_BC), cpu.read8Reg(REG_A))
	},
	0x03: func() {
		// INC BC
		cpu.inc16reg(REG_BC)
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
	0x07: func() {
		// RLCA
		cpu.rlca8Reg(REG_A)
	},
	0x08: func() {
		// LD (a16), SP
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		addr := joinu8(msb, lsb)
		memory.write(addr, lo(cpu.sp))
		memory.write(addr+1, hi(cpu.sp))
	},
	0x09: func() {
		// ADD HL, BC
		cpu.add16Reg(REG_HL, cpu.read16Reg(REG_BC))
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
		cpu.dec8Reg(REG_C)
	},
	0x0e: func() {
		// LD C, d8
		cpu.set8Reg(REG_C, memory.read(cpu.readPc()))
	},
	0x0f: func() {
		// RRCA
		cpu.rrca8Reg(REG_A)
	},

	0x10: func() {
		// STOP d8
	},
	0x11: func() {
		// LD DE, d16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		cpu.set16Reg(REG_DE, joinu8(msb, lsb))
	},
	0x12: func() {
		// LD (DE), A
		memory.write(cpu.read16Reg(REG_DE), cpu.read8Reg(REG_A))
	},
	0x13: func() {
		// INC DE
		cpu.inc16reg(REG_DE)
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
	0x17: func() {
		// RLA
		cpu.rla8Reg(REG_A)
	},
	0x18: func() {
		// JR r8
		val := int8(memory.read(cpu.readPc()))
		pc := int32(cpu.pc)
		addr := uint16(pc + int32(val))
		cpu.jump(addr)
	},
	0x19: func() {
		// ADD HL, DE
		cpu.add16Reg(REG_HL, cpu.read16Reg(REG_DE))
	},
	0x1a: func() {
		// LD A, (DE)
		addr := cpu.read16Reg(REG_DE)
		cpu.set8Reg(REG_A, memory.read(addr))
	},
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
	0x1e: func() {
		// LD E, d8
		val := memory.read(cpu.readPc())
		cpu.set8Reg(REG_E, val)
	},
	0x1f: func() {
		// RRA
		reg := cpu.read8Reg(REG_A)
		rotation := rotateRight(reg, 1)
		rotation = toggleBit(rotation, 7, cpu.C())
		cpu.set8Reg(REG_A, rotation)

		cpu.setC(isBitSet(reg, 0))
		cpu.setH(false)
		cpu.setN(false)
		cpu.setZ(false)
	},

	0x20: func() {
		// JR NZ, r8
		addr := cpu.readPc()
		val := int8(memory.read(addr))

		if !cpu.Z() {
			addr := int32(cpu.pc) + int32(val)
			cpu.jump(uint16(addr))
		}
	},
	0x21: func() {
		// LD HL, d16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		cpu.set16Reg(REG_HL, joinu8(msb, lsb))
	},
	0x22: func() {
		// LD (HL+), A
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_A))
		cpu.set16Reg(REG_HL, addr+1)
	},
	0x23: func() {
		// INC HL
		cpu.inc16reg(REG_HL)
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
	0x27: func() {
		// DAA
		a := cpu.read8Reg(REG_A)

		if !cpu.N() {
			if cpu.C() || a > 0x99 {
				a += 0x60
				cpu.setC(true)
			}

			if cpu.H() || (a&0x0f) > 0x09 {
				a += 0x6
			}
		} else {
			if cpu.C() {
				a -= 0x60
			}

			if cpu.H() {
				a -= 0x6
			}
		}

		cpu.setZ(a == 0)
		cpu.setH(false)

		cpu.set8Reg(REG_A, a)
	},
	0x28: func() {
		// JR Z, r8
		addr := cpu.readPc()
		val := int8(memory.read(addr))

		if cpu.Z() {
			addr := int32(cpu.pc) + int32(val)
			cpu.jump(uint16(addr))
		}
	},
	0x29: func() {
		// ADD HL, HL
		cpu.add16Reg(REG_HL, cpu.read16Reg(REG_HL))
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
	0x2e: func() {
		// LD L, d8
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.set8Reg(REG_L, val)
	},
	0x2f: func() {
		// CPL
		val := cpu.read8Reg(REG_A)
		cpu.set8Reg(REG_A, 0xff^val)
		cpu.setN(true)
		cpu.setH(true)
	},

	0x30: func() {
		// JR NC, r8
		pc := cpu.readPc()

		if !cpu.C() {
			val := int8(memory.read(pc))
			addr := int32(cpu.pc) + int32(val)
			cpu.jump(uint16(addr))
		}
	},
	0x31: func() {
		// LD SP, d16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		cpu.set16Reg(REG_SP, joinu8(msb, lsb))
	},
	0x32: func() {
		// LD (HL-), A
		addr := cpu.read16Reg(REG_HL)
		memory.write(addr, cpu.read8Reg(REG_A))
		cpu.set16Reg(REG_HL, addr-1)
	},
	0x33: func() {
		// INC SP
		cpu.sp = cpu.sp + 1
	},
	0x34: func() {
		// INC (HL)
		cpu.inc16reg(REG_HL)
	},
	0x35: func() {
		// DEC (HL)
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		dec := val - 1
		memory.write(addr, dec)

		cpu.setN(true)
		cpu.setH((val & 0xf) == 0)
		cpu.setZ(dec == 0)
	},
	0x36: func() {
		// LD (HL), d8
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(cpu.readPc())
		memory.write(addr, val)
	},
	0x37: func() {
		// SCF
		cpu.setN(false)
		cpu.setH(false)
		cpu.setC(true)
	},
	0x38: func() {
		// JR C, r8
		pc := cpu.readPc()

		if cpu.C() {
			val := int8(memory.read(pc))
			addr := int32(cpu.pc) + int32(val)
			cpu.jump(uint16(addr))
		}
	},
	0x39: func() {
		// ADD HL, SP
		cpu.add16Reg(REG_HL, cpu.sp)
	},
	0x3a: func() {
		// LD A, (HL-)
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.set8Reg(REG_A, val)
		cpu.set16Reg(REG_HL, addr-1)
	},
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
		cpu.set8Reg(REG_A, memory.read(cpu.readPc()))
	},
	0x3f: func() {
		// CCF
		cpu.setN(false)
		cpu.setH(false)
		cpu.setC(!cpu.C())
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
		// LD B, (HL)
		cpu.set8Reg(REG_B, memory.read(cpu.read16Reg(REG_HL)))
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
		cpu.set8Reg(REG_C, memory.read(cpu.read16Reg(REG_HL)))
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
		// LD D, (HL)
		cpu.set8Reg(REG_D, memory.read(cpu.read16Reg(REG_HL)))
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
		// LD E, (HL)
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
		cpu.halted = true
	},
	0x77: func() {
		// LD (HL), A
		memory.write(cpu.read16Reg(REG_HL), cpu.read8Reg(REG_A))
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
		b := cpu.read8Reg(REG_B)
		cpu.add8Reg(REG_A, b)
	},
	0x81: func() {
		// ADD A, C
		c := cpu.read8Reg(REG_C)
		cpu.add8Reg(REG_A, c)
	},
	0x82: func() {
		// ADD A, D
		d := cpu.read8Reg(REG_D)
		cpu.add8Reg(REG_A, d)
	},
	0x83: func() {
		// ADD A, E
		cpu.add8Reg(REG_A, cpu.read8Reg(REG_E))
	},
	0x84: func() {
		// ADD A, H
		h := cpu.read8Reg(REG_H)
		cpu.add8Reg(REG_A, h)
	},
	0x85: func() {
		// ADD A, L
		l := cpu.read8Reg(REG_L)
		cpu.add8Reg(REG_A, l)
	},
	0x86: func() {
		// ADD A, HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.add8Reg(REG_A, val)
	},
	0x87: func() {
		// ADD A, A
		a := cpu.read8Reg(REG_A)
		cpu.add8Reg(REG_A, a)
	},
	0x88: func() {
		// ADC A, B
		cpu.adc8Reg(REG_A, cpu.read8Reg(REG_B))
	},
	0x89: func() {
		// ADC A, C
		cpu.adc8Reg(REG_A, cpu.read8Reg(REG_C))
	},
	0x8a: func() {
		// ADC A, D
		cpu.adc8Reg(REG_A, cpu.read8Reg(REG_D))
	},
	0x8b: func() {
		// ADC A, E
		cpu.adc8Reg(REG_A, cpu.read8Reg(REG_E))
	},
	0x8c: func() {
		// ADC A, H
		cpu.adc8Reg(REG_A, cpu.read8Reg(REG_H))
	},
	0x8d: func() {
		// ADC A, L
		cpu.adc8Reg(REG_A, cpu.read8Reg(REG_L))
	},
	0x8e: func() {
		// ADC A, HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.adc8Reg(REG_A, val)
	},
	0x8f: func() {
		// ADC A, A
		cpu.adc8Reg(REG_A, cpu.read8Reg(REG_A))
	},

	0x90: func() {
		// SUB B
		cpu.sub8Reg(REG_A, cpu.read8Reg(REG_B))
	},
	0x91: func() {
		// SUB C
		cpu.sub8Reg(REG_A, cpu.read8Reg(REG_C))
	},
	0x92: func() {
		// SUB D
		cpu.sub8Reg(REG_A, cpu.read8Reg(REG_D))
	},
	0x93: func() {
		// SUB E
		cpu.sub8Reg(REG_A, cpu.read8Reg(REG_E))
	},
	0x94: func() {
		// SUB H
		cpu.sub8Reg(REG_A, cpu.read8Reg(REG_H))
	},
	0x95: func() {
		// SUB L
		cpu.sub8Reg(REG_A, cpu.read8Reg(REG_L))
	},
	0x96: func() {
		// SUB (HL)
		val := memory.read(cpu.read16Reg(REG_HL))
		cpu.sub8Reg(REG_A, val)
	},
	0x97: func() {
		// SUB A
		cpu.sub8Reg(REG_A, cpu.read8Reg(REG_A))
	},
	0x98: func() {
		// SBC A, B
		cpu.sbc8Reg(REG_A, cpu.read8Reg(REG_B))
	},
	0x99: func() {
		// SBC A, C
		cpu.sbc8Reg(REG_A, cpu.read8Reg(REG_C))
	},
	0x9a: func() {
		// SBC A, D
		cpu.sbc8Reg(REG_A, cpu.read8Reg(REG_D))
	},
	0x9b: func() {
		// SBC A, E
		cpu.sbc8Reg(REG_A, cpu.read8Reg(REG_E))
	},
	0x9c: func() {
		// SBC A, H
		cpu.sbc8Reg(REG_A, cpu.read8Reg(REG_H))
	},
	0x9d: func() {
		// SBC A, L
		cpu.sbc8Reg(REG_A, cpu.read8Reg(REG_L))
	},
	0x9e: func() {
		// SBC A, (HL)
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.sbc8Reg(REG_A, val)
	},
	0x9f: func() {
		// SBC A, A
		cpu.sbc8Reg(REG_A, cpu.read8Reg(REG_A))
	},

	0xa0: func() {
		// AND B
		cpu.and8Reg(REG_A, cpu.read8Reg(REG_B))
	},
	0xa1: func() {
		// AND C
		cpu.and8Reg(REG_A, cpu.read8Reg(REG_C))
	},
	0xa2: func() {
		// AND D
		cpu.and8Reg(REG_A, cpu.read8Reg(REG_D))
	},
	0xa3: func() {
		// AND E
		cpu.and8Reg(REG_A, cpu.read8Reg(REG_E))
	},
	0xa4: func() {
		// AND H
		cpu.and8Reg(REG_A, cpu.read8Reg(REG_H))
	},
	0xa5: func() {
		// AND L
		cpu.and8Reg(REG_A, cpu.read8Reg(REG_L))
	},
	0xa6: func() {
		// AND HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.and8Reg(REG_A, val)
	},
	0xa7: func() {
		// AND A
		cpu.and8Reg(REG_A, cpu.read8Reg(REG_A))
	},
	0xa8: func() {
		// XOR B
		cpu.xor8Reg(REG_A, cpu.read8Reg(REG_B))
	},
	0xa9: func() {
		// XOR C
		cpu.xor8Reg(REG_A, cpu.read8Reg(REG_C))
	},
	0xaa: func() {
		// XOR D
		cpu.xor8Reg(REG_A, cpu.read8Reg(REG_D))
	},
	0xab: func() {
		// XOR E
		cpu.xor8Reg(REG_A, cpu.read8Reg(REG_E))
	},
	0xac: func() {
		// XOR H
		cpu.xor8Reg(REG_A, cpu.read8Reg(REG_H))
	},
	0xad: func() {
		// XOR L
		cpu.xor8Reg(REG_A, cpu.read8Reg(REG_L))
	},
	0xae: func() {
		// XOR HL
		cpu.xor8Reg(REG_A, memory.read(cpu.read16Reg(REG_HL)))
	},
	0xaf: func() {
		// XOR A
		cpu.xor8Reg(REG_A, cpu.read8Reg(REG_A))
	},
	0xb0: func() {
		// OR B
		cpu.or8Reg(REG_A, cpu.read8Reg(REG_B))
	},
	0xb1: func() {
		// OR C
		cpu.or8Reg(REG_A, cpu.read8Reg(REG_C))
	},
	0xb2: func() {
		// OR D
		cpu.or8Reg(REG_A, cpu.read8Reg(REG_D))
	},
	0xb3: func() {
		// OR E
		cpu.or8Reg(REG_A, cpu.read8Reg(REG_E))
	},
	0xb4: func() {
		// OR H
		cpu.or8Reg(REG_A, cpu.read8Reg(REG_H))
	},
	0xb5: func() {
		// OR L
		cpu.or8Reg(REG_A, cpu.read8Reg(REG_L))
	},
	0xb6: func() {
		// OR HL
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.or8Reg(REG_A, val)
	},
	0xb7: func() {
		// OR A
		cpu.or8Reg(REG_A, cpu.read8Reg(REG_A))
	},
	0xb8: func() {
		// CP B
		cpu.cp8Reg(REG_A, cpu.read8Reg(REG_B))
	},
	0xb9: func() {
		// CP C
		cpu.cp8Reg(REG_A, cpu.read8Reg(REG_C))
	},
	0xba: func() {
		// CP D
		cpu.cp8Reg(REG_A, cpu.read8Reg(REG_D))
	},
	0xbb: func() {
		// CP E
		cpu.cp8Reg(REG_A, cpu.read8Reg(REG_E))
	},
	0xbc: func() {
		// CP H
		cpu.cp8Reg(REG_A, cpu.read8Reg(REG_H))
	},
	0xbd: func() {
		// CP L
		cpu.cp8Reg(REG_A, cpu.read8Reg(REG_L))
	},
	0xbe: func() {
		// CP (HL)
		addr := cpu.read16Reg(REG_HL)
		val := memory.read(addr)
		cpu.cp8Reg(REG_A, val)
	},
	0xbf: func() {
		// CP A
		cpu.cp8Reg(REG_A, cpu.read8Reg(REG_A))
	},

	0xc0: func() {
		// RET NZ
		if !cpu.Z() {
			cpu.ret()
		}
	},
	0xc1: func() {
		// POP BC
		cpu.popSp(REG_BC)
	},
	0xc2: func() {
		// JP NZ, a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		addr := joinu8(msb, lsb)

		if !cpu.Z() {
			cpu.jump(addr)
		}
	},
	0xc3: func() {
		// JP a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		addr := joinu8(msb, lsb)
		cpu.jump(addr)
	},
	0xc4: func() {
		// CALL NZ, a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())

		if !cpu.Z() {
			addr := joinu8(msb, lsb)
			cpu.call(addr)
		}
	},
	0xc5: func() {
		// PUSH BC
		cpu.pushSp(REG_BC)
	},
	0xc6: func() {
		// ADD A, d8
		cpu.add8Reg(REG_A, memory.read(cpu.readPc()))
	},
	0xc7: func() {
		// RST 00H
		cpu.call(0x0000)
	},
	0xc8: func() {
		// RET Z
		if cpu.Z() {
			cpu.ret()
		}
	},
	0xc9: func() {
		//  RET
		cpu.ret()
	},
	0xca: func() {
		// JP Z, a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())

		if cpu.Z() {
			cpu.jump(joinu8(msb, lsb))
		}
	},
	0xcc: func() {
		//  CALL Z, a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())

		if cpu.Z() {
			addr := joinu8(msb, lsb)
			cpu.call(addr)
		}
	},
	0xcd: func() {
		// CALL a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		addr := joinu8(msb, lsb)
		cpu.call(addr)
	},
	0xce: func() {
		// ADC A, d8
		cpu.adc8Reg(REG_A, memory.read(cpu.readPc()))
	},
	0xcf: func() {
		// RST 08H
		cpu.call(0x0008)
	},

	0xd0: func() {
		// RET NC
		if !cpu.C() {
			cpu.ret()
		}
	},
	0xd1: func() {
		// POP DE
		cpu.popSp(REG_DE)
	},
	0xd2: func() {
		// JP NC, a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())

		if !cpu.C() {
			cpu.jump(joinu8(msb, lsb))
		}
	},
	0xd4: func() {
		// CALL NC, a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())

		if !cpu.C() {
			cpu.call(joinu8(msb, lsb))
		}
	},
	0xd5: func() {
		// PUSH DE
		cpu.pushSp(REG_DE)
	},
	0xd6: func() {
		// SUB d8
		cpu.sub8Reg(REG_A, memory.read(cpu.readPc()))
	},
	0xd7: func() {
		// RST 10H
		cpu.call(0x0010)
	},
	0xd8: func() {
		// RET C
		if cpu.C() {
			cpu.ret()
		}
	},
	0xd9: func() {
		// RETI
		cpu.ret()
		cpu.setIME(true)
	},
	0xda: func() {
		// JP C, a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())

		if cpu.C() {
			cpu.jump(joinu8(msb, lsb))
		}
	},
	0xdc: func() {
		// CALL C, a16
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())

		if cpu.C() {
			addr := joinu8(msb, lsb)
			cpu.call(addr)
		}
	},
	0xde: func() {
		// SBC A, d8
		cpu.sbc8Reg(REG_A, memory.read(cpu.readPc()))
	},
	0xdf: func() {
		// RST 18H
		cpu.call(0x0018)
	},

	0xe0: func() {
		// LDH (a8), A
		lsb := memory.read(cpu.readPc())
		addr := joinu8(0xff, lsb)
		memory.write(addr, cpu.read8Reg(REG_A))
	},
	0xe1: func() {
		// POP HL
		cpu.popSp(REG_HL)
	},
	0xe2: func() {
		// LD (C), A
		addr := joinu8(cpu.read8Reg(REG_C), 0xff)
		val := cpu.read8Reg(REG_A)
		memory.write(addr, val)
	},
	0xe5: func() {
		// PUSH HL
		cpu.pushSp(REG_HL)
	},
	0xe6: func() {
		// AND d8
		addr := cpu.readPc()
		val := memory.read(addr)
		cpu.and8Reg(REG_A, val)
	},
	0xe7: func() {
		// RST 20H
		cpu.call(0x0020)
	},
	0xe8: func() {
		// ADD SP, r8
		sp := cpu.read16Reg(REG_SP)
		r8 := int8(memory.read(cpu.readPc()))
		add := sp + uint16(r8)
		cpu.set16Reg(REG_SP, add)

		carry := sp ^ uint16(r8) ^ add
		cpu.setZ(false)
		cpu.setN(false)
		cpu.setH(carry&0x10 == 0x10)
		cpu.setC(carry&0x100 == 0x100)
	},
	0xe9: func() {
		// JP HL
		addr := cpu.read16Reg(REG_HL)
		cpu.jump(addr)
	},
	0xea: func() {
		// LD (a16), A
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		addr := joinu8(msb, lsb)
		memory.write(addr, cpu.read8Reg(REG_A))
	},
	0xee: func() {
		// XOR d8
		cpu.xor8Reg(REG_A, memory.read(cpu.readPc()))
	},
	0xef: func() {
		// RST 28H
		cpu.call(0x0028)
	},

	0xf0: func() {
		// LDH A, (a8)
		lsb := memory.read(cpu.readPc())
		addr := joinu8(0xff, lsb)
		cpu.set8Reg(REG_A, memory.read(addr))
	},
	0xf1: func() {
		// POP AF
		cpu.popSp(REG_AF)
		cpu.set16Reg(REG_AF, cpu.read16Reg(REG_AF)&0xfff0)
	},
	0xf2: func() {
		// LD A, (C)
		lsb := cpu.read8Reg(REG_C)
		addr := joinu8(0xff, lsb)
		val := memory.read(addr)
		cpu.set8Reg(REG_A, val)
	},
	0xf3: func() {
		// DI
		cpu.setIME(false)
	},
	0xf5: func() {
		// PUSH AF
		cpu.pushSp(REG_AF)
	},
	0xf6: func() {
		// OR d8
		cpu.or8Reg(REG_A, memory.read(cpu.readPc()))
	},
	0xf7: func() {
		// RST 30H
		cpu.call(0x0030)
	},
	0xf8: func() {
		//  LD HL, SP + r8
		r8 := int8(memory.read(cpu.readPc()))
		add := int32(cpu.sp) + int32(r8)
		cpu.set16Reg(REG_HL, uint16(add))

		carry := cpu.sp ^ uint16(r8) ^ uint16(add)
		cpu.setZ(false)
		cpu.setN(false)
		cpu.setH(carry&0x10 == 0x10)
		cpu.setC(carry&0x100 == 0x100)
	},
	0xf9: func() {
		// LD SP, HL
		cpu.set16Reg(REG_SP, cpu.read16Reg(REG_HL))
	},
	0xfa: func() {
		// LD A, (a16)
		lsb := memory.read(cpu.readPc())
		msb := memory.read(cpu.readPc())
		addr := joinu8(msb, lsb)
		cpu.set8Reg(REG_A, memory.read(addr))
	},
	0xfb: func() {
		// IE
		cpu.setIme = true
	},
	0xfe: func() {
		cpu.cp8Reg(REG_A, memory.read(cpu.readPc()))
	},
	0xff: func() {
		// RST 38H
		cpu.call(0x0038)
	},
}
