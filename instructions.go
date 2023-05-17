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

var instructions = [0x100]func(gb *Gameboy){
	0x00: func(gb *Gameboy) {}, // NOP
	0x01: func(gb *Gameboy) {
		// LD BC, d16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		gb.cpu.bc.Set(joinu8(msb, lsb))
	},
	0x02: func(gb *Gameboy) {
		// LD (BC), A
		gb.memoryBus.write(gb.cpu.bc.val, gb.cpu.af.Hi())
	},
	0x03: func(gb *Gameboy) {
		// INC BC
		gb.cpu.bc.val++
	},
	0x04: func(gb *Gameboy) {
		// INC B
		b := inc(gb, gb.cpu.bc.Hi())
		gb.cpu.bc.SetHi(b)
	},
	0x05: func(gb *Gameboy) {
		// DEC B
		b := dec(gb, gb.cpu.bc.Hi())
		gb.cpu.bc.SetHi(b)
	},
	0x06: func(gb *Gameboy) {
		// LD B, d8
		val := gb.memoryBus.read(gb.cpu.readPc())
		gb.cpu.bc.SetHi(val)
	},
	0x07: func(gb *Gameboy) {
		// RLCA
		rlca(gb)
	},
	0x08: func(gb *Gameboy) {
		// LD (a16), SP
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		gb.memoryBus.write(addr, lo(gb.cpu.sp))
		gb.memoryBus.write(addr+1, hi(gb.cpu.sp))
	},
	0x09: func(gb *Gameboy) {
		// ADD HL, BC
		add := add16(gb, gb.cpu.hl.val, gb.cpu.bc.val)
		gb.cpu.hl.Set(add)
	},
	0x0a: func(gb *Gameboy) {
		// LD A, (BC)
		val := gb.memoryBus.read(gb.cpu.bc.val)
		gb.cpu.af.SetHi(val)
	},
	0x0b: func(gb *Gameboy) {
		// DEC BC
		gb.cpu.bc.val--
	},
	0x0c: func(gb *Gameboy) {
		// INC D
		c := inc(gb, gb.cpu.bc.Lo())
		gb.cpu.bc.SetLo(c)
	},
	0x0d: func(gb *Gameboy) {
		// DEC C
		c := dec(gb, gb.cpu.bc.Lo())
		gb.cpu.bc.SetLo(c)
	},
	0x0e: func(gb *Gameboy) {
		// LD C, d8
		gb.cpu.bc.SetLo(gb.memoryBus.read(gb.cpu.readPc()))
	},
	0x0f: func(gb *Gameboy) {
		// RRCA
		rrca(gb)
	},
	0x10: func(gb *Gameboy) {
		// STOP d8
		gb.timer.resetDIV()
	},
	0x11: func(gb *Gameboy) {
		// LD DE, d16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		gb.cpu.de.Set(joinu8(msb, lsb))
	},
	0x12: func(gb *Gameboy) {
		// LD (DE), A
		gb.memoryBus.write(gb.cpu.de.val, gb.cpu.af.Hi())
	},
	0x13: func(gb *Gameboy) {
		// INC DE
		gb.cpu.de.val++
	},
	0x14: func(gb *Gameboy) {
		// INC D
		d := inc(gb, gb.cpu.de.Hi())
		gb.cpu.de.SetHi(d)
	},
	0x15: func(gb *Gameboy) {
		// DEC D
		d := dec(gb, gb.cpu.de.Lo())
		gb.cpu.de.SetHi(d)
	},
	0x16: func(gb *Gameboy) {
		// LD D, d8
		gb.cpu.de.SetHi(gb.memoryBus.read(gb.cpu.readPc()))
	},
	0x17: func(gb *Gameboy) {
		// RLA
		rla(gb)
	},
	0x18: func(gb *Gameboy) {
		// JR r8
		val := int8(gb.memoryBus.read(gb.cpu.readPc()))
		pc := int32(gb.cpu.pc)
		addr := uint16(pc + int32(val))
		jump(gb, addr)
	},
	0x19: func(gb *Gameboy) {
		// ADD HL, DE
		add := add16(gb, gb.cpu.hl.val, gb.cpu.de.val)
		gb.cpu.hl.Set(add)
	},
	0x1a: func(gb *Gameboy) {
		// LD A, (DE)
		gb.cpu.af.SetHi(gb.memoryBus.read(gb.cpu.de.val))
	},
	0x1b: func(gb *Gameboy) {
		// DEC DE
		gb.cpu.de.val--
	},
	0x1c: func(gb *Gameboy) {
		// INC E
		e := inc(gb, gb.cpu.de.Lo())
		gb.cpu.de.SetLo(e)
	},
	0x1d: func(gb *Gameboy) {
		// DEC E
		e := dec(gb, gb.cpu.de.Lo())
		gb.cpu.de.SetLo(e)
	},
	0x1e: func(gb *Gameboy) {
		// LD E, d8
		val := gb.memoryBus.read(gb.cpu.readPc())
		gb.cpu.de.SetLo(val)
	},
	0x1f: func(gb *Gameboy) {
		// RRA
		reg := gb.cpu.af.Hi()
		rot := rotateRight(reg, 1)
		rot = toggleBit(rot, 7, gb.cpu.C())
		gb.cpu.af.SetHi(rot)

		gb.cpu.setC(isBitSet(reg, 0))
		gb.cpu.setH(false)
		gb.cpu.setN(false)
		gb.cpu.setZ(false)
	},

	0x20: func(gb *Gameboy) {
		// JR NZ, r8
		addr := gb.cpu.readPc()
		val := int8(gb.memoryBus.read(addr))

		if !gb.cpu.Z() {
			addr := int32(gb.cpu.pc) + int32(val)
			jump(gb, uint16(addr))
			gb.cpu.clockCycles += 4
		}
	},
	0x21: func(gb *Gameboy) {
		// LD HL, d16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		gb.cpu.hl.Set(joinu8(msb, lsb))
	},
	0x22: func(gb *Gameboy) {
		// LD (HL+), A
		addr := gb.cpu.hl.val
		gb.memoryBus.write(addr, gb.cpu.af.Hi())
		gb.cpu.hl.Set(addr + 1)
	},
	0x23: func(gb *Gameboy) {
		// INC HL
		gb.cpu.hl.val++
	},
	0x24: func(gb *Gameboy) {
		// INC H
		h := inc(gb, gb.cpu.hl.Hi())
		gb.cpu.hl.SetHi(h)
	},
	0x25: func(gb *Gameboy) {
		// DEC H
		h := dec(gb, gb.cpu.hl.Hi())
		gb.cpu.hl.SetHi(h)
	},
	0x26: func(gb *Gameboy) {
		// LD H, d8
		gb.cpu.hl.SetHi(gb.memoryBus.read(gb.cpu.readPc()))
	},
	0x27: func(gb *Gameboy) {
		// DAA
		a := gb.cpu.af.Hi()

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

		gb.cpu.af.SetHi(a)
	},
	0x28: func(gb *Gameboy) {
		// JR Z, r8
		addr := gb.cpu.readPc()
		val := int8(gb.memoryBus.read(addr))

		if gb.cpu.Z() {
			addr := int32(gb.cpu.pc) + int32(val)
			jump(gb, uint16(addr))
			gb.cpu.clockCycles += 4
		}
	},
	0x29: func(gb *Gameboy) {
		// ADD HL, HL
		add := add16(gb, gb.cpu.hl.val, gb.cpu.hl.val)
		gb.cpu.hl.Set(add)
	},
	0x2a: func(gb *Gameboy) {
		// LD A, (HL+)
		val := gb.memoryBus.read(gb.cpu.hl.val)
		gb.cpu.af.SetHi(val)
		gb.cpu.hl.val++
	},
	0x2b: func(gb *Gameboy) {
		// DEC HL
		gb.cpu.hl.val--
	},
	0x2c: func(gb *Gameboy) {
		// INC L
		l := inc(gb, gb.cpu.hl.Lo())
		gb.cpu.hl.SetLo(l)
	},
	0x2d: func(gb *Gameboy) {
		// DEC L
		l := dec(gb, gb.cpu.hl.Lo())
		gb.cpu.hl.SetLo(l)
	},
	0x2e: func(gb *Gameboy) {
		// LD L, d8
		addr := gb.cpu.readPc()
		val := gb.memoryBus.read(addr)
		gb.cpu.hl.SetHi(val)
	},
	0x2f: func(gb *Gameboy) {
		// CPL
		val := gb.cpu.af.Hi()
		gb.cpu.af.SetHi(0xff ^ val)
		gb.cpu.setN(true)
		gb.cpu.setH(true)
	},

	0x30: func(gb *Gameboy) {
		// JR NC, r8
		pc := gb.cpu.readPc()

		if !gb.cpu.C() {
			val := int8(gb.memoryBus.read(pc))
			addr := int32(gb.cpu.pc) + int32(val)
			jump(gb, uint16(addr))
			gb.cpu.clockCycles += 4
		}
	},
	0x31: func(gb *Gameboy) {
		// LD SP, d16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		gb.cpu.sp = joinu8(msb, lsb)
	},
	0x32: func(gb *Gameboy) {
		// LD (HL-), A
		addr := gb.cpu.hl.val
		gb.memoryBus.write(addr, gb.cpu.af.Hi())
		gb.cpu.hl.val--
	},
	0x33: func(gb *Gameboy) {
		// INC SP
		gb.cpu.sp = gb.cpu.sp + 1
	},
	0x34: func(gb *Gameboy) {
		// INC (HL)
		addr := gb.cpu.hl.val
		val := gb.memoryBus.read(addr)
		inc := val + 1
		gb.memoryBus.write(addr, inc)

		gb.cpu.setN(false)
		gb.cpu.setH((val&0xf)+1 > 0xf)
		gb.cpu.setZ(inc == 0)
	},
	0x35: func(gb *Gameboy) {
		// DEC (HL)
		addr := gb.cpu.hl.val
		val := gb.memoryBus.read(addr)
		dec := val - 1
		gb.memoryBus.write(addr, dec)

		gb.cpu.setN(true)
		gb.cpu.setH((val & 0xf) == 0)
		gb.cpu.setZ(dec == 0)
	},
	0x36: func(gb *Gameboy) {
		// LD (HL), d8
		addr := gb.cpu.hl.val
		val := gb.memoryBus.read(gb.cpu.readPc())
		gb.memoryBus.write(addr, val)
	},
	0x37: func(gb *Gameboy) {
		// SCF
		gb.cpu.setN(false)
		gb.cpu.setH(false)
		gb.cpu.setC(true)
	},
	0x38: func(gb *Gameboy) {
		// JR C, r8
		pc := gb.cpu.readPc()

		if gb.cpu.C() {
			val := int8(gb.memoryBus.read(pc))
			addr := int32(gb.cpu.pc) + int32(val)
			jump(gb, uint16(addr))
			gb.cpu.clockCycles += 4
		}
	},
	0x39: func(gb *Gameboy) {
		// ADD HL, SP
		gb.cpu.hl.Set(gb.cpu.sp)
	},
	0x3a: func(gb *Gameboy) {
		// LD A, (HL-)
		addr := gb.cpu.hl.val
		val := gb.memoryBus.read(addr)
		gb.cpu.af.SetHi(val)
		gb.cpu.hl.val--
	},
	0x3b: func(gb *Gameboy) {
		// DEC SP
		gb.cpu.sp = gb.cpu.sp - 1
	},
	0x3c: func(gb *Gameboy) {
		// INC A
		a := inc(gb, gb.cpu.af.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x3d: func(gb *Gameboy) {
		// DEC A
		a := dec(gb, gb.cpu.af.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x3e: func(gb *Gameboy) {
		// LD A, d8
		gb.cpu.af.SetHi(gb.memoryBus.read(gb.cpu.readPc()))
	},
	0x3f: func(gb *Gameboy) {
		// CCF
		gb.cpu.setN(false)
		gb.cpu.setH(false)
		gb.cpu.setC(!gb.cpu.C())
	},
	0x40: func(gb *Gameboy) {
		// LD B, B
		gb.cpu.bc.SetHi(gb.cpu.bc.Hi())
	},
	0x41: func(gb *Gameboy) {
		// LD B, C
		gb.cpu.bc.SetHi(gb.cpu.bc.Lo())
	},
	0x42: func(gb *Gameboy) {
		// LD B, D
		gb.cpu.bc.SetHi(gb.cpu.de.Hi())
	},
	0x43: func(gb *Gameboy) {
		// LD B, E
		gb.cpu.bc.SetHi(gb.cpu.de.Lo())
	},
	0x44: func(gb *Gameboy) {
		// LD B, H
		gb.cpu.bc.SetHi(gb.cpu.hl.Hi())
	},
	0x45: func(gb *Gameboy) {
		// LD B, L
		gb.cpu.bc.SetHi(gb.cpu.hl.Lo())
	},
	0x46: func(gb *Gameboy) {
		// LD B, (HL)
		gb.cpu.bc.SetHi(gb.memoryBus.read(gb.cpu.hl.val))
	},
	0x47: func(gb *Gameboy) {
		// LD B, A
		gb.cpu.bc.SetHi(gb.cpu.af.Hi())
	},
	0x48: func(gb *Gameboy) {
		// LD C, B
		gb.cpu.bc.SetLo(gb.cpu.bc.Hi())
	},
	0x49: func(gb *Gameboy) {
		// LD C, C
		gb.cpu.bc.SetLo(gb.cpu.bc.Lo())
	},
	0x4a: func(gb *Gameboy) {
		// LD C, D
		gb.cpu.bc.SetLo(gb.cpu.de.Hi())
	},
	0x4b: func(gb *Gameboy) {
		// LD C, E
		gb.cpu.bc.SetLo(gb.cpu.de.Lo())
	},
	0x4c: func(gb *Gameboy) {
		// LD C, H
		gb.cpu.bc.SetLo(gb.cpu.hl.Hi())
	},
	0x4d: func(gb *Gameboy) {
		// LD C, L
		gb.cpu.bc.SetLo(gb.cpu.hl.Lo())
	},
	0x4e: func(gb *Gameboy) {
		// LD C, (HL)
		gb.cpu.bc.SetLo(gb.memoryBus.read(gb.cpu.hl.val))
	},
	0x4f: func(gb *Gameboy) {
		// LD C, A
		gb.cpu.bc.SetLo(gb.cpu.af.Hi())
	},
	0x50: func(gb *Gameboy) {
		// LD D, B
		gb.cpu.de.SetHi(gb.cpu.bc.Hi())
	},
	0x51: func(gb *Gameboy) {
		// LD D, C
		gb.cpu.de.SetHi(gb.cpu.bc.Lo())
	},
	0x52: func(gb *Gameboy) {
		// LD D, D
		gb.cpu.de.SetHi(gb.cpu.de.Hi())
	},
	0x53: func(gb *Gameboy) {
		// LD D, E
		gb.cpu.de.SetHi(gb.cpu.de.Lo())
	},
	0x54: func(gb *Gameboy) {
		// LD D, H
		gb.cpu.de.SetHi(gb.cpu.hl.Hi())
	},
	0x55: func(gb *Gameboy) {
		// LD D, L
		gb.cpu.de.SetHi(gb.cpu.hl.Lo())
	},
	0x56: func(gb *Gameboy) {
		// LD D, (HL)
		gb.cpu.de.SetHi(gb.memoryBus.read(gb.cpu.hl.val))
	},
	0x57: func(gb *Gameboy) {
		// LD D, A
		gb.cpu.de.SetHi(gb.cpu.af.Hi())
	},
	0x58: func(gb *Gameboy) {
		// LD E, B
		gb.cpu.de.SetHi(gb.cpu.bc.Hi())
	},
	0x59: func(gb *Gameboy) {
		// LD E, C
		gb.cpu.de.SetHi(gb.cpu.bc.Lo())
	},
	0x5a: func(gb *Gameboy) {
		// LD E, D
		gb.cpu.de.SetHi(gb.cpu.de.Hi())
	},
	0x5b: func(gb *Gameboy) {
		// LD E, E
		gb.cpu.de.SetHi(gb.cpu.de.Hi())
	},
	0x5c: func(gb *Gameboy) {
		// LD E, H
		gb.cpu.de.SetHi(gb.cpu.hl.Hi())
	},
	0x5d: func(gb *Gameboy) {
		// LD E, L
		gb.cpu.de.SetHi(gb.cpu.hl.Lo())
	},
	0x5e: func(gb *Gameboy) {
		// LD E, (HL)
		gb.cpu.de.SetLo(gb.memoryBus.read(gb.cpu.hl.val))
	},
	0x5f: func(gb *Gameboy) {
		// LD E, A
		gb.cpu.de.SetLo(gb.cpu.af.Hi())
	},
	0x60: func(gb *Gameboy) {
		// LD H, B
		gb.cpu.hl.SetHi(gb.cpu.bc.Hi())
	},
	0x61: func(gb *Gameboy) {
		// LD H, C
		gb.cpu.hl.SetHi(gb.cpu.bc.Lo())
	},
	0x62: func(gb *Gameboy) {
		// LD H, D
		gb.cpu.hl.SetHi(gb.cpu.de.Hi())
	},
	0x63: func(gb *Gameboy) {
		// LD H, E
		gb.cpu.hl.SetHi(gb.cpu.de.Lo())
	},
	0x64: func(gb *Gameboy) {
		// LD H, H
		gb.cpu.hl.SetHi(gb.cpu.hl.Hi())
	},
	0x65: func(gb *Gameboy) {
		// LD H, L
		gb.cpu.hl.SetHi(gb.cpu.hl.Lo())
	},
	0x66: func(gb *Gameboy) {
		// LD H, (HL)
		gb.cpu.hl.SetHi(gb.memoryBus.read(gb.cpu.hl.val))
	},
	0x67: func(gb *Gameboy) {
		// LD H, A
		gb.cpu.hl.SetHi(gb.cpu.af.Hi())
	},
	0x68: func(gb *Gameboy) {
		// LD L, B
		gb.cpu.hl.SetLo(gb.cpu.bc.Hi())
	},
	0x69: func(gb *Gameboy) {
		// LD L, C
		gb.cpu.hl.SetLo(gb.cpu.bc.Lo())
	},
	0x6a: func(gb *Gameboy) {
		// LD L, D
		gb.cpu.hl.SetLo(gb.cpu.de.Hi())
	},
	0x6b: func(gb *Gameboy) {
		// LD L, E
		gb.cpu.hl.SetLo(gb.cpu.de.Lo())
	},
	0x6c: func(gb *Gameboy) {
		// LD L, H
		gb.cpu.hl.SetLo(gb.cpu.hl.Hi())
	},
	0x6d: func(gb *Gameboy) {
		// LD L, L
		gb.cpu.hl.SetLo(gb.cpu.hl.Lo())
	},
	0x6e: func(gb *Gameboy) {
		// LD L, (HL)
		gb.cpu.hl.SetLo(gb.memoryBus.read(gb.cpu.hl.val))
	},
	0x6f: func(gb *Gameboy) {
		// LD L, A
		gb.cpu.hl.SetLo(gb.cpu.af.Hi())
	},
	0x70: func(gb *Gameboy) {
		// LD (HL), B
		gb.memoryBus.write(gb.cpu.hl.val, gb.cpu.bc.Hi())
	},
	0x71: func(gb *Gameboy) {
		// LD (HL), C
		gb.memoryBus.write(gb.cpu.hl.val, gb.cpu.bc.Lo())
	},
	0x72: func(gb *Gameboy) {
		// LD (HL), D
		gb.memoryBus.write(gb.cpu.hl.val, gb.cpu.de.Hi())
	},
	0x73: func(gb *Gameboy) {
		// LD (HL), E
		gb.memoryBus.write(gb.cpu.hl.val, gb.cpu.de.Lo())
	},
	0x74: func(gb *Gameboy) {
		// LD (HL), H
		gb.memoryBus.write(gb.cpu.hl.val, gb.cpu.hl.Hi())
	},
	0x75: func(gb *Gameboy) {
		// LD (HL), L
		gb.memoryBus.write(gb.cpu.hl.val, gb.cpu.hl.Lo())
	},
	0x76: func(gb *Gameboy) {
		// HALT
		gb.cpu.halted = true
	},
	0x77: func(gb *Gameboy) {
		// LD (HL), A
		gb.memoryBus.write(gb.cpu.hl.val, gb.cpu.af.Hi())
	},
	0x78: func(gb *Gameboy) {
		// LD A, B
		gb.cpu.af.SetHi(gb.cpu.bc.Hi())
	},
	0x79: func(gb *Gameboy) {
		// LD A, C
		gb.cpu.af.SetHi(gb.cpu.bc.Lo())
	},
	0x7a: func(gb *Gameboy) {
		// LD A, D
		gb.cpu.af.SetHi(gb.cpu.de.Hi())
	},
	0x7b: func(gb *Gameboy) {
		// LD A, E
		gb.cpu.af.SetHi(gb.cpu.de.Lo())
	},
	0x7c: func(gb *Gameboy) {
		// LD A, H
		gb.cpu.af.SetHi(gb.cpu.hl.Hi())
	},
	0x7d: func(gb *Gameboy) {
		// LD A, L
		gb.cpu.af.SetHi(gb.cpu.hl.Lo())
	},
	0x7e: func(gb *Gameboy) {
		// LD A, (HL)
		gb.cpu.af.SetHi(gb.memoryBus.read(gb.cpu.hl.val))
	},
	0x7f: func(gb *Gameboy) {
		// LD A, A
		gb.cpu.af.SetHi(gb.cpu.af.Hi())
	},
	0x80: func(gb *Gameboy) {
		// ADD A, B
		a := add8(gb, gb.cpu.af.Hi(), gb.cpu.bc.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x81: func(gb *Gameboy) {
		// ADD A, C
		a := add8(gb, gb.cpu.af.Hi(), gb.cpu.bc.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x82: func(gb *Gameboy) {
		// ADD A, D
		a := add8(gb, gb.cpu.af.Hi(), gb.cpu.de.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x83: func(gb *Gameboy) {
		// ADD A, E
		a := add8(gb, gb.cpu.af.Hi(), gb.cpu.de.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x84: func(gb *Gameboy) {
		// ADD A, H
		a := add8(gb, gb.cpu.af.Hi(), gb.cpu.hl.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x85: func(gb *Gameboy) {
		// ADD A, L
		a := add8(gb, gb.cpu.af.Hi(), gb.cpu.hl.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x86: func(gb *Gameboy) {
		// ADD A, HL
		a := add8(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.hl.val))
		gb.cpu.af.SetHi(a)
	},
	0x87: func(gb *Gameboy) {
		// ADD A, A
		a := add8(gb, gb.cpu.af.Hi(), gb.cpu.af.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x88: func(gb *Gameboy) {
		// ADC A, B
		a := adc(gb, gb.cpu.af.Hi(), gb.cpu.bc.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x89: func(gb *Gameboy) {
		// ADC A, C
		a := adc(gb, gb.cpu.af.Hi(), gb.cpu.bc.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x8a: func(gb *Gameboy) {
		// ADC A, D
		a := adc(gb, gb.cpu.af.Hi(), gb.cpu.de.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x8b: func(gb *Gameboy) {
		// ADC A, E
		a := adc(gb, gb.cpu.af.Hi(), gb.cpu.de.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x8c: func(gb *Gameboy) {
		// ADC A, H
		a := adc(gb, gb.cpu.af.Hi(), gb.cpu.hl.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x8d: func(gb *Gameboy) {
		// ADC A, L
		a := adc(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.hl.val))
		gb.cpu.af.SetHi(a)
	},
	0x8e: func(gb *Gameboy) {
		// ADC A, HL
		a := adc(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.hl.val))
		gb.cpu.af.SetHi(a)
	},
	0x8f: func(gb *Gameboy) {
		// ADC A, A
		a := adc(gb, gb.cpu.af.Hi(), gb.cpu.af.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x90: func(gb *Gameboy) {
		// SUB B
		a := sub(gb, gb.cpu.af.Hi(), gb.cpu.bc.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x91: func(gb *Gameboy) {
		// SUB C
		a := sub(gb, gb.cpu.af.Hi(), gb.cpu.bc.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x92: func(gb *Gameboy) {
		// SUB D
		a := sub(gb, gb.cpu.af.Hi(), gb.cpu.de.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x93: func(gb *Gameboy) {
		// SUB E
		a := sub(gb, gb.cpu.af.Hi(), gb.cpu.de.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x94: func(gb *Gameboy) {
		// SUB H
		a := sub(gb, gb.cpu.af.Hi(), gb.cpu.hl.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x95: func(gb *Gameboy) {
		// SUB L
		a := sub(gb, gb.cpu.af.Hi(), gb.cpu.hl.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x96: func(gb *Gameboy) {
		// SUB (HL)
		a := sub(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.hl.val))
		gb.cpu.af.SetHi(a)
	},
	0x97: func(gb *Gameboy) {
		// SUB A
		a := sub(gb, gb.cpu.af.Hi(), gb.cpu.af.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x98: func(gb *Gameboy) {
		// SBC A, B
		a := sbc(gb, gb.cpu.af.Hi(), gb.cpu.bc.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x99: func(gb *Gameboy) {
		// SBC A, C
		a := sbc(gb, gb.cpu.af.Hi(), gb.cpu.bc.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x9a: func(gb *Gameboy) {
		// SBC A, D
		a := sbc(gb, gb.cpu.af.Hi(), gb.cpu.de.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x9b: func(gb *Gameboy) {
		// SBC A, E
		a := sbc(gb, gb.cpu.af.Hi(), gb.cpu.de.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x9c: func(gb *Gameboy) {
		// SBC A, H
		a := sbc(gb, gb.cpu.af.Hi(), gb.cpu.hl.Hi())
		gb.cpu.af.SetHi(a)
	},
	0x9d: func(gb *Gameboy) {
		// SBC A, L
		a := sbc(gb, gb.cpu.af.Hi(), gb.cpu.hl.Lo())
		gb.cpu.af.SetHi(a)
	},
	0x9e: func(gb *Gameboy) {
		// SBC A, (HL)
		a := sbc(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.hl.val))
		gb.cpu.af.SetHi(a)
	},
	0x9f: func(gb *Gameboy) {
		// SBC A, A
		a := sbc(gb, gb.cpu.af.Hi(), gb.cpu.af.Hi())
		gb.cpu.af.SetHi(a)
	},
	0xa0: func(gb *Gameboy) {
		// AND B
		and(gb, gb.cpu.bc.Hi())
	},
	0xa1: func(gb *Gameboy) {
		// AND C
		and(gb, gb.cpu.bc.Lo())
	},
	0xa2: func(gb *Gameboy) {
		// AND D
		and(gb, gb.cpu.de.Hi())
	},
	0xa3: func(gb *Gameboy) {
		// AND E
		and(gb, gb.cpu.de.Lo())
	},
	0xa4: func(gb *Gameboy) {
		// AND H
		and(gb, gb.cpu.hl.Hi())
	},
	0xa5: func(gb *Gameboy) {
		// AND L
		and(gb, gb.cpu.hl.Lo())
	},
	0xa6: func(gb *Gameboy) {
		// AND HL
		and(gb, gb.memoryBus.read(gb.cpu.hl.val))
	},
	0xa7: func(gb *Gameboy) {
		// AND A
		and(gb, gb.cpu.af.Hi())
	},
	0xa8: func(gb *Gameboy) {
		// XOR B
		xor(gb, gb.cpu.bc.Hi())
	},
	0xa9: func(gb *Gameboy) {
		// XOR C
		xor(gb, gb.cpu.bc.Lo())
	},
	0xaa: func(gb *Gameboy) {
		// XOR D
		xor(gb, gb.cpu.de.Hi())
	},
	0xab: func(gb *Gameboy) {
		// XOR E
		xor(gb, gb.cpu.de.Lo())
	},
	0xac: func(gb *Gameboy) {
		// XOR H
		xor(gb, gb.cpu.hl.Hi())
	},
	0xad: func(gb *Gameboy) {
		// XOR L
		xor(gb, gb.cpu.hl.Lo())
	},
	0xae: func(gb *Gameboy) {
		// XOR HL
		xor(gb, gb.memoryBus.read(gb.cpu.hl.val))
	},
	0xaf: func(gb *Gameboy) {
		// XOR A
		xor(gb, gb.cpu.af.Hi())
	},
	0xb0: func(gb *Gameboy) {
		// OR B
		or(gb, gb.cpu.bc.Hi())
	},
	0xb1: func(gb *Gameboy) {
		// OR C
		or(gb, gb.cpu.bc.Lo())
	},
	0xb2: func(gb *Gameboy) {
		// OR D
		or(gb, gb.cpu.de.Hi())
	},
	0xb3: func(gb *Gameboy) {
		// OR E
		or(gb, gb.cpu.de.Lo())
	},
	0xb4: func(gb *Gameboy) {
		// OR H
		or(gb, gb.cpu.hl.Hi())
	},
	0xb5: func(gb *Gameboy) {
		// OR L
		or(gb, gb.cpu.hl.Lo())
	},
	0xb6: func(gb *Gameboy) {
		// OR HL
		or(gb, gb.memoryBus.read(gb.cpu.hl.val))
	},
	0xb7: func(gb *Gameboy) {
		// OR A
		or(gb, gb.cpu.af.Hi())
	},
	0xb8: func(gb *Gameboy) {
		// CP B
		cp(gb, gb.cpu.bc.Hi())
	},
	0xb9: func(gb *Gameboy) {
		// CP C
		cp(gb, gb.cpu.bc.Lo())
	},
	0xba: func(gb *Gameboy) {
		// CP D
		cp(gb, gb.cpu.de.Hi())
	},
	0xbb: func(gb *Gameboy) {
		// CP E
		cp(gb, gb.cpu.de.Lo())
	},
	0xbc: func(gb *Gameboy) {
		// CP H
		cp(gb, gb.cpu.hl.Hi())
	},
	0xbd: func(gb *Gameboy) {
		// CP L
		cp(gb, gb.cpu.hl.Lo())
	},
	0xbe: func(gb *Gameboy) {
		// CP (HL)
		cp(gb, gb.memoryBus.read(gb.cpu.hl.val))
	},
	0xbf: func(gb *Gameboy) {
		// CP A
		cp(gb, gb.cpu.af.Hi())
	},

	0xc0: func(gb *Gameboy) {
		// RET NZ
		if !gb.cpu.Z() {
			ret(gb)
			gb.cpu.clockCycles += 12
		}
	},
	0xc1: func(gb *Gameboy) {
		// POP BC
		gb.cpu.bc.Set(gb.popSp())
	},
	0xc2: func(gb *Gameboy) {
		// JP NZ, a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)

		if !gb.cpu.Z() {
			jump(gb, addr)
			gb.cpu.clockCycles += 4
		}
	},
	0xc3: func(gb *Gameboy) {
		// JP a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		jump(gb, addr)
	},
	0xc4: func(gb *Gameboy) {
		// CALL NZ, a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())

		if !gb.cpu.Z() {
			addr := joinu8(msb, lsb)
			jump(gb, addr)
			gb.cpu.clockCycles += 12
		}
	},
	0xc5: func(gb *Gameboy) {
		// PUSH BC
		gb.pushSp(gb.cpu.bc.val)
	},
	0xc6: func(gb *Gameboy) {
		// ADD A, d8
		a := add8(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.readPc()))
		gb.cpu.af.SetHi(a)
	},
	0xc7: func(gb *Gameboy) {
		// RST 00H
		call(gb, 0x0000)
	},
	0xc8: func(gb *Gameboy) {
		// RET Z
		if gb.cpu.Z() {
			ret(gb)
			gb.cpu.clockCycles += 12
		}
	},
	0xc9: func(gb *Gameboy) {
		//  RET
		ret(gb)
	},
	0xca: func(gb *Gameboy) {
		// JP Z, a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())

		if gb.cpu.Z() {
			jump(gb, joinu8(msb, lsb))
			gb.cpu.clockCycles += 4
		}
	},
	0xcc: func(gb *Gameboy) {
		//  CALL Z, a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())

		if gb.cpu.Z() {
			addr := joinu8(msb, lsb)
			call(gb, addr)
			gb.cpu.clockCycles += 12
		}
	},
	0xcd: func(gb *Gameboy) {
		// CALL a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		call(gb, addr)
	},
	0xce: func(gb *Gameboy) {
		// ADC A, d8
		a := adc(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.readPc()))
		gb.cpu.af.SetHi(a)
	},
	0xcf: func(gb *Gameboy) {
		// RST 08H
		call(gb, 0x0008)
	},
	0xd0: func(gb *Gameboy) {
		// RET NC
		if !gb.cpu.C() {
			ret(gb)
			gb.cpu.clockCycles += 12
		}
	},
	0xd1: func(gb *Gameboy) {
		// POP DE
		gb.cpu.de.Set(gb.popSp())
	},
	0xd2: func(gb *Gameboy) {
		// JP NC, a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())

		if !gb.cpu.C() {
			jump(gb, joinu8(msb, lsb))
			gb.cpu.clockCycles += 4
		}
	},
	0xd4: func(gb *Gameboy) {
		// CALL NC, a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())

		if !gb.cpu.C() {
			call(gb, joinu8(msb, lsb))
			gb.cpu.clockCycles += 12
		}
	},
	0xd5: func(gb *Gameboy) {
		// PUSH DE
		gb.pushSp(gb.cpu.de.val)
	},
	0xd6: func(gb *Gameboy) {
		// SUB d8
		a := sub(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.readPc()))
		gb.cpu.af.SetHi(a)
	},
	0xd7: func(gb *Gameboy) {
		// RST 10H
		call(gb, 0x0010)
	},
	0xd8: func(gb *Gameboy) {
		// RET C
		if gb.cpu.C() {
			ret(gb)
			gb.cpu.clockCycles += 12
		}
	},
	0xd9: func(gb *Gameboy) {
		// RETI
		gb.interruptBus.enablingIme = true
		ret(gb)
	},
	0xda: func(gb *Gameboy) {
		// JP C, a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())

		if gb.cpu.C() {
			jump(gb, joinu8(msb, lsb))
			gb.cpu.clockCycles += 4
		}
	},
	0xdc: func(gb *Gameboy) {
		// CALL C, a16
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())

		if gb.cpu.C() {
			addr := joinu8(msb, lsb)
			call(gb, addr)
			gb.cpu.clockCycles += 12
		}
	},
	0xde: func(gb *Gameboy) {
		// SBC A, d8
		a := sbc(gb, gb.cpu.af.Hi(), gb.memoryBus.read(gb.cpu.readPc()))
		gb.cpu.af.SetHi(a)
	},
	0xdf: func(gb *Gameboy) {
		// RST 18H
		call(gb, 0x0018)
	},
	0xe0: func(gb *Gameboy) {
		// LDH (a8), A
		addr := 0xff00 + uint16(gb.memoryBus.read(gb.cpu.readPc()))
		gb.memoryBus.write(addr, gb.cpu.af.Hi())
	},
	0xe1: func(gb *Gameboy) {
		// POP HL
		gb.cpu.hl.Set(gb.popSp())
	},
	0xe2: func(gb *Gameboy) {
		// LD (C), A
		addr := joinu8(0xff, gb.cpu.bc.Lo())
		val := gb.cpu.af.Hi()
		gb.memoryBus.write(addr, val)
	},
	0xe5: func(gb *Gameboy) {
		// PUSH HL
		gb.pushSp(gb.cpu.hl.val)
	},
	0xe6: func(gb *Gameboy) {
		// AND d8
		addr := gb.cpu.readPc()
		val := gb.memoryBus.read(addr)
		and(gb, val)
	},
	0xe7: func(gb *Gameboy) {
		// RST 20H
		call(gb, 0x0020)
	},
	0xe8: func(gb *Gameboy) {
		// ADD SP, r8
		sp := gb.cpu.sp
		r8 := int8(gb.memoryBus.read(gb.cpu.readPc()))
		add := sp + uint16(r8)
		gb.cpu.sp = add

		carry := sp ^ uint16(r8) ^ add
		gb.cpu.setZ(false)
		gb.cpu.setN(false)
		gb.cpu.setH(carry&0x10 == 0x10)
		gb.cpu.setC(carry&0x100 == 0x100)
	},
	0xe9: func(gb *Gameboy) {
		// JP HL
		jump(gb, gb.cpu.hl.val)
	},
	0xea: func(gb *Gameboy) {
		// LD (a16), A
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		gb.memoryBus.write(addr, gb.cpu.af.Hi())
	},
	0xee: func(gb *Gameboy) {
		// XOR d8
		xor(gb, gb.memoryBus.read(gb.cpu.readPc()))
	},
	0xef: func(gb *Gameboy) {
		// RST 28H
		call(gb, 0x0028)
	},

	0xf0: func(gb *Gameboy) {
		// LDH A, (a8)
		addr := 0xff00 + uint16(gb.memoryBus.read(gb.cpu.readPc()))
		gb.cpu.af.SetHi(gb.memoryBus.read(addr))
	},
	0xf1: func(gb *Gameboy) {
		// POP AF
		gb.cpu.af.Set(gb.popSp())
		gb.cpu.af.Set(gb.cpu.af.val & 0xfff0)
	},
	0xf2: func(gb *Gameboy) {
		// LD A, (C)
		lsb := gb.cpu.bc.Lo()
		addr := joinu8(0xff, lsb)
		val := gb.memoryBus.read(addr)
		gb.cpu.af.SetHi(val)
	},
	0xf3: func(gb *Gameboy) {
		// DI
		gb.interruptBus.enablingIme = false
	},
	0xf5: func(gb *Gameboy) {
		// PUSH AF
		gb.pushSp(gb.cpu.af.val)
	},
	0xf6: func(gb *Gameboy) {
		// OR d8
		or(gb, gb.memoryBus.read(gb.cpu.readPc()))
	},
	0xf7: func(gb *Gameboy) {
		// RST 30H
		call(gb, 0x0030)
	},
	0xf8: func(gb *Gameboy) {
		//  LD HL, SP + r8
		r8 := int8(gb.memoryBus.read(gb.cpu.readPc()))
		add := int32(gb.cpu.sp) + int32(r8)
		gb.cpu.hl.Set(uint16(add))

		carry := gb.cpu.sp ^ uint16(r8) ^ uint16(add)
		gb.cpu.setZ(false)
		gb.cpu.setN(false)
		gb.cpu.setH(carry&0x10 == 0x10)
		gb.cpu.setC(carry&0x100 == 0x100)
	},
	0xf9: func(gb *Gameboy) {
		// LD SP, HL
		gb.cpu.sp = gb.cpu.hl.val
	},
	0xfa: func(gb *Gameboy) {
		// LD A, (a16)
		lsb := gb.memoryBus.read(gb.cpu.readPc())
		msb := gb.memoryBus.read(gb.cpu.readPc())
		addr := joinu8(msb, lsb)
		gb.cpu.af.SetHi(gb.memoryBus.read(addr))
	},
	0xfb: func(gb *Gameboy) {
		// IE
		gb.interruptBus.enablingIme = true
	},
	0xfe: func(gb *Gameboy) {
		cp(gb, gb.memoryBus.read(gb.cpu.readPc()))
	},
	0xff: func(gb *Gameboy) {
		// RST 38H
		call(gb, 0x0038)
	},
}
