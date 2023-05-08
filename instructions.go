package main

var instructionCycles = []int{
	1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1,
	0, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1,
	2, 3, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1,
	2, 3, 2, 2, 3, 3, 3, 1, 2, 2, 2, 2, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1,
	2, 3, 3, 4, 3, 4, 2, 4, 2, 4, 3, 0, 3, 6, 2, 4,
	2, 3, 3, 0, 3, 4, 2, 4, 2, 4, 3, 0, 3, 0, 2, 4,
	3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4,
	3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4,
}

var instructions = [0x100]func(){
	0x00: func() {}, // NOP
	0x01: func() {
		// LD BC, d16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		gb.cpu.set16Reg(REG_BC, joinu8(msb, lsb))
	},
	0x02: func() {
		// LD (BC), A
		gb.bus.write(gb.cpu.read16Reg(REG_BC), gb.cpu.read8Reg(REG_A))
	},
	0x03: func() {
		// INC BC
		gb.cpu.inc16reg(REG_BC)
	},
	0x04: func() {
		// INC B
		gb.cpu.inc8Reg(REG_B)
	},
	0x05: func() {
		// DEC B
		gb.cpu.dec8Reg(REG_B)
	},
	0x06: func() {
		// LD B, d8
		val := gb.bus.read(gb.cpu.readPc())
		gb.cpu.set8Reg(REG_B, val)
	},
	0x07: func() {
		// RLCA
		gb.cpu.rlca8Reg(REG_A)
	},
	0x08: func() {
		// LD (a16), SP
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		gb.bus.write(addr, lo(gb.cpu.sp))
		gb.bus.write(addr+1, hi(gb.cpu.sp))
	},
	0x09: func() {
		// ADD HL, BC
		gb.cpu.add16Reg(REG_HL, gb.cpu.read16Reg(REG_BC))
	},
	0x0a: func() {
		// LD A, (BC)
		val := gb.bus.read(gb.cpu.read16Reg(REG_BC))
		gb.cpu.set8Reg(REG_A, val)
	},
	0x0b: func() {
		// DEC BC
		gb.cpu.dec16Reg(REG_BC)
	},
	0x0c: func() {
		// INC C
		gb.cpu.inc8Reg(REG_C)
	},

	0x0d: func() {
		// DEC C
		gb.cpu.dec8Reg(REG_C)
	},
	0x0e: func() {
		// LD C, d8
		gb.cpu.set8Reg(REG_C, gb.bus.read(gb.cpu.readPc()))
	},
	0x0f: func() {
		// RRCA
		gb.cpu.rrca8Reg(REG_A)
	},

	0x10: func() {
		// STOP d8
		gb.timer.resetDIV()
	},
	0x11: func() {
		// LD DE, d16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		gb.cpu.set16Reg(REG_DE, joinu8(msb, lsb))
	},
	0x12: func() {
		// LD (DE), A
		gb.bus.write(gb.cpu.read16Reg(REG_DE), gb.cpu.read8Reg(REG_A))
	},
	0x13: func() {
		// INC DE
		gb.cpu.inc16reg(REG_DE)
	},
	0x14: func() {
		// INC D
		gb.cpu.inc8Reg(REG_D)
	},
	0x15: func() {
		// DEC D
		gb.cpu.dec8Reg(REG_D)
	},
	0x16: func() {
		// LD D, d8
		gb.cpu.set8Reg(REG_D, gb.bus.read(gb.cpu.readPc()))
	},
	0x17: func() {
		// RLA
		gb.cpu.rla8Reg(REG_A)
	},
	0x18: func() {
		// JR r8
		val := int8(gb.bus.read(gb.cpu.readPc()))
		pc := int32(gb.cpu.pc)
		addr := uint16(pc + int32(val))
		gb.cpu.jump(addr)
	},
	0x19: func() {
		// ADD HL, DE
		gb.cpu.add16Reg(REG_HL, gb.cpu.read16Reg(REG_DE))
	},
	0x1a: func() {
		// LD A, (DE)
		addr := gb.cpu.read16Reg(REG_DE)
		gb.cpu.set8Reg(REG_A, gb.bus.read(addr))
	},
	0x1b: func() {
		// DEC DE
		gb.cpu.dec16Reg(REG_DE)
	},
	0x1c: func() {
		// INC E
		gb.cpu.inc8Reg(REG_E)
	},
	0x1d: func() {
		// DEC E
		gb.cpu.dec8Reg(REG_E)
	},
	0x1e: func() {
		// LD E, d8
		val := gb.bus.read(gb.cpu.readPc())
		gb.cpu.set8Reg(REG_E, val)
	},
	0x1f: func() {
		// RRA
		reg := gb.cpu.read8Reg(REG_A)
		rotation := rotateRight(reg, 1)
		rotation = toggleBit(rotation, 7, gb.cpu.C())
		gb.cpu.set8Reg(REG_A, rotation)

		gb.cpu.setC(isBitSet(reg, 0))
		gb.cpu.setH(false)
		gb.cpu.setN(false)
		gb.cpu.setZ(false)
	},

	0x20: func() {
		// JR NZ, r8
		addr := gb.cpu.readPc()
		val := int8(gb.bus.read(addr))

		if !gb.cpu.Z() {
			addr := int32(gb.cpu.pc) + int32(val)
			gb.cpu.jump(uint16(addr))
			gb.cpu.clockCycles += 4
		}
	},
	0x21: func() {
		// LD HL, d16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		gb.cpu.set16Reg(REG_HL, joinu8(msb, lsb))
	},
	0x22: func() {
		// LD (HL+), A
		addr := gb.cpu.read16Reg(REG_HL)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_A))
		gb.cpu.set16Reg(REG_HL, addr+1)
	},
	0x23: func() {
		// INC HL
		gb.cpu.inc16reg(REG_HL)
	},
	0x24: func() {
		// INC H
		gb.cpu.inc8Reg(REG_H)
	},
	0x25: func() {
		// DEC H
		gb.cpu.dec8Reg(REG_H)
	},
	0x26: func() {
		// LD H, d8
		gb.cpu.set8Reg(REG_H, gb.bus.read(gb.cpu.readPc()))
	},
	0x27: func() {
		// DAA
		a := gb.cpu.read8Reg(REG_A)

		if !gb.cpu.N() {
			if gb.cpu.C() || a > 0x99 {
				a += 0x60
				gb.cpu.setC(true)
			}

			if gb.cpu.H() || (a&0x0f) > 0x09 {
				a += 0x6
			}
		} else {
			if gb.cpu.C() {
				a -= 0x60
			}

			if gb.cpu.H() {
				a -= 0x6
			}
		}

		gb.cpu.setZ(a == 0)
		gb.cpu.setH(false)

		gb.cpu.set8Reg(REG_A, a)
	},
	0x28: func() {
		// JR Z, r8
		addr := gb.cpu.readPc()
		val := int8(gb.bus.read(addr))

		if gb.cpu.Z() {
			addr := int32(gb.cpu.pc) + int32(val)
			gb.cpu.jump(uint16(addr))
			gb.cpu.clockCycles += 4
		}
	},
	0x29: func() {
		// ADD HL, HL
		gb.cpu.add16Reg(REG_HL, gb.cpu.read16Reg(REG_HL))
	},
	0x2a: func() {
		// LD A, (HL+)
		val := gb.bus.read(gb.cpu.read16Reg(REG_HL))
		gb.cpu.set8Reg(REG_A, val)
		gb.cpu.set16Reg(REG_HL, gb.cpu.read16Reg(REG_HL)+1)
	},
	0x2b: func() {
		// DEC HL
		gb.cpu.dec16Reg(REG_HL)
	},
	0x2c: func() {
		// INC L
		gb.cpu.inc8Reg(REG_L)
	},
	0x2d: func() {
		// DEC L
		gb.cpu.dec8Reg(REG_L)
	},
	0x2e: func() {
		// LD L, d8
		addr := gb.cpu.readPc()
		val := gb.bus.read(addr)
		gb.cpu.set8Reg(REG_L, val)
	},
	0x2f: func() {
		// CPL
		val := gb.cpu.read8Reg(REG_A)
		gb.cpu.set8Reg(REG_A, 0xff^val)
		gb.cpu.setN(true)
		gb.cpu.setH(true)
	},

	0x30: func() {
		// JR NC, r8
		pc := gb.cpu.readPc()

		if !gb.cpu.C() {
			val := int8(gb.bus.read(pc))
			addr := int32(gb.cpu.pc) + int32(val)
			gb.cpu.jump(uint16(addr))
			gb.cpu.clockCycles += 4
		}
	},
	0x31: func() {
		// LD SP, d16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		gb.cpu.set16Reg(REG_SP, joinu8(msb, lsb))
	},
	0x32: func() {
		// LD (HL-), A
		addr := gb.cpu.read16Reg(REG_HL)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_A))
		gb.cpu.set16Reg(REG_HL, addr-1)
	},
	0x33: func() {
		// INC SP
		gb.cpu.sp = gb.cpu.sp + 1
	},
	0x34: func() {
		// INC (HL)
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		inc := val + 1
		gb.bus.write(addr, inc)

		gb.cpu.setN(false)
		gb.cpu.setH((val&0xf)+1 > 0xf)
		gb.cpu.setZ(inc == 0)
	},
	0x35: func() {
		// DEC (HL)
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		dec := val - 1
		gb.bus.write(addr, dec)

		gb.cpu.setN(true)
		gb.cpu.setH((val & 0xf) == 0)
		gb.cpu.setZ(dec == 0)
	},
	0x36: func() {
		// LD (HL), d8
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(gb.cpu.readPc())
		gb.bus.write(addr, val)
	},
	0x37: func() {
		// SCF
		gb.cpu.setN(false)
		gb.cpu.setH(false)
		gb.cpu.setC(true)
	},
	0x38: func() {
		// JR C, r8
		pc := gb.cpu.readPc()

		if gb.cpu.C() {
			val := int8(gb.bus.read(pc))
			addr := int32(gb.cpu.pc) + int32(val)
			gb.cpu.jump(uint16(addr))
			gb.cpu.clockCycles += 4
		}
	},
	0x39: func() {
		// ADD HL, SP
		gb.cpu.add16Reg(REG_HL, gb.cpu.sp)
	},
	0x3a: func() {
		// LD A, (HL-)
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		gb.cpu.set8Reg(REG_A, val)
		gb.cpu.set16Reg(REG_HL, addr-1)
	},
	0x3b: func() {
		// DEC SP
		gb.cpu.sp = gb.cpu.sp - 1
	},
	0x3c: func() {
		// INC A
		gb.cpu.inc8Reg(REG_A)
	},
	0x3d: func() {
		// DEC A
		gb.cpu.dec8Reg(REG_A)
	},
	0x3e: func() {
		// LD A, d8
		gb.cpu.set8Reg(REG_A, gb.bus.read(gb.cpu.readPc()))
	},
	0x3f: func() {
		// CCF
		gb.cpu.setN(false)
		gb.cpu.setH(false)
		gb.cpu.setC(!gb.cpu.C())
	},

	0x40: func() {
		// LD B, B
		gb.cpu.load8Reg(REG_B, REG_B)
	},
	0x41: func() {
		// LD B, C
		gb.cpu.load8Reg(REG_B, REG_C)
	},
	0x42: func() {
		// LD B, D
		gb.cpu.load8Reg(REG_B, REG_D)
	},
	0x43: func() {
		// LD B, E
		gb.cpu.load8Reg(REG_B, REG_E)
	},
	0x44: func() {
		// LD B, H
		gb.cpu.load8Reg(REG_B, REG_H)
	},
	0x45: func() {
		// LD B, L
		gb.cpu.load8Reg(REG_B, REG_L)
	},
	0x46: func() {
		// LD B, (HL)
		gb.cpu.set8Reg(REG_B, gb.bus.read(gb.cpu.read16Reg(REG_HL)))
	},
	0x47: func() {
		// LD B, A
		gb.cpu.load8Reg(REG_B, REG_A)
	},
	0x48: func() {
		// LD C, B
		gb.cpu.load8Reg(REG_C, REG_B)
	},
	0x49: func() {
		// LD C, C
		gb.cpu.load8Reg(REG_C, REG_C)
	},
	0x4a: func() {
		// LD C, D
		gb.cpu.load8Reg(REG_C, REG_D)
	},
	0x4b: func() {
		// LD C, E
		gb.cpu.load8Reg(REG_C, REG_E)
	},
	0x4c: func() {
		// LD C, H
		gb.cpu.load8Reg(REG_C, REG_H)
	},
	0x4d: func() {
		// LD C, L
		gb.cpu.load8Reg(REG_C, REG_L)
	},
	0x4e: func() {
		// LD C, (HL)
		gb.cpu.set8Reg(REG_C, gb.bus.read(gb.cpu.read16Reg(REG_HL)))
	},
	0x4f: func() {
		// LD C, A
		gb.cpu.load8Reg(REG_C, REG_A)
	},
	0x50: func() {
		// LD D, B
		gb.cpu.load8Reg(REG_D, REG_B)
	},
	0x51: func() {
		// LD D, C
		gb.cpu.load8Reg(REG_D, REG_C)
	},
	0x52: func() {
		// LD D, D
		gb.cpu.load8Reg(REG_D, REG_D)
	},
	0x53: func() {
		// LD D, E
		gb.cpu.load8Reg(REG_D, REG_E)
	},
	0x54: func() {
		// LD D, H
		gb.cpu.load8Reg(REG_D, REG_H)
	},
	0x55: func() {
		// LD D, L
		gb.cpu.load8Reg(REG_D, REG_L)
	},
	0x56: func() {
		// LD D, (HL)
		gb.cpu.set8Reg(REG_D, gb.bus.read(gb.cpu.read16Reg(REG_HL)))
	},
	0x57: func() {
		// LD D, A
		gb.cpu.load8Reg(REG_D, REG_A)
	},
	0x58: func() {
		// LD E, B
		gb.cpu.load8Reg(REG_E, REG_B)
	},
	0x59: func() {
		// LD E, C
		gb.cpu.load8Reg(REG_E, REG_C)
	},
	0x5a: func() {
		// LD E, D
		gb.cpu.load8Reg(REG_E, REG_D)
	},
	0x5b: func() {
		// LD E, E
		gb.cpu.load8Reg(REG_E, REG_E)
	},
	0x5c: func() {
		// LD E, H
		gb.cpu.load8Reg(REG_E, REG_H)
	},
	0x5d: func() {
		// LD E, L
		gb.cpu.load8Reg(REG_E, REG_L)
	},
	0x5e: func() {
		// LD E, (HL)
		val := gb.bus.read(gb.cpu.read16Reg(REG_HL))
		gb.cpu.set8Reg(REG_E, val)
	},
	0x5f: func() {
		// LD E, A
		gb.cpu.load8Reg(REG_E, REG_A)
	},

	0x60: func() {
		// LD H, B
		gb.cpu.load8Reg(REG_H, REG_B)
	},
	0x61: func() {
		// LD H, C
		gb.cpu.load8Reg(REG_H, REG_C)
	},
	0x62: func() {
		// LD H, D
		gb.cpu.load8Reg(REG_H, REG_D)
	},
	0x63: func() {
		// LD H, E
		gb.cpu.load8Reg(REG_H, REG_E)
	},
	0x64: func() {
		// LD H, H
		gb.cpu.load8Reg(REG_H, REG_H)
	},
	0x65: func() {
		// LD H, L
		gb.cpu.load8Reg(REG_H, REG_L)
	},
	0x66: func() {
		// LD H, (HL)
		val := gb.bus.read(gb.cpu.read16Reg(REG_HL))
		gb.cpu.set8Reg(REG_H, val)
	},
	0x67: func() {
		// LD H, A
		gb.cpu.load8Reg(REG_H, REG_A)
	},
	0x68: func() {
		// LD L, B
		gb.cpu.load8Reg(REG_L, REG_B)
	},
	0x69: func() {
		// LD L, C
		gb.cpu.load8Reg(REG_L, REG_C)
	},
	0x6a: func() {
		// LD L, D
		gb.cpu.load8Reg(REG_L, REG_D)
	},
	0x6b: func() {
		// LD L, E
		gb.cpu.load8Reg(REG_L, REG_E)
	},
	0x6c: func() {
		// LD L, H
		gb.cpu.load8Reg(REG_L, REG_H)
	},
	0x6d: func() {
		// LD L, L
		gb.cpu.load8Reg(REG_L, REG_L)
	},
	0x6e: func() {
		// LD L, (HL)
		val := gb.bus.read(gb.cpu.read16Reg(REG_HL))
		gb.cpu.set8Reg(REG_L, val)
	},
	0x6f: func() {
		// LD L, A
		gb.cpu.load8Reg(REG_L, REG_A)
	},
	0x70: func() {
		// LD (HL), B
		addr := gb.cpu.read16Reg(REG_HL)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_B))
	},
	0x71: func() {
		// LD (HL), C
		addr := gb.cpu.read16Reg(REG_HL)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_C))
	},
	0x72: func() {
		// LD (HL), D
		addr := gb.cpu.read16Reg(REG_HL)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_D))
	},
	0x73: func() {
		// LD (HL), E
		addr := gb.cpu.read16Reg(REG_HL)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_E))
	},
	0x74: func() {
		// LD (HL), H
		addr := gb.cpu.read16Reg(REG_HL)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_H))
	},
	0x75: func() {
		// LD (HL), L
		addr := gb.cpu.read16Reg(REG_HL)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_L))
	},
	0x76: func() {
		// HALT
		gb.cpu.halted = true
	},
	0x77: func() {
		// LD (HL), A
		gb.bus.write(gb.cpu.read16Reg(REG_HL), gb.cpu.read8Reg(REG_A))
	},
	0x78: func() {
		// LD A, B
		gb.cpu.load8Reg(REG_A, REG_B)
	},
	0x79: func() {
		// LD A, C
		gb.cpu.load8Reg(REG_A, REG_C)
	},
	0x7a: func() {
		// LD A, D
		gb.cpu.load8Reg(REG_A, REG_D)
	},
	0x7b: func() {
		// LD A, E
		gb.cpu.load8Reg(REG_A, REG_E)
	},
	0x7c: func() {
		// LD A, H
		gb.cpu.load8Reg(REG_A, REG_H)
	},
	0x7d: func() {
		// LD A, L
		gb.cpu.load8Reg(REG_A, REG_L)
	},
	0x7e: func() {
		// LD A, (HL)
		val := gb.bus.read(gb.cpu.read16Reg(REG_HL))
		gb.cpu.set8Reg(REG_A, val)
	},
	0x7f: func() {
		// LD A, A
		gb.cpu.load8Reg(REG_A, REG_A)
	},

	0x80: func() {
		// ADD A, B
		b := gb.cpu.read8Reg(REG_B)
		gb.cpu.add8Reg(REG_A, b)
	},
	0x81: func() {
		// ADD A, C
		c := gb.cpu.read8Reg(REG_C)
		gb.cpu.add8Reg(REG_A, c)
	},
	0x82: func() {
		// ADD A, D
		d := gb.cpu.read8Reg(REG_D)
		gb.cpu.add8Reg(REG_A, d)
	},
	0x83: func() {
		// ADD A, E
		gb.cpu.add8Reg(REG_A, gb.cpu.read8Reg(REG_E))
	},
	0x84: func() {
		// ADD A, H
		h := gb.cpu.read8Reg(REG_H)
		gb.cpu.add8Reg(REG_A, h)
	},
	0x85: func() {
		// ADD A, L
		l := gb.cpu.read8Reg(REG_L)
		gb.cpu.add8Reg(REG_A, l)
	},
	0x86: func() {
		// ADD A, HL
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		gb.cpu.add8Reg(REG_A, val)
	},
	0x87: func() {
		// ADD A, A
		a := gb.cpu.read8Reg(REG_A)
		gb.cpu.add8Reg(REG_A, a)
	},
	0x88: func() {
		// ADC A, B
		gb.cpu.adc8Reg(REG_A, gb.cpu.read8Reg(REG_B))
	},
	0x89: func() {
		// ADC A, C
		gb.cpu.adc8Reg(REG_A, gb.cpu.read8Reg(REG_C))
	},
	0x8a: func() {
		// ADC A, D
		gb.cpu.adc8Reg(REG_A, gb.cpu.read8Reg(REG_D))
	},
	0x8b: func() {
		// ADC A, E
		gb.cpu.adc8Reg(REG_A, gb.cpu.read8Reg(REG_E))
	},
	0x8c: func() {
		// ADC A, H
		gb.cpu.adc8Reg(REG_A, gb.cpu.read8Reg(REG_H))
	},
	0x8d: func() {
		// ADC A, L
		gb.cpu.adc8Reg(REG_A, gb.cpu.read8Reg(REG_L))
	},
	0x8e: func() {
		// ADC A, HL
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		gb.cpu.adc8Reg(REG_A, val)
	},
	0x8f: func() {
		// ADC A, A
		gb.cpu.adc8Reg(REG_A, gb.cpu.read8Reg(REG_A))
	},

	0x90: func() {
		// SUB B
		gb.cpu.sub8Reg(REG_A, gb.cpu.read8Reg(REG_B))
	},
	0x91: func() {
		// SUB C
		gb.cpu.sub8Reg(REG_A, gb.cpu.read8Reg(REG_C))
	},
	0x92: func() {
		// SUB D
		gb.cpu.sub8Reg(REG_A, gb.cpu.read8Reg(REG_D))
	},
	0x93: func() {
		// SUB E
		gb.cpu.sub8Reg(REG_A, gb.cpu.read8Reg(REG_E))
	},
	0x94: func() {
		// SUB H
		gb.cpu.sub8Reg(REG_A, gb.cpu.read8Reg(REG_H))
	},
	0x95: func() {
		// SUB L
		gb.cpu.sub8Reg(REG_A, gb.cpu.read8Reg(REG_L))
	},
	0x96: func() {
		// SUB (HL)
		val := gb.bus.read(gb.cpu.read16Reg(REG_HL))
		gb.cpu.sub8Reg(REG_A, val)
	},
	0x97: func() {
		// SUB A
		gb.cpu.sub8Reg(REG_A, gb.cpu.read8Reg(REG_A))
	},
	0x98: func() {
		// SBC A, B
		gb.cpu.sbc8Reg(REG_A, gb.cpu.read8Reg(REG_B))
	},
	0x99: func() {
		// SBC A, C
		gb.cpu.sbc8Reg(REG_A, gb.cpu.read8Reg(REG_C))
	},
	0x9a: func() {
		// SBC A, D
		gb.cpu.sbc8Reg(REG_A, gb.cpu.read8Reg(REG_D))
	},
	0x9b: func() {
		// SBC A, E
		gb.cpu.sbc8Reg(REG_A, gb.cpu.read8Reg(REG_E))
	},
	0x9c: func() {
		// SBC A, H
		gb.cpu.sbc8Reg(REG_A, gb.cpu.read8Reg(REG_H))
	},
	0x9d: func() {
		// SBC A, L
		gb.cpu.sbc8Reg(REG_A, gb.cpu.read8Reg(REG_L))
	},
	0x9e: func() {
		// SBC A, (HL)
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		gb.cpu.sbc8Reg(REG_A, val)
	},
	0x9f: func() {
		// SBC A, A
		gb.cpu.sbc8Reg(REG_A, gb.cpu.read8Reg(REG_A))
	},

	0xa0: func() {
		// AND B
		gb.cpu.and8Reg(REG_A, gb.cpu.read8Reg(REG_B))
	},
	0xa1: func() {
		// AND C
		gb.cpu.and8Reg(REG_A, gb.cpu.read8Reg(REG_C))
	},
	0xa2: func() {
		// AND D
		gb.cpu.and8Reg(REG_A, gb.cpu.read8Reg(REG_D))
	},
	0xa3: func() {
		// AND E
		gb.cpu.and8Reg(REG_A, gb.cpu.read8Reg(REG_E))
	},
	0xa4: func() {
		// AND H
		gb.cpu.and8Reg(REG_A, gb.cpu.read8Reg(REG_H))
	},
	0xa5: func() {
		// AND L
		gb.cpu.and8Reg(REG_A, gb.cpu.read8Reg(REG_L))
	},
	0xa6: func() {
		// AND HL
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		gb.cpu.and8Reg(REG_A, val)
	},
	0xa7: func() {
		// AND A
		gb.cpu.and8Reg(REG_A, gb.cpu.read8Reg(REG_A))
	},
	0xa8: func() {
		// XOR B
		gb.cpu.xor8Reg(REG_A, gb.cpu.read8Reg(REG_B))
	},
	0xa9: func() {
		// XOR C
		gb.cpu.xor8Reg(REG_A, gb.cpu.read8Reg(REG_C))
	},
	0xaa: func() {
		// XOR D
		gb.cpu.xor8Reg(REG_A, gb.cpu.read8Reg(REG_D))
	},
	0xab: func() {
		// XOR E
		gb.cpu.xor8Reg(REG_A, gb.cpu.read8Reg(REG_E))
	},
	0xac: func() {
		// XOR H
		gb.cpu.xor8Reg(REG_A, gb.cpu.read8Reg(REG_H))
	},
	0xad: func() {
		// XOR L
		gb.cpu.xor8Reg(REG_A, gb.cpu.read8Reg(REG_L))
	},
	0xae: func() {
		// XOR HL
		gb.cpu.xor8Reg(REG_A, gb.bus.read(gb.cpu.read16Reg(REG_HL)))
	},
	0xaf: func() {
		// XOR A
		gb.cpu.xor8Reg(REG_A, gb.cpu.read8Reg(REG_A))
	},
	0xb0: func() {
		// OR B
		gb.cpu.or8Reg(REG_A, gb.cpu.read8Reg(REG_B))
	},
	0xb1: func() {
		// OR C
		gb.cpu.or8Reg(REG_A, gb.cpu.read8Reg(REG_C))
	},
	0xb2: func() {
		// OR D
		gb.cpu.or8Reg(REG_A, gb.cpu.read8Reg(REG_D))
	},
	0xb3: func() {
		// OR E
		gb.cpu.or8Reg(REG_A, gb.cpu.read8Reg(REG_E))
	},
	0xb4: func() {
		// OR H
		gb.cpu.or8Reg(REG_A, gb.cpu.read8Reg(REG_H))
	},
	0xb5: func() {
		// OR L
		gb.cpu.or8Reg(REG_A, gb.cpu.read8Reg(REG_L))
	},
	0xb6: func() {
		// OR HL
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		gb.cpu.or8Reg(REG_A, val)
	},
	0xb7: func() {
		// OR A
		gb.cpu.or8Reg(REG_A, gb.cpu.read8Reg(REG_A))
	},
	0xb8: func() {
		// CP B
		gb.cpu.cp8Reg(REG_A, gb.cpu.read8Reg(REG_B))
	},
	0xb9: func() {
		// CP C
		gb.cpu.cp8Reg(REG_A, gb.cpu.read8Reg(REG_C))
	},
	0xba: func() {
		// CP D
		gb.cpu.cp8Reg(REG_A, gb.cpu.read8Reg(REG_D))
	},
	0xbb: func() {
		// CP E
		gb.cpu.cp8Reg(REG_A, gb.cpu.read8Reg(REG_E))
	},
	0xbc: func() {
		// CP H
		gb.cpu.cp8Reg(REG_A, gb.cpu.read8Reg(REG_H))
	},
	0xbd: func() {
		// CP L
		gb.cpu.cp8Reg(REG_A, gb.cpu.read8Reg(REG_L))
	},
	0xbe: func() {
		// CP (HL)
		addr := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(addr)
		gb.cpu.cp8Reg(REG_A, val)
	},
	0xbf: func() {
		// CP A
		gb.cpu.cp8Reg(REG_A, gb.cpu.read8Reg(REG_A))
	},

	0xc0: func() {
		// RET NZ
		if !gb.cpu.Z() {
			gb.cpu.ret()
			gb.cpu.clockCycles += 12
		}
	},
	0xc1: func() {
		// POP BC
		gb.cpu.popSp(REG_BC)
	},
	0xc2: func() {
		// JP NZ, a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)

		if !gb.cpu.Z() {
			gb.cpu.jump(addr)
			gb.cpu.clockCycles += 4
		}
	},
	0xc3: func() {
		// JP a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		gb.cpu.jump(addr)
	},
	0xc4: func() {
		// CALL NZ, a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())

		if !gb.cpu.Z() {
			addr := joinu8(msb, lsb)
			gb.cpu.call(addr)
			gb.cpu.clockCycles += 12
		}
	},
	0xc5: func() {
		// PUSH BC
		gb.cpu.pushSp(REG_BC)
	},
	0xc6: func() {
		// ADD A, d8
		gb.cpu.add8Reg(REG_A, gb.bus.read(gb.cpu.readPc()))
	},
	0xc7: func() {
		// RST 00H
		gb.cpu.call(0x0000)
	},
	0xc8: func() {
		// RET Z
		if gb.cpu.Z() {
			gb.cpu.ret()
			gb.cpu.clockCycles += 12
		}
	},
	0xc9: func() {
		//  RET
		gb.cpu.ret()
	},
	0xca: func() {
		// JP Z, a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())

		if gb.cpu.Z() {
			gb.cpu.jump(joinu8(msb, lsb))
			gb.cpu.clockCycles += 4
		}
	},
	0xcc: func() {
		//  CALL Z, a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())

		if gb.cpu.Z() {
			addr := joinu8(msb, lsb)
			gb.cpu.call(addr)
			gb.cpu.clockCycles += 12
		}
	},
	0xcd: func() {
		// CALL a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		gb.cpu.call(addr)
	},
	0xce: func() {
		// ADC A, d8
		gb.cpu.adc8Reg(REG_A, gb.bus.read(gb.cpu.readPc()))
	},
	0xcf: func() {
		// RST 08H
		gb.cpu.call(0x0008)
	},

	0xd0: func() {
		// RET NC
		if !gb.cpu.C() {
			gb.cpu.ret()
			gb.cpu.clockCycles += 12
		}
	},
	0xd1: func() {
		// POP DE
		gb.cpu.popSp(REG_DE)
	},
	0xd2: func() {
		// JP NC, a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())

		if !gb.cpu.C() {
			gb.cpu.jump(joinu8(msb, lsb))
			gb.cpu.clockCycles += 4
		}
	},
	0xd4: func() {
		// CALL NC, a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())

		if !gb.cpu.C() {
			gb.cpu.call(joinu8(msb, lsb))
			gb.cpu.clockCycles += 12
		}
	},
	0xd5: func() {
		// PUSH DE
		gb.cpu.pushSp(REG_DE)
	},
	0xd6: func() {
		// SUB d8
		gb.cpu.sub8Reg(REG_A, gb.bus.read(gb.cpu.readPc()))
	},
	0xd7: func() {
		// RST 10H
		gb.cpu.call(0x0010)
	},
	0xd8: func() {
		// RET C
		if gb.cpu.C() {
			gb.cpu.ret()
			gb.cpu.clockCycles += 12
		}
	},
	0xd9: func() {
		// RETI
		gb.cpu.enablingIme = true
		gb.cpu.ret()
	},
	0xda: func() {
		// JP C, a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())

		if gb.cpu.C() {
			gb.cpu.jump(joinu8(msb, lsb))
			gb.cpu.clockCycles += 4
		}
	},
	0xdc: func() {
		// CALL C, a16
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())

		if gb.cpu.C() {
			addr := joinu8(msb, lsb)
			gb.cpu.call(addr)
			gb.cpu.clockCycles += 12
		}
	},
	0xde: func() {
		// SBC A, d8
		gb.cpu.sbc8Reg(REG_A, gb.bus.read(gb.cpu.readPc()))
	},
	0xdf: func() {
		// RST 18H
		gb.cpu.call(0x0018)
	},

	0xe0: func() {
		// LDH (a8), A
		addr := 0xff00 + uint16(gb.bus.read(gb.cpu.readPc()))
		gb.bus.write(addr, gb.cpu.read8Reg(REG_A))
	},
	0xe1: func() {
		// POP HL
		gb.cpu.popSp(REG_HL)
	},
	0xe2: func() {
		// LD (C), A
		addr := joinu8(0xff, gb.cpu.read8Reg(REG_C))
		val := gb.cpu.read8Reg(REG_A)
		gb.bus.write(addr, val)
	},
	0xe5: func() {
		// PUSH HL
		gb.cpu.pushSp(REG_HL)
	},
	0xe6: func() {
		// AND d8
		addr := gb.cpu.readPc()
		val := gb.bus.read(addr)
		gb.cpu.and8Reg(REG_A, val)
	},
	0xe7: func() {
		// RST 20H
		gb.cpu.call(0x0020)
	},
	0xe8: func() {
		// ADD SP, r8
		sp := gb.cpu.read16Reg(REG_SP)
		r8 := int8(gb.bus.read(gb.cpu.readPc()))
		add := sp + uint16(r8)
		gb.cpu.set16Reg(REG_SP, add)

		carry := sp ^ uint16(r8) ^ add
		gb.cpu.setZ(false)
		gb.cpu.setN(false)
		gb.cpu.setH(carry&0x10 == 0x10)
		gb.cpu.setC(carry&0x100 == 0x100)
	},
	0xe9: func() {
		// JP HL
		addr := gb.cpu.read16Reg(REG_HL)
		gb.cpu.jump(addr)
	},
	0xea: func() {
		// LD (a16), A
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		gb.bus.write(addr, gb.cpu.read8Reg(REG_A))
	},
	0xee: func() {
		// XOR d8
		gb.cpu.xor8Reg(REG_A, gb.bus.read(gb.cpu.readPc()))
	},
	0xef: func() {
		// RST 28H
		gb.cpu.call(0x0028)
	},

	0xf0: func() {
		// LDH A, (a8)
		addr := 0xff00 + uint16(gb.bus.read(gb.cpu.readPc()))
		gb.cpu.set8Reg(REG_A, gb.bus.read(addr))
	},
	0xf1: func() {
		// POP AF
		gb.cpu.popSp(REG_AF)
		gb.cpu.set16Reg(REG_AF, gb.cpu.read16Reg(REG_AF)&0xfff0)
	},
	0xf2: func() {
		// LD A, (C)
		lsb := gb.cpu.read8Reg(REG_C)
		addr := joinu8(0xff, lsb)
		val := gb.bus.read(addr)
		gb.cpu.set8Reg(REG_A, val)
	},
	0xf3: func() {
		// DI
		gb.cpu.enablingIme = false
	},
	0xf5: func() {
		// PUSH AF
		gb.cpu.pushSp(REG_AF)
	},
	0xf6: func() {
		// OR d8
		gb.cpu.or8Reg(REG_A, gb.bus.read(gb.cpu.readPc()))
	},
	0xf7: func() {
		// RST 30H
		gb.cpu.call(0x0030)
	},
	0xf8: func() {
		//  LD HL, SP + r8
		r8 := int8(gb.bus.read(gb.cpu.readPc()))
		add := int32(gb.cpu.sp) + int32(r8)
		gb.cpu.set16Reg(REG_HL, uint16(add))

		carry := gb.cpu.sp ^ uint16(r8) ^ uint16(add)
		gb.cpu.setZ(false)
		gb.cpu.setN(false)
		gb.cpu.setH(carry&0x10 == 0x10)
		gb.cpu.setC(carry&0x100 == 0x100)
	},
	0xf9: func() {
		// LD SP, HL
		gb.cpu.set16Reg(REG_SP, gb.cpu.read16Reg(REG_HL))
	},
	0xfa: func() {
		// LD A, (a16)
		lsb := gb.bus.read(gb.cpu.readPc())
		msb := gb.bus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		gb.cpu.set8Reg(REG_A, gb.bus.read(addr))
	},
	0xfb: func() {
		// IE
		gb.cpu.enablingIme = true
	},
	0xfe: func() {
		gb.cpu.cp8Reg(REG_A, gb.bus.read(gb.cpu.readPc()))
	},
	0xff: func() {
		// RST 38H
		gb.cpu.call(0x0038)
	},
}
