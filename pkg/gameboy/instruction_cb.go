package gameboy

import "github.com/elmarsan/gbgo/pkg/bit"

var cbInstr = [0x100]func(gb *Gameboy){
	0x00: func(gb *Gameboy) {
		// RLC B
		gb.cpu.bc.SetHi(rlc(gb, gb.cpu.bc.Hi()))
	},
	0x01: func(gb *Gameboy) {
		// RLC C
		gb.cpu.bc.SetLo(rlc(gb, gb.cpu.bc.Lo()))
	},
	0x02: func(gb *Gameboy) {
		// RLC D
		gb.cpu.de.SetHi(rlc(gb, gb.cpu.de.Hi()))
	},
	0x03: func(gb *Gameboy) {
		// RLC E
		gb.cpu.de.SetLo(rlc(gb, gb.cpu.de.Lo()))
	},
	0x04: func(gb *Gameboy) {
		// RLC H
		gb.cpu.hl.SetHi(rlc(gb, gb.cpu.hl.Hi()))
	},
	0x05: func(gb *Gameboy) {
		// RLC L
		gb.cpu.hl.SetLo(rlc(gb, gb.cpu.hl.Lo()))
	},
	0x06: func(gb *Gameboy) {
		// RLC (hl)
		hl := gb.cpu.hl.val
		gb.memoryBus.write(hl, rlc(gb, gb.memoryBus.read(hl)))
	},
	0x07: func(gb *Gameboy) {
		// RLC A
		gb.cpu.af.SetHi(rlc(gb, gb.cpu.af.Hi()))
	},
	0x08: func(gb *Gameboy) {
		// RRC B
		gb.cpu.bc.SetHi(rrc(gb, gb.cpu.bc.Hi()))
	},
	0x09: func(gb *Gameboy) {
		// RRC C
		gb.cpu.bc.SetLo(rrc(gb, gb.cpu.bc.Lo()))
	},
	0x0a: func(gb *Gameboy) {
		// RRC D
		gb.cpu.de.SetHi(rrc(gb, gb.cpu.de.Hi()))
	},
	0x0b: func(gb *Gameboy) {
		// RRC E
		gb.cpu.de.SetLo(rrc(gb, gb.cpu.de.Lo()))
	},
	0x0c: func(gb *Gameboy) {
		// RRC H
		gb.cpu.hl.SetHi(rrc(gb, gb.cpu.hl.Hi()))
	},
	0x0d: func(gb *Gameboy) {
		// RRC L
		gb.cpu.hl.SetLo(rrc(gb, gb.cpu.hl.Lo()))
	},
	0x0e: func(gb *Gameboy) {
		// RRC (hl)
		hl := gb.cpu.hl.val
		gb.memoryBus.write(hl, rrc(gb, gb.memoryBus.read(hl)))
	},
	0x0f: func(gb *Gameboy) {
		// RRC A
		gb.cpu.af.SetHi(rrc(gb, gb.cpu.af.Hi()))
	},
	0x10: func(gb *Gameboy) {
		// RL B
		rot := rl(gb, gb.cpu.bc.Hi())
		gb.cpu.bc.SetHi(rot)
	},
	0x11: func(gb *Gameboy) {
		// RL C
		rot := rl(gb, gb.cpu.bc.Lo())
		gb.cpu.bc.SetLo(rot)
	},
	0x12: func(gb *Gameboy) {
		// RL D
		rot := rl(gb, gb.cpu.de.Hi())
		gb.cpu.de.SetHi(rot)
	},
	0x13: func(gb *Gameboy) {
		// RL E
		rot := rl(gb, gb.cpu.de.Lo())
		gb.cpu.de.SetLo(rot)
	},
	0x14: func(gb *Gameboy) {
		// RL H
		rot := rl(gb, gb.cpu.hl.Hi())
		gb.cpu.hl.SetHi(rot)
	},
	0x15: func(gb *Gameboy) {
		// RL L
		rot := rl(gb, gb.cpu.hl.Lo())
		gb.cpu.hl.SetLo(rot)
	},
	0x16: func(gb *Gameboy) {
		// RL (hl)
		addr := gb.cpu.hl.val
		rot := rl(gb, gb.memoryBus.read(addr))
		gb.memoryBus.write(addr, rot)
	},
	0x17: func(gb *Gameboy) {
		// RL A
		rot := rl(gb, gb.cpu.af.Hi())
		gb.cpu.af.SetHi(rot)
	},
	0x18: func(gb *Gameboy) {
		// RR B
		gb.cpu.bc.SetHi(rr(gb, gb.cpu.bc.Hi()))
	},
	0x19: func(gb *Gameboy) {
		// RR C
		gb.cpu.bc.SetLo(rr(gb, gb.cpu.bc.Lo()))
	},
	0x1a: func(gb *Gameboy) {
		// RR D
		gb.cpu.de.SetHi(rr(gb, gb.cpu.de.Hi()))
	},
	0x1b: func(gb *Gameboy) {
		// RR E
		gb.cpu.de.SetLo(rr(gb, gb.cpu.de.Lo()))
	},
	0x1c: func(gb *Gameboy) {
		// RR H
		gb.cpu.hl.SetHi(rr(gb, gb.cpu.hl.Hi()))
	},
	0x1d: func(gb *Gameboy) {
		// RR L
		gb.cpu.hl.SetLo(rr(gb, gb.cpu.hl.Lo()))
	},
	0x1e: func(gb *Gameboy) {
		// RR (hl)
		hl := gb.cpu.hl.val
		gb.memoryBus.write(hl, rr(gb, gb.memoryBus.read(hl)))
	},
	0x1f: func(gb *Gameboy) {
		// RR A
		gb.cpu.af.SetHi(rr(gb, gb.cpu.af.Hi()))
	},
	0x20: func(gb *Gameboy) {
		// SLA B
		rot := sla(gb, gb.cpu.bc.Hi())
		gb.cpu.bc.SetHi(rot)
	},
	0x21: func(gb *Gameboy) {
		// SLA C
		rot := sla(gb, gb.cpu.bc.Lo())
		gb.cpu.bc.SetLo(rot)
	},
	0x22: func(gb *Gameboy) {
		// SLA D
		rot := sla(gb, gb.cpu.de.Hi())
		gb.cpu.de.SetHi(rot)
	},
	0x23: func(gb *Gameboy) {
		// SLA E
		rot := sla(gb, gb.cpu.de.Lo())
		gb.cpu.de.SetLo(rot)
	},
	0x24: func(gb *Gameboy) {
		// SLA H
		rot := sla(gb, gb.cpu.hl.Hi())
		gb.cpu.hl.SetHi(rot)
	},
	0x25: func(gb *Gameboy) {
		// SLA L
		rot := sla(gb, gb.cpu.hl.Lo())
		gb.cpu.hl.SetLo(rot)
	},
	0x26: func(gb *Gameboy) {
		// SLA (hl)
		addr := gb.cpu.hl.val
		rot := sla(gb, gb.memoryBus.read(addr))
		gb.memoryBus.write(addr, rot)
	},
	0x27: func(gb *Gameboy) {
		// SLA A
		rot := sla(gb, gb.cpu.af.Hi())
		gb.cpu.af.SetHi(rot)
	},
	0x28: func(gb *Gameboy) {
		// SRA B
		gb.cpu.bc.SetHi(sra(gb, gb.cpu.bc.Hi()))
	},
	0x29: func(gb *Gameboy) {
		// SRA C
		rot := sra(gb, gb.cpu.bc.Lo())
		gb.cpu.bc.SetLo(rot)
	},
	0x2a: func(gb *Gameboy) {
		// SRA D
		rot := sra(gb, gb.cpu.de.Hi())
		gb.cpu.de.SetHi(rot)
	},
	0x2b: func(gb *Gameboy) {
		// SRA E
		rot := sra(gb, gb.cpu.de.Lo())
		gb.cpu.de.SetLo(rot)
	},
	0x2c: func(gb *Gameboy) {
		// SRA H
		rot := sra(gb, gb.cpu.hl.Hi())
		gb.cpu.hl.SetHi(rot)
	},
	0x2d: func(gb *Gameboy) {
		// SRA L
		rot := sra(gb, gb.cpu.hl.Lo())
		gb.cpu.hl.SetLo(rot)
	},
	0x2e: func(gb *Gameboy) {
		// SRA (hl)
		gb.memoryBus.write(gb.cpu.hl.val, sra(gb, gb.memoryBus.read(gb.cpu.hl.val)))
	},
	0x2f: func(gb *Gameboy) {
		// SRA A
		rot := sra(gb, gb.cpu.af.Hi())
		gb.cpu.af.SetHi(rot)
	},
	0x30: func(gb *Gameboy) {
		// SWAP B
		swap := swap(gb, gb.cpu.bc.Hi())
		gb.cpu.bc.SetHi(swap)
	},
	0x31: func(gb *Gameboy) {
		// SWAP C
		swap := swap(gb, gb.cpu.bc.Lo())
		gb.cpu.bc.SetLo(swap)
	},
	0x32: func(gb *Gameboy) {
		// SWAP D
		swap := swap(gb, gb.cpu.de.Hi())
		gb.cpu.de.SetHi(swap)
	},
	0x33: func(gb *Gameboy) {
		// SWAP E
		swap := swap(gb, gb.cpu.de.Lo())
		gb.cpu.de.SetLo(swap)
	},
	0x34: func(gb *Gameboy) {
		// SWAP H
		swap := swap(gb, gb.cpu.hl.Hi())
		gb.cpu.hl.SetHi(swap)
	},
	0x35: func(gb *Gameboy) {
		// SWAP L
		swap := swap(gb, gb.cpu.hl.Lo())
		gb.cpu.hl.SetLo(swap)
	},
	0x36: func(gb *Gameboy) {
		// SWAP (hl)
		addr := gb.cpu.hl.val
		swap := swap(gb, gb.memoryBus.read(addr))
		gb.memoryBus.write(addr, swap)
	},
	0x37: func(gb *Gameboy) {
		// SWAP A
		swap := swap(gb, gb.cpu.af.Hi())
		gb.cpu.af.SetHi(swap)
	},
	0x38: func(gb *Gameboy) {
		// SRL B
		shift := srl(gb, gb.cpu.bc.Hi())
		gb.cpu.bc.SetHi(shift)
	},
	0x39: func(gb *Gameboy) {
		// SRL C
		shift := srl(gb, gb.cpu.bc.Lo())
		gb.cpu.bc.SetLo(shift)
	},
	0x3a: func(gb *Gameboy) {
		// SRL D
		shift := srl(gb, gb.cpu.de.Hi())
		gb.cpu.de.SetHi(shift)
	},
	0x3b: func(gb *Gameboy) {
		// SRL E
		shift := srl(gb, gb.cpu.de.Lo())
		gb.cpu.de.SetLo(shift)
	},
	0x3c: func(gb *Gameboy) {
		// SRL H
		shift := srl(gb, gb.cpu.hl.Hi())
		gb.cpu.hl.SetHi(shift)
	},
	0x3d: func(gb *Gameboy) {
		// SRL L
		shift := srl(gb, gb.cpu.hl.Lo())
		gb.cpu.hl.SetLo(shift)
	},
	0x3e: func(gb *Gameboy) {
		// SRL (hl)
		addr := gb.cpu.hl.val
		shift := srl(gb, gb.memoryBus.read(addr))
		gb.memoryBus.write(addr, shift)
	},
	0x3f: func(gb *Gameboy) {
		// SRL A
		shift := srl(gb, gb.cpu.af.Hi())
		gb.cpu.af.SetHi(shift)
	},
	0x40: func(gb *Gameboy) {
		// BIT 0, B
		bitInstr(gb, gb.cpu.bc.Hi(), 0)
	},
	0x41: func(gb *Gameboy) {
		// BIT 0, C
		bitInstr(gb, gb.cpu.bc.Lo(), 0)
	},
	0x42: func(gb *Gameboy) {
		// BIT 0, D
		bitInstr(gb, gb.cpu.de.Hi(), 0)
	},
	0x43: func(gb *Gameboy) {
		// BIT 0, E
		bitInstr(gb, gb.cpu.de.Lo(), 0)
	},
	0x44: func(gb *Gameboy) {
		// BIT 0, H
		bitInstr(gb, gb.cpu.hl.Hi(), 0)
	},
	0x45: func(gb *Gameboy) {
		// BIT 0, L
		bitInstr(gb, gb.cpu.hl.Lo(), 0)
	},
	0x46: func(gb *Gameboy) {
		// BIT 0, (hl)
		bitInstr(gb, gb.memoryBus.read(gb.cpu.hl.val), 0)
	},
	0x47: func(gb *Gameboy) {
		// BIT 0, A
		bitInstr(gb, gb.cpu.af.Hi(), 0)
	},
	0x48: func(gb *Gameboy) {
		// BIT 1, B
		bitInstr(gb, gb.cpu.bc.Hi(), 1)
	},
	0x49: func(gb *Gameboy) {
		// BIT 1, C
		bitInstr(gb, gb.cpu.bc.Lo(), 1)
	},
	0x4a: func(gb *Gameboy) {
		// BIT 1, D
		bitInstr(gb, gb.cpu.de.Hi(), 1)
	},
	0x4b: func(gb *Gameboy) {
		// BIT 1, E
		bitInstr(gb, gb.cpu.de.Lo(), 1)
	},
	0x4c: func(gb *Gameboy) {
		// BIT 1, H
		bitInstr(gb, gb.cpu.hl.Hi(), 1)
	},
	0x4d: func(gb *Gameboy) {
		// BIT 1, L
		bitInstr(gb, gb.cpu.hl.Lo(), 1)
	},
	0x4e: func(gb *Gameboy) {
		// BIT 1, (hl)
		bitInstr(gb, gb.memoryBus.read(gb.cpu.hl.val), 1)
	},
	0x4f: func(gb *Gameboy) {
		// BIT 1, A
		bitInstr(gb, gb.cpu.af.Hi(), 1)
	},
	0x50: func(gb *Gameboy) {
		// BIT 2, B
		bitInstr(gb, gb.cpu.bc.Hi(), 2)
	},
	0x51: func(gb *Gameboy) {
		// BIT 2, C
		bitInstr(gb, gb.cpu.bc.Lo(), 2)
	},
	0x52: func(gb *Gameboy) {
		// BIT 2, D
		bitInstr(gb, gb.cpu.de.Hi(), 2)
	},
	0x53: func(gb *Gameboy) {
		// BIT 2, E
		bitInstr(gb, gb.cpu.de.Lo(), 2)
	},
	0x54: func(gb *Gameboy) {
		// BIT 2, H
		bitInstr(gb, gb.cpu.hl.Hi(), 2)
	},
	0x55: func(gb *Gameboy) {
		// BIT 2, L
		bitInstr(gb, gb.cpu.hl.Lo(), 2)
	},
	0x56: func(gb *Gameboy) {
		// BIT 2, (hl)
		bitInstr(gb, gb.memoryBus.read(gb.cpu.hl.val), 2)
	},
	0x57: func(gb *Gameboy) {
		// BIT 2, A
		bitInstr(gb, gb.cpu.af.Hi(), 2)
	},
	0x58: func(gb *Gameboy) {
		// BIT 3, B
		bitInstr(gb, gb.cpu.bc.Hi(), 3)
	},
	0x59: func(gb *Gameboy) {
		// BIT 3, C
		bitInstr(gb, gb.cpu.bc.Lo(), 3)
	},
	0x5a: func(gb *Gameboy) {
		// BIT 3, D
		bitInstr(gb, gb.cpu.de.Hi(), 3)
	},
	0x5b: func(gb *Gameboy) {
		// BIT 3, E
		bitInstr(gb, gb.cpu.de.Lo(), 3)
	},
	0x5c: func(gb *Gameboy) {
		// BIT 3, H
		bitInstr(gb, gb.cpu.hl.Hi(), 3)
	},
	0x5d: func(gb *Gameboy) {
		// BIT 3, L
		bitInstr(gb, gb.cpu.hl.Lo(), 3)
	},
	0x5e: func(gb *Gameboy) {
		// BIT 3, (hl)
		bitInstr(gb, gb.memoryBus.read(gb.cpu.hl.val), 3)
	},
	0x5f: func(gb *Gameboy) {
		// BIT 3, A
		bitInstr(gb, gb.cpu.af.Hi(), 3)
	},
	0x60: func(gb *Gameboy) {
		// BIT 4, B
		bitInstr(gb, gb.cpu.bc.Hi(), 4)
	},
	0x61: func(gb *Gameboy) {
		// BIT 4, C
		bitInstr(gb, gb.cpu.bc.Lo(), 4)
	},
	0x62: func(gb *Gameboy) {
		// BIT 4, D
		bitInstr(gb, gb.cpu.de.Hi(), 4)
	},
	0x63: func(gb *Gameboy) {
		// BIT 4, E
		bitInstr(gb, gb.cpu.de.Lo(), 4)
	},
	0x64: func(gb *Gameboy) {
		// BIT 4, H
		bitInstr(gb, gb.cpu.hl.Hi(), 4)
	},
	0x65: func(gb *Gameboy) {
		// BIT 4, L
		bitInstr(gb, gb.cpu.hl.Lo(), 4)
	},
	0x66: func(gb *Gameboy) {
		// BIT 4, (hl)
		bitInstr(gb, gb.memoryBus.read(gb.cpu.hl.val), 4)
	},
	0x67: func(gb *Gameboy) {
		// BIT 4, A
		bitInstr(gb, gb.cpu.af.Hi(), 4)
	},
	0x68: func(gb *Gameboy) {
		// BIT 5, B
		bitInstr(gb, gb.cpu.bc.Hi(), 5)
	},
	0x69: func(gb *Gameboy) {
		// BIT 5, C
		bitInstr(gb, gb.cpu.bc.Lo(), 5)
	},
	0x6a: func(gb *Gameboy) {
		// BIT 5, D
		bitInstr(gb, gb.cpu.de.Hi(), 5)
	},
	0x6b: func(gb *Gameboy) {
		// BIT 5, E
		bitInstr(gb, gb.cpu.de.Lo(), 5)
	},
	0x6c: func(gb *Gameboy) {
		// BIT 5, H
		bitInstr(gb, gb.cpu.hl.Hi(), 5)
	},
	0x6d: func(gb *Gameboy) {
		// BIT 5, L
		bitInstr(gb, gb.cpu.hl.Lo(), 5)
	},
	0x6e: func(gb *Gameboy) {
		// BIT 5, (hl)
		bitInstr(gb, gb.memoryBus.read(gb.cpu.hl.val), 5)
	},
	0x6f: func(gb *Gameboy) {
		// BIT 5, A
		bitInstr(gb, gb.cpu.af.Hi(), 5)
	},
	0x70: func(gb *Gameboy) {
		// BIT 6, B
		bitInstr(gb, gb.cpu.bc.Hi(), 6)
	},
	0x71: func(gb *Gameboy) {
		// BIT 6, C
		bitInstr(gb, gb.cpu.bc.Lo(), 6)
	},
	0x72: func(gb *Gameboy) {
		// BIT 6, D
		bitInstr(gb, gb.cpu.de.Hi(), 6)
	},
	0x73: func(gb *Gameboy) {
		// BIT 6, E
		bitInstr(gb, gb.cpu.de.Lo(), 6)
	},
	0x74: func(gb *Gameboy) {
		// BIT 6, H
		bitInstr(gb, gb.cpu.hl.Hi(), 6)
	},
	0x75: func(gb *Gameboy) {
		// BIT 6, L
		bitInstr(gb, gb.cpu.hl.Lo(), 6)
	},
	0x76: func(gb *Gameboy) {
		// BIT 6, (hl)
		bitInstr(gb, gb.memoryBus.read(gb.cpu.hl.val), 6)
	},
	0x77: func(gb *Gameboy) {
		// BIT 6, A
		bitInstr(gb, gb.cpu.af.Hi(), 6)
	},
	0x78: func(gb *Gameboy) {
		// BIT 7, B
		bitInstr(gb, gb.cpu.bc.Hi(), 7)
	},
	0x79: func(gb *Gameboy) {
		// BIT 7, C
		bitInstr(gb, gb.cpu.bc.Lo(), 7)
	},
	0x7a: func(gb *Gameboy) {
		// BIT 7, D
		bitInstr(gb, gb.cpu.de.Hi(), 7)
	},
	0x7b: func(gb *Gameboy) {
		// BIT 7, E
		bitInstr(gb, gb.cpu.de.Lo(), 7)
	},
	0x7c: func(gb *Gameboy) {
		// BIT 7, H
		bitInstr(gb, gb.cpu.hl.Hi(), 7)
	},
	0x7d: func(gb *Gameboy) {
		// BIT 7, L
		bitInstr(gb, gb.cpu.hl.Lo(), 7)
	},
	0x7e: func(gb *Gameboy) {
		// BIT 7, (hl)
		bitInstr(gb, gb.memoryBus.read(gb.cpu.hl.val), 7)
	},
	0x7f: func(gb *Gameboy) {
		// BIT 7, A
		bitInstr(gb, gb.cpu.af.Hi(), 7)
	},
	0x80: func(gb *Gameboy) {
		// RES 0, B
		gb.cpu.bc.SetHi(bit.Clear(gb.cpu.bc.Hi(), 0))
	},
	0x81: func(gb *Gameboy) {
		// RES 0, C
		gb.cpu.bc.SetLo(bit.Clear(gb.cpu.bc.Lo(), 0))
	},
	0x82: func(gb *Gameboy) {
		// RES 0, D
		gb.cpu.de.SetHi(bit.Clear(gb.cpu.de.Hi(), 0))
	},
	0x83: func(gb *Gameboy) {
		// RES 0, E
		gb.cpu.de.SetLo(bit.Clear(gb.cpu.de.Lo(), 0))
	},
	0x84: func(gb *Gameboy) {
		// RES 0, H
		gb.cpu.hl.SetHi(bit.Clear(gb.cpu.hl.Hi(), 0))
	},
	0x85: func(gb *Gameboy) {
		// RES 0, L
		gb.cpu.hl.SetLo(bit.Clear(gb.cpu.hl.Lo(), 0))
	},
	0x86: func(gb *Gameboy) {
		// RES 0, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Clear(gb.memoryBus.read(gb.cpu.hl.val), 0))
	},
	0x87: func(gb *Gameboy) {
		// RES 0, A
		gb.cpu.af.SetHi(bit.Clear(gb.cpu.af.Hi(), 0))
	},
	0x88: func(gb *Gameboy) {
		// RES 1, B
		gb.cpu.bc.SetHi(bit.Clear(gb.cpu.bc.Hi(), 1))
	},
	0x89: func(gb *Gameboy) {
		// RES 1, C
		gb.cpu.bc.SetLo(bit.Clear(gb.cpu.bc.Lo(), 1))
	},
	0x8a: func(gb *Gameboy) {
		// RES 1, D
		gb.cpu.de.SetHi(bit.Clear(gb.cpu.de.Hi(), 1))
	},
	0x8b: func(gb *Gameboy) {
		// RES 1, E
		gb.cpu.de.SetLo(bit.Clear(gb.cpu.de.Lo(), 1))
	},
	0x8c: func(gb *Gameboy) {
		// RES 1, H
		gb.cpu.hl.SetHi(bit.Clear(gb.cpu.hl.Hi(), 1))
	},
	0x8d: func(gb *Gameboy) {
		// RES 1, L
		gb.cpu.hl.SetLo(bit.Clear(gb.cpu.hl.Lo(), 1))
	},
	0x8e: func(gb *Gameboy) {
		// RES 1, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Clear(gb.memoryBus.read(gb.cpu.hl.val), 1))
	},
	0x8f: func(gb *Gameboy) {
		// RES 1, A
		gb.cpu.af.SetHi(bit.Clear(gb.cpu.af.Hi(), 1))
	},
	0x90: func(gb *Gameboy) {
		// RES 2, B
		gb.cpu.bc.SetHi(bit.Clear(gb.cpu.bc.Hi(), 2))
	},
	0x91: func(gb *Gameboy) {
		// RES 2, C
		gb.cpu.bc.SetLo(bit.Clear(gb.cpu.bc.Lo(), 2))
	},
	0x92: func(gb *Gameboy) {
		// RES 2, D
		gb.cpu.de.SetHi(bit.Clear(gb.cpu.de.Hi(), 2))
	},
	0x93: func(gb *Gameboy) {
		// RES 2, E
		gb.cpu.de.SetLo(bit.Clear(gb.cpu.de.Lo(), 2))
	},
	0x94: func(gb *Gameboy) {
		// RES 2, H
		gb.cpu.hl.SetHi(bit.Clear(gb.cpu.hl.Hi(), 2))
	},
	0x95: func(gb *Gameboy) {
		// RES 2, L
		gb.cpu.hl.SetLo(bit.Clear(gb.cpu.hl.Lo(), 2))
	},
	0x96: func(gb *Gameboy) {
		// RES 2, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Clear(gb.memoryBus.read(gb.cpu.hl.val), 2))
	},
	0x97: func(gb *Gameboy) {
		// RES 2, A
		gb.cpu.af.SetHi(bit.Clear(gb.cpu.af.Hi(), 2))
	},
	0x98: func(gb *Gameboy) {
		// RES 3, B
		gb.cpu.bc.SetHi(bit.Clear(gb.cpu.bc.Hi(), 3))
	},
	0x99: func(gb *Gameboy) {
		// RES 3, C
		gb.cpu.bc.SetLo(bit.Clear(gb.cpu.bc.Lo(), 3))
	},
	0x9a: func(gb *Gameboy) {
		// RES 3, D
		gb.cpu.de.SetHi(bit.Clear(gb.cpu.de.Hi(), 3))
	},
	0x9b: func(gb *Gameboy) {
		// RES 3, E
		gb.cpu.de.SetLo(bit.Clear(gb.cpu.de.Lo(), 3))
	},
	0x9c: func(gb *Gameboy) {
		// RES 3, H
		gb.cpu.hl.SetHi(bit.Clear(gb.cpu.hl.Hi(), 3))
	},
	0x9d: func(gb *Gameboy) {
		// RES 3, L
		gb.cpu.hl.SetLo(bit.Clear(gb.cpu.hl.Lo(), 3))
	},
	0x9e: func(gb *Gameboy) {
		// RES 3, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Clear(gb.memoryBus.read(gb.cpu.hl.val), 3))
	},
	0x9f: func(gb *Gameboy) {
		// RES 3, A
		gb.cpu.af.SetHi(bit.Clear(gb.cpu.af.Hi(), 3))
	},

	0xa0: func(gb *Gameboy) {
		// RES 4, B
		gb.cpu.bc.SetHi(bit.Clear(gb.cpu.bc.Hi(), 4))
	},
	0xa1: func(gb *Gameboy) {
		// RES 4, C
		gb.cpu.bc.SetLo(bit.Clear(gb.cpu.bc.Lo(), 4))
	},
	0xa2: func(gb *Gameboy) {
		// RES 4, D
		gb.cpu.de.SetHi(bit.Clear(gb.cpu.de.Hi(), 4))
	},
	0xa3: func(gb *Gameboy) {
		// RES 4, E
		gb.cpu.de.SetLo(bit.Clear(gb.cpu.de.Lo(), 4))
	},
	0xa4: func(gb *Gameboy) {
		// RES 4, H
		gb.cpu.hl.SetHi(bit.Clear(gb.cpu.hl.Hi(), 4))
	},
	0xa5: func(gb *Gameboy) {
		// RES 4, L
		gb.cpu.hl.SetLo(bit.Clear(gb.cpu.hl.Lo(), 4))
	},
	0xa6: func(gb *Gameboy) {
		// RES 4, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Clear(gb.memoryBus.read(gb.cpu.hl.val), 4))
	},
	0xa7: func(gb *Gameboy) {
		// RES 4, A
		gb.cpu.af.SetHi(bit.Clear(gb.cpu.af.Hi(), 4))
	},
	0xa8: func(gb *Gameboy) {
		// RES 5, B
		gb.cpu.bc.SetHi(bit.Clear(gb.cpu.bc.Hi(), 5))
	},
	0xa9: func(gb *Gameboy) {
		// RES 5, C
		gb.cpu.bc.SetLo(bit.Clear(gb.cpu.bc.Lo(), 5))
	},
	0xaa: func(gb *Gameboy) {
		// RES 5, D
		gb.cpu.de.SetHi(bit.Clear(gb.cpu.de.Hi(), 5))
	},
	0xab: func(gb *Gameboy) {
		// RES 5, E
		gb.cpu.de.SetLo(bit.Clear(gb.cpu.de.Lo(), 5))
	},
	0xac: func(gb *Gameboy) {
		// RES 5, H
		gb.cpu.hl.SetHi(bit.Clear(gb.cpu.hl.Hi(), 5))
	},
	0xad: func(gb *Gameboy) {
		// RES 5, L
		gb.cpu.hl.SetLo(bit.Clear(gb.cpu.hl.Lo(), 5))
	},
	0xae: func(gb *Gameboy) {
		// RES 5, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Clear(gb.memoryBus.read(gb.cpu.hl.val), 5))
	},
	0xaf: func(gb *Gameboy) {
		// RES 5, A
		gb.cpu.af.SetHi(bit.Clear(gb.cpu.af.Hi(), 5))
	},
	0xb0: func(gb *Gameboy) {
		// RES 6, B
		gb.cpu.bc.SetHi(bit.Clear(gb.cpu.bc.Hi(), 6))
	},
	0xb1: func(gb *Gameboy) {
		// RES 6, C
		gb.cpu.bc.SetLo(bit.Clear(gb.cpu.bc.Lo(), 6))
	},
	0xb2: func(gb *Gameboy) {
		// RES 6, D
		gb.cpu.de.SetHi(bit.Clear(gb.cpu.de.Hi(), 6))
	},
	0xb3: func(gb *Gameboy) {
		// RES 6, E
		gb.cpu.de.SetLo(bit.Clear(gb.cpu.de.Lo(), 6))
	},
	0xb4: func(gb *Gameboy) {
		// RES 6, H
		gb.cpu.hl.SetHi(bit.Clear(gb.cpu.hl.Hi(), 6))
	},
	0xb5: func(gb *Gameboy) {
		// RES 6, L
		gb.cpu.hl.SetLo(bit.Clear(gb.cpu.hl.Lo(), 6))
	},
	0xb6: func(gb *Gameboy) {
		// RES 6, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Clear(gb.memoryBus.read(gb.cpu.hl.val), 6))
	},
	0xb7: func(gb *Gameboy) {
		// RES 6, A
		gb.cpu.af.SetHi(bit.Clear(gb.cpu.af.Hi(), 6))
	},
	0xb8: func(gb *Gameboy) {
		// RES 7, B
		gb.cpu.bc.SetHi(bit.Clear(gb.cpu.bc.Hi(), 7))
	},
	0xb9: func(gb *Gameboy) {
		// RES 7, C
		gb.cpu.bc.SetLo(bit.Clear(gb.cpu.bc.Lo(), 7))
	},
	0xba: func(gb *Gameboy) {
		// RES 7, D
		gb.cpu.de.SetHi(bit.Clear(gb.cpu.de.Hi(), 7))
	},
	0xbb: func(gb *Gameboy) {
		// RES 7, E
		gb.cpu.de.SetLo(bit.Clear(gb.cpu.de.Lo(), 7))
	},
	0xbc: func(gb *Gameboy) {
		// RES 7, H
		gb.cpu.hl.SetHi(bit.Clear(gb.cpu.hl.Hi(), 7))
	},
	0xbd: func(gb *Gameboy) {
		// RES 7, L
		gb.cpu.hl.SetLo(bit.Clear(gb.cpu.hl.Lo(), 7))
	},
	0xbe: func(gb *Gameboy) {
		// RES 7, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Clear(gb.memoryBus.read(gb.cpu.hl.val), 7))
	},
	0xbf: func(gb *Gameboy) {
		// RES 7, A
		gb.cpu.af.SetHi(bit.Clear(gb.cpu.af.Hi(), 7))
	},

	0xc0: func(gb *Gameboy) {
		// SET 0, B
		gb.cpu.bc.SetHi(bit.Set(gb.cpu.bc.Hi(), 0))
	},
	0xc1: func(gb *Gameboy) {
		// SET 0, C
		gb.cpu.bc.SetLo(bit.Set(gb.cpu.bc.Lo(), 0))
	},
	0xc2: func(gb *Gameboy) {
		// SET 0, D
		gb.cpu.de.SetHi(bit.Set(gb.cpu.de.Hi(), 0))
	},
	0xc3: func(gb *Gameboy) {
		// SET 0, E
		gb.cpu.de.SetLo(bit.Set(gb.cpu.de.Lo(), 0))
	},
	0xc4: func(gb *Gameboy) {
		// SET 0, H
		gb.cpu.hl.SetHi(bit.Set(gb.cpu.hl.Hi(), 0))
	},
	0xc5: func(gb *Gameboy) {
		// SET 0, L
		gb.cpu.hl.SetLo(bit.Set(gb.cpu.hl.Lo(), 0))
	},
	0xc6: func(gb *Gameboy) {
		// SET 0, (hl)
		gb.memoryBus.write(gb.cpu.hl.val, bit.Set(gb.memoryBus.read(gb.cpu.hl.val), 0))
	},
	0xc7: func(gb *Gameboy) {
		// SET 0, A
		gb.cpu.af.SetHi(bit.Set(gb.cpu.af.Hi(), 0))
	},
	0xc8: func(gb *Gameboy) {
		// SET 1, B
		gb.cpu.bc.SetHi(bit.Set(gb.cpu.bc.Hi(), 1))
	},
	0xc9: func(gb *Gameboy) {
		// SET 1, C
		gb.cpu.bc.SetLo(bit.Set(gb.cpu.bc.Lo(), 1))
	},
	0xca: func(gb *Gameboy) {
		// SET 1, D
		gb.cpu.de.SetHi(bit.Set(gb.cpu.de.Hi(), 1))
	},
	0xcb: func(gb *Gameboy) {
		// SET 1, E
		gb.cpu.de.SetLo(bit.Set(gb.cpu.de.Lo(), 1))
	},
	0xcc: func(gb *Gameboy) {
		// SET 1, H
		gb.cpu.hl.SetHi(bit.Set(gb.cpu.hl.Hi(), 1))
	},
	0xcd: func(gb *Gameboy) {
		// SET 1, L
		gb.cpu.hl.SetLo(bit.Set(gb.cpu.hl.Lo(), 1))
	},
	0xce: func(gb *Gameboy) {
		// SET 1, hl
		gb.memoryBus.write(gb.cpu.hl.val, bit.Set(gb.memoryBus.read(gb.cpu.hl.val), 1))
	},
	0xcf: func(gb *Gameboy) {
		// SET 1, A
		gb.cpu.af.SetHi(bit.Set(gb.cpu.af.Hi(), 1))
	},
	0xd0: func(gb *Gameboy) {
		// SET 2, B
		gb.cpu.bc.SetHi(bit.Set(gb.cpu.bc.Hi(), 2))
	},
	0xd1: func(gb *Gameboy) {
		// SET 2, C
		gb.cpu.bc.SetLo(bit.Set(gb.cpu.bc.Lo(), 2))
	},
	0xd2: func(gb *Gameboy) {
		// SET 2, D
		gb.cpu.de.SetHi(bit.Set(gb.cpu.de.Hi(), 2))
	},
	0xd3: func(gb *Gameboy) {
		// SET 2, E
		gb.cpu.de.SetLo(bit.Set(gb.cpu.de.Lo(), 2))
	},
	0xd4: func(gb *Gameboy) {
		// SET 2, H
		gb.cpu.hl.SetHi(bit.Set(gb.cpu.hl.Hi(), 2))
	},
	0xd5: func(gb *Gameboy) {
		// SET 2, L
		gb.cpu.hl.SetLo(bit.Set(gb.cpu.hl.Lo(), 2))
	},
	0xd6: func(gb *Gameboy) {
		// SET 02 hl
		gb.memoryBus.write(gb.cpu.hl.val, bit.Set(gb.memoryBus.read(gb.cpu.hl.val), 2))
	},
	0xd7: func(gb *Gameboy) {
		// SET 2, A
		gb.cpu.af.SetHi(bit.Set(gb.cpu.af.Hi(), 2))
	},
	0xd8: func(gb *Gameboy) {
		// SET 3, B
		gb.cpu.bc.SetHi(bit.Set(gb.cpu.bc.Hi(), 3))
	},
	0xd9: func(gb *Gameboy) {
		// SET 3, C
		gb.cpu.bc.SetLo(bit.Set(gb.cpu.bc.Lo(), 3))
	},
	0xda: func(gb *Gameboy) {
		// SET 3, D
		gb.cpu.de.SetHi(bit.Set(gb.cpu.de.Hi(), 3))
	},
	0xdb: func(gb *Gameboy) {
		// SET 3, E
		gb.cpu.de.SetLo(bit.Set(gb.cpu.de.Lo(), 3))
	},
	0xdc: func(gb *Gameboy) {
		// SET 3, H
		gb.cpu.hl.SetHi(bit.Set(gb.cpu.hl.Hi(), 3))
	},
	0xdd: func(gb *Gameboy) {
		// SET 3, L
		gb.cpu.hl.SetLo(bit.Set(gb.cpu.hl.Lo(), 3))
	},
	0xde: func(gb *Gameboy) {
		// SET 3, hl
		gb.memoryBus.write(gb.cpu.hl.val, bit.Set(gb.memoryBus.read(gb.cpu.hl.val), 3))
	},
	0xdf: func(gb *Gameboy) {
		// SET 3, A
		gb.cpu.af.SetHi(bit.Set(gb.cpu.af.Hi(), 3))
	},
	0xe0: func(gb *Gameboy) {
		// SET 4, B
		gb.cpu.bc.SetHi(bit.Set(gb.cpu.bc.Hi(), 4))
	},
	0xe1: func(gb *Gameboy) {
		// SET 4, C
		gb.cpu.bc.SetLo(bit.Set(gb.cpu.bc.Lo(), 4))
	},
	0xe2: func(gb *Gameboy) {
		// SET 4, D
		gb.cpu.de.SetHi(bit.Set(gb.cpu.de.Hi(), 4))
	},
	0xe3: func(gb *Gameboy) {
		// SET 4, E
		gb.cpu.de.SetLo(bit.Set(gb.cpu.de.Lo(), 4))
	},
	0xe4: func(gb *Gameboy) {
		// SET 4, H
		gb.cpu.hl.SetHi(bit.Set(gb.cpu.hl.Hi(), 4))
	},
	0xe5: func(gb *Gameboy) {
		// SET 4, L
		gb.cpu.hl.SetLo(bit.Set(gb.cpu.hl.Lo(), 4))
	},
	0xe6: func(gb *Gameboy) {
		// SET 04 hl
		gb.memoryBus.write(gb.cpu.hl.val, bit.Set(gb.memoryBus.read(gb.cpu.hl.val), 4))
	},
	0xe7: func(gb *Gameboy) {
		// SET 4, A
		gb.cpu.af.SetHi(bit.Set(gb.cpu.af.Hi(), 4))
	},
	0xe8: func(gb *Gameboy) {
		// SET 5, B
		gb.cpu.bc.SetHi(bit.Set(gb.cpu.bc.Hi(), 5))
	},
	0xe9: func(gb *Gameboy) {
		// SET 5, C
		gb.cpu.bc.SetLo(bit.Set(gb.cpu.bc.Lo(), 5))
	},
	0xea: func(gb *Gameboy) {
		// SET 5, D
		gb.cpu.de.SetHi(bit.Set(gb.cpu.de.Hi(), 5))
	},
	0xeb: func(gb *Gameboy) {
		// SET 5, E
		gb.cpu.de.SetLo(bit.Set(gb.cpu.de.Lo(), 5))
	},
	0xec: func(gb *Gameboy) {
		// SET 5, H
		gb.cpu.hl.SetHi(bit.Set(gb.cpu.hl.Hi(), 5))
	},
	0xed: func(gb *Gameboy) {
		// SET 5, L
		gb.cpu.hl.SetLo(bit.Set(gb.cpu.hl.Lo(), 5))
	},
	0xee: func(gb *Gameboy) {
		// SET 5, hl
		gb.memoryBus.write(gb.cpu.hl.val, bit.Set(gb.memoryBus.read(gb.cpu.hl.val), 5))
	},
	0xef: func(gb *Gameboy) {
		// SET 5, A
		gb.cpu.af.SetHi(bit.Set(gb.cpu.af.Hi(), 5))
	},
	0xf0: func(gb *Gameboy) {
		// SET 6, B
		gb.cpu.bc.SetHi(bit.Set(gb.cpu.bc.Hi(), 6))
	},
	0xf1: func(gb *Gameboy) {
		// SET 6, C
		gb.cpu.bc.SetLo(bit.Set(gb.cpu.bc.Lo(), 6))
	},
	0xf2: func(gb *Gameboy) {
		// SET 6, D
		gb.cpu.de.SetHi(bit.Set(gb.cpu.de.Hi(), 6))
	},
	0xf3: func(gb *Gameboy) {
		// SET 6, E
		gb.cpu.de.SetLo(bit.Set(gb.cpu.de.Lo(), 6))
	},
	0xf4: func(gb *Gameboy) {
		// SET 6, H
		gb.cpu.hl.SetHi(bit.Set(gb.cpu.hl.Hi(), 6))
	},
	0xf5: func(gb *Gameboy) {
		// SET 6, L
		gb.cpu.hl.SetLo(bit.Set(gb.cpu.hl.Lo(), 6))
	},
	0xf6: func(gb *Gameboy) {
		// SET 6, hl
		gb.memoryBus.write(gb.cpu.hl.val, bit.Set(gb.memoryBus.read(gb.cpu.hl.val), 6))
	},
	0xf7: func(gb *Gameboy) {
		// SET 6, A
		gb.cpu.af.SetHi(bit.Set(gb.cpu.af.Hi(), 6))
	},
	0xf8: func(gb *Gameboy) {
		// SET 7, B
		gb.cpu.bc.SetHi(bit.Set(gb.cpu.bc.Hi(), 7))
	},
	0xf9: func(gb *Gameboy) {
		// SET 7, C
		gb.cpu.bc.SetLo(bit.Set(gb.cpu.bc.Lo(), 7))
	},
	0xfa: func(gb *Gameboy) {
		// SET 7, D
		gb.cpu.de.SetHi(bit.Set(gb.cpu.de.Hi(), 7))
	},
	0xfb: func(gb *Gameboy) {
		// SET 7, E
		gb.cpu.de.SetLo(bit.Set(gb.cpu.de.Lo(), 7))
	},
	0xfc: func(gb *Gameboy) {
		// SET 7, H
		gb.cpu.hl.SetHi(bit.Set(gb.cpu.hl.Hi(), 7))
	},
	0xfd: func(gb *Gameboy) {
		// SET 7, L
		gb.cpu.hl.SetLo(bit.Set(gb.cpu.hl.Lo(), 7))
	},
	0xfe: func(gb *Gameboy) {
		// SET 7, hl
		gb.memoryBus.write(gb.cpu.hl.val, bit.Set(gb.memoryBus.read(gb.cpu.hl.val), 7))
	},
	0xff: func(gb *Gameboy) {
		// SET 7, A
		gb.cpu.af.SetHi(bit.Set(gb.cpu.af.Hi(), 7))
	},
}
