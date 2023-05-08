package main

var cbInstructionCycles = []int{
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2,
}

var prefixedInstructions = [0x100]func(){
	0x00: func() {
		// RLC B
		gb.cpu.rlc8Reg(REG_B)
	},
	0x01: func() {
		// RLC C
		gb.cpu.rlc8Reg(REG_C)
	},
	0x02: func() {
		// RLC D
		gb.cpu.rlc8Reg(REG_D)
	},
	0x03: func() {
		// RLC E
		gb.cpu.rlc8Reg(REG_E)
	},
	0x04: func() {
		// RLC H
		gb.cpu.rlc8Reg(REG_H)
	},
	0x05: func() {
		// RLC L
		gb.cpu.rlc8Reg(REG_L)
	},
	0x06: func() {
		// RLC (HL)
		gb.cpu.rlcHL()
	},
	0x07: func() {
		// RLC A
		gb.cpu.rlc8Reg(REG_A)
	},
	0x08: func() {
		// RRC B
		gb.cpu.rrc8Reg(REG_B)
	},
	0x09: func() {
		// RRC C
		gb.cpu.rrc8Reg(REG_C)
	},
	0x0a: func() {
		// RRC D
		gb.cpu.rrc8Reg(REG_D)
	},
	0x0b: func() {
		// RRC E
		gb.cpu.rrc8Reg(REG_E)
	},
	0x0c: func() {
		// RRC H
		gb.cpu.rrc8Reg(REG_H)
	},
	0x0d: func() {
		// RRC L
		gb.cpu.rrc8Reg(REG_L)
	},
	0x0e: func() {
		// RRC (HL)
		gb.cpu.rrcHL()
	},
	0x0f: func() {
		// RRC A
		gb.cpu.rrc8Reg(REG_A)
	},

	0x10: func() {
		// RL B
		gb.cpu.rl8Reg(REG_B)
	},
	0x11: func() {
		// RL C
		gb.cpu.rl8Reg(REG_C)
	},
	0x12: func() {
		// RL D
		gb.cpu.rl8Reg(REG_D)
	},
	0x13: func() {
		// RL E
		gb.cpu.rl8Reg(REG_E)
	},
	0x14: func() {
		// RL H
		gb.cpu.rl8Reg(REG_H)
	},
	0x15: func() {
		// RL L
		gb.cpu.rl8Reg(REG_L)
	},
	0x16: func() {
		// RL (HL)
		gb.cpu.rlHL()
	},
	0x17: func() {
		// RL A
		gb.cpu.rl8Reg(REG_A)
	},
	0x18: func() {
		// RR B
		gb.cpu.rr8Reg(REG_B)
	},
	0x19: func() {
		// RR C
		gb.cpu.rr8Reg(REG_C)
	},
	0x1a: func() {
		// RR D
		gb.cpu.rr8Reg(REG_D)
	},
	0x1b: func() {
		// RR E
		gb.cpu.rr8Reg(REG_E)
	},
	0x1c: func() {
		// RR H
		gb.cpu.rr8Reg(REG_H)
	},
	0x1d: func() {
		// RR L
		gb.cpu.rr8Reg(REG_L)
	},
	0x1e: func() {
		// RR (HL)
		gb.cpu.rrHL()
	},
	0x1f: func() {
		// RR A
		gb.cpu.rr8Reg(REG_A)
	},
	0x20: func() {
		// SLA B
		gb.cpu.sla8Reg(REG_B)
	},
	0x21: func() {
		// SLA C
		gb.cpu.sla8Reg(REG_C)
	},
	0x22: func() {
		// SLA D
		gb.cpu.sla8Reg(REG_D)
	},
	0x23: func() {
		// SLA E
		gb.cpu.sla8Reg(REG_E)
	},
	0x24: func() {
		// SLA H
		gb.cpu.sla8Reg(REG_H)
	},
	0x25: func() {
		// SLA L
		gb.cpu.sla8Reg(REG_L)
	},
	0x26: func() {
		// SLA (HL)
		gb.cpu.slaHL()
	},
	0x27: func() {
		// SLA A
		gb.cpu.sla8Reg(REG_A)
	},
	0x28: func() {
		// SRA B
		gb.cpu.sra8Reg(REG_B)
	},
	0x29: func() {
		// SRA C
		gb.cpu.sra8Reg(REG_C)
	},
	0x2a: func() {
		// SRA D
		gb.cpu.sra8Reg(REG_D)
	},
	0x2b: func() {
		// SRA E
		gb.cpu.sra8Reg(REG_E)
	},
	0x2c: func() {
		// SRA H
		gb.cpu.sra8Reg(REG_H)
	},
	0x2d: func() {
		// SRA L
		gb.cpu.sra8Reg(REG_L)
	},
	0x2e: func() {
		// SRA (HL)
		gb.cpu.sraHL()
	},
	0x2f: func() {
		// SRA A
		gb.cpu.sra8Reg(REG_A)
	},

	0x30: func() {
		// SWAP B
		gb.cpu.swap8Reg(REG_B)
	},
	0x31: func() {
		// SWAP C
		gb.cpu.swap8Reg(REG_C)
	},
	0x32: func() {
		// SWAP D
		gb.cpu.swap8Reg(REG_D)
	},
	0x33: func() {
		// SWAP E
		gb.cpu.swap8Reg(REG_E)
	},
	0x34: func() {
		// SWAP H
		gb.cpu.swap8Reg(REG_H)
	},
	0x35: func() {
		// SWAP L
		gb.cpu.swap8Reg(REG_L)
	},
	0x36: func() {
		// SWAP (HL)
		gb.cpu.swapHL()
	},
	0x37: func() {
		// SWAP A
		gb.cpu.swap8Reg(REG_A)
	},
	0x38: func() {
		// SRL B
		gb.cpu.srl8Reg(REG_B)
	},
	0x39: func() {
		// SRL C
		gb.cpu.srl8Reg(REG_C)
	},
	0x3a: func() {
		// SRL D
		gb.cpu.srl8Reg(REG_D)
	},
	0x3b: func() {
		// SRL E
		gb.cpu.srl8Reg(REG_E)
	},
	0x3c: func() {
		// SRL H
		gb.cpu.srl8Reg(REG_H)
	},
	0x3d: func() {
		// SRL L
		gb.cpu.srl8Reg(REG_L)
	},
	0x3e: func() {
		// SRL (HL)
		gb.cpu.srlHL()
	},
	0x3f: func() {
		// SRL A
		gb.cpu.srl8Reg(REG_A)
	},

	0x40: func() {
		// BIT 0, B
		gb.cpu.bit8Reg(REG_B, 0)
	},
	0x41: func() {
		// BIT 0, C
		gb.cpu.bit8Reg(REG_C, 0)
	},
	0x42: func() {
		// BIT 0, D
		gb.cpu.bit8Reg(REG_D, 0)
	},
	0x43: func() {
		// BIT 0, E
		gb.cpu.bit8Reg(REG_E, 0)
	},
	0x44: func() {
		// BIT 0, H
		gb.cpu.bit8Reg(REG_H, 0)
	},
	0x45: func() {
		// BIT 0, L
		gb.cpu.bit8Reg(REG_L, 0)
	},
	0x46: func() {
		// BIT 0, (HL)
		gb.cpu.bitHL(0)
	},
	0x47: func() {
		// BIT 0, A
		gb.cpu.bit8Reg(REG_A, 0)
	},
	0x48: func() {
		// BIT 1, B
		gb.cpu.bit8Reg(REG_B, 1)
	},
	0x49: func() {
		// BIT 1, C
		gb.cpu.bit8Reg(REG_C, 1)
	},
	0x4a: func() {
		// BIT 1, D
		gb.cpu.bit8Reg(REG_D, 1)
	},
	0x4b: func() {
		// BIT 1, E
		gb.cpu.bit8Reg(REG_E, 1)
	},
	0x4c: func() {
		// BIT 1, H
		gb.cpu.bit8Reg(REG_H, 1)
	},
	0x4d: func() {
		// BIT 1, L
		gb.cpu.bit8Reg(REG_L, 1)

	},
	0x4e: func() {
		// BIT 1, (HL)
		gb.cpu.bitHL(1)
	},
	0x4f: func() {
		// BIT 1, A
		gb.cpu.bit8Reg(REG_A, 1)
	},
	0x50: func() {
		// BIT 2, B
		gb.cpu.bit8Reg(REG_B, 2)
	},
	0x51: func() {
		// BIT 2, C
		gb.cpu.bit8Reg(REG_C, 2)

	},
	0x52: func() {
		// BIT 2, D
		gb.cpu.bit8Reg(REG_D, 2)

	},
	0x53: func() {
		// BIT 2, E
		gb.cpu.bit8Reg(REG_E, 2)
	},
	0x54: func() {
		// BIT 2, H
		gb.cpu.bit8Reg(REG_H, 2)
	},
	0x55: func() {
		// BIT 2, L
		gb.cpu.bit8Reg(REG_L, 2)
	},
	0x56: func() {
		// BIT 2, (HL)
		gb.cpu.bitHL(2)
	},
	0x57: func() {
		// BIT 2, A
		gb.cpu.bit8Reg(REG_A, 2)
	},
	0x58: func() {
		// BIT 3, B
		gb.cpu.bit8Reg(REG_B, 3)
	},
	0x59: func() {
		// BIT 3, C
		gb.cpu.bit8Reg(REG_C, 3)
	},
	0x5a: func() {
		// BIT 3, D
		gb.cpu.bit8Reg(REG_D, 3)
	},
	0x5b: func() {
		// BIT 3, E
		gb.cpu.bit8Reg(REG_E, 3)
	},
	0x5c: func() {
		// BIT 3, H
		gb.cpu.bit8Reg(REG_H, 3)
	},
	0x5d: func() {
		// BIT 3, L
		gb.cpu.bit8Reg(REG_L, 3)
	},
	0x5e: func() {
		// BIT 3, (HL)
		gb.cpu.bitHL(3)
	},
	0x5f: func() {
		// BIT 3, A
		gb.cpu.bit8Reg(REG_A, 3)
	},
	0x60: func() {
		// BIT 4, B
		gb.cpu.bit8Reg(REG_B, 4)
	},
	0x61: func() {
		// BIT 4, C
		gb.cpu.bit8Reg(REG_C, 4)
	},
	0x62: func() {
		// BIT 4, D
		gb.cpu.bit8Reg(REG_D, 4)
	},
	0x63: func() {
		// BIT 4, E
		gb.cpu.bit8Reg(REG_E, 4)
	},
	0x64: func() {
		// BIT 4, H
		gb.cpu.bit8Reg(REG_H, 4)
	},
	0x65: func() {
		// BIT 4, L
		gb.cpu.bit8Reg(REG_L, 4)
	},
	0x66: func() {
		// BIT 4, (HL)
		gb.cpu.bitHL(4)
	},
	0x67: func() {
		// BIT 4, A
		gb.cpu.bit8Reg(REG_A, 4)
	},
	0x68: func() {
		// BIT 5, B
		gb.cpu.bit8Reg(REG_B, 5)
	},
	0x69: func() {
		// BIT 5, C
		gb.cpu.bit8Reg(REG_C, 5)
	},
	0x6a: func() {
		// BIT 5, D
		gb.cpu.bit8Reg(REG_D, 5)
	},
	0x6b: func() {
		// BIT 5, E
		gb.cpu.bit8Reg(REG_E, 5)
	},
	0x6c: func() {
		// BIT 5, H
		gb.cpu.bit8Reg(REG_H, 5)
	},
	0x6d: func() {
		// BIT 5, L
		gb.cpu.bit8Reg(REG_L, 5)
	},
	0x6e: func() {
		// BIT 5, (HL)
		gb.cpu.bitHL(5)
	},
	0x6f: func() {
		// BIT 5, A
		gb.cpu.bit8Reg(REG_A, 5)
	},
	0x70: func() {
		// BIT 6, B
		gb.cpu.bit8Reg(REG_B, 6)

	},
	0x71: func() {
		// BIT 6, C
		gb.cpu.bit8Reg(REG_C, 6)

	},
	0x72: func() {
		// BIT 6, D
		gb.cpu.bit8Reg(REG_D, 6)

	},
	0x73: func() {
		// BIT 6, E
		gb.cpu.bit8Reg(REG_E, 6)

	},
	0x74: func() {
		// BIT 6, H
		gb.cpu.bit8Reg(REG_H, 6)

	},
	0x75: func() {
		// BIT 6, L
		gb.cpu.bit8Reg(REG_L, 6)

	},
	0x76: func() {
		// BIT 6, (HL)
		gb.cpu.bitHL(6)
	},
	0x77: func() {
		// BIT 6, A
		gb.cpu.bit8Reg(REG_A, 6)

	},
	0x78: func() {
		// BIT 7, B
		gb.cpu.bit8Reg(REG_B, 7)

	},
	0x79: func() {
		// BIT 7, C
		gb.cpu.bit8Reg(REG_C, 7)

	},
	0x7a: func() {
		// BIT 7, D
		gb.cpu.bit8Reg(REG_D, 7)

	},
	0x7b: func() {
		// BIT 7, E
		gb.cpu.bit8Reg(REG_E, 7)

	},
	0x7c: func() {
		// BIT 7, H
		gb.cpu.bit8Reg(REG_H, 7)

	},
	0x7d: func() {
		// BIT 7, L
		gb.cpu.bit8Reg(REG_L, 7)

	},
	0x7e: func() {
		// BIT 7, (HL)
		gb.cpu.bitHL(7)
	},
	0x7f: func() {
		// BIT 7, A
		gb.cpu.bit8Reg(REG_A, 7)

	},

	0x80: func() {
		// RES 0, B
		gb.cpu.set8Reg(REG_B, clearBit(gb.cpu.read8Reg(REG_B), 0))
	},
	0x81: func() {
		// RES 0, C
		gb.cpu.set8Reg(REG_C, clearBit(gb.cpu.read8Reg(REG_C), 0))
	},
	0x82: func() {
		// RES 0, D
		gb.cpu.set8Reg(REG_D, clearBit(gb.cpu.read8Reg(REG_D), 0))
	},
	0x83: func() {
		// RES 0, E
		gb.cpu.set8Reg(REG_E, clearBit(gb.cpu.read8Reg(REG_E), 0))
	},
	0x84: func() {
		// RES 0, H
		gb.cpu.set8Reg(REG_H, clearBit(gb.cpu.read8Reg(REG_H), 0))
	},
	0x85: func() {
		// RES 0, L
		gb.cpu.set8Reg(REG_L, clearBit(gb.cpu.read8Reg(REG_L), 0))
	},
	0x86: func() {
		// RES 0, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, clearBit(val, 0))
	},
	0x87: func() {
		// RES 0, A
		gb.cpu.set8Reg(REG_A, clearBit(gb.cpu.read8Reg(REG_A), 0))
	},
	0x88: func() {
		// RES 1, B
		gb.cpu.set8Reg(REG_B, clearBit(gb.cpu.read8Reg(REG_B), 1))
	},
	0x89: func() {
		// RES 1, C
		gb.cpu.set8Reg(REG_C, clearBit(gb.cpu.read8Reg(REG_C), 1))
	},
	0x8a: func() {
		// RES 1, D
		gb.cpu.set8Reg(REG_D, clearBit(gb.cpu.read8Reg(REG_D), 1))
	},
	0x8b: func() {
		// RES 1, E
		gb.cpu.set8Reg(REG_E, clearBit(gb.cpu.read8Reg(REG_E), 1))
	},
	0x8c: func() {
		// RES 1, H
		gb.cpu.set8Reg(REG_H, clearBit(gb.cpu.read8Reg(REG_H), 1))
	},
	0x8d: func() {
		// RES 1, L
		gb.cpu.set8Reg(REG_L, clearBit(gb.cpu.read8Reg(REG_L), 1))
	},
	0x8e: func() {
		// RES 1, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, clearBit(val, 1))
	},
	0x8f: func() {
		// RES 1, A
		gb.cpu.set8Reg(REG_A, clearBit(gb.cpu.read8Reg(REG_A), 1))
	},
	0x90: func() {
		// RES 2, B
		gb.cpu.set8Reg(REG_B, clearBit(gb.cpu.read8Reg(REG_B), 2))
	},
	0x91: func() {
		// RES 2, C
		gb.cpu.set8Reg(REG_C, clearBit(gb.cpu.read8Reg(REG_C), 2))
	},
	0x92: func() {
		// RES 2, D
		gb.cpu.set8Reg(REG_D, clearBit(gb.cpu.read8Reg(REG_D), 2))
	},
	0x93: func() {
		// RES 2, E
		gb.cpu.set8Reg(REG_E, clearBit(gb.cpu.read8Reg(REG_E), 2))
	},
	0x94: func() {
		// RES 2, H
		gb.cpu.set8Reg(REG_H, clearBit(gb.cpu.read8Reg(REG_H), 2))
	},
	0x95: func() {
		// RES 2, L
		gb.cpu.set8Reg(REG_L, clearBit(gb.cpu.read8Reg(REG_L), 2))
	},
	0x96: func() {
		// RES 2, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, clearBit(val, 2))
	},
	0x97: func() {
		// RES 2, A
		gb.cpu.set8Reg(REG_A, clearBit(gb.cpu.read8Reg(REG_A), 2))
	},
	0x98: func() {
		// RES 3, B
		gb.cpu.set8Reg(REG_B, clearBit(gb.cpu.read8Reg(REG_B), 3))
	},
	0x99: func() {
		// RES 3, C
		gb.cpu.set8Reg(REG_C, clearBit(gb.cpu.read8Reg(REG_C), 3))
	},
	0x9a: func() {
		// RES 3, D
		gb.cpu.set8Reg(REG_D, clearBit(gb.cpu.read8Reg(REG_D), 3))
	},
	0x9b: func() {
		// RES 3, E
		gb.cpu.set8Reg(REG_E, clearBit(gb.cpu.read8Reg(REG_E), 3))
	},
	0x9c: func() {
		// RES 3, H
		gb.cpu.set8Reg(REG_H, clearBit(gb.cpu.read8Reg(REG_H), 3))
	},
	0x9d: func() {
		// RES 3, L
		gb.cpu.set8Reg(REG_L, clearBit(gb.cpu.read8Reg(REG_L), 3))
	},
	0x9e: func() {
		// RES 3, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, clearBit(val, 3))
	},
	0x9f: func() {
		// RES 3, A
		gb.cpu.set8Reg(REG_A, clearBit(gb.cpu.read8Reg(REG_A), 3))
	},

	0xa0: func() {
		// RES 4, B
		gb.cpu.set8Reg(REG_B, clearBit(gb.cpu.read8Reg(REG_B), 4))
	},
	0xa1: func() {
		// RES 4, C
		gb.cpu.set8Reg(REG_C, clearBit(gb.cpu.read8Reg(REG_C), 4))
	},
	0xa2: func() {
		// RES 4, D
		gb.cpu.set8Reg(REG_D, clearBit(gb.cpu.read8Reg(REG_D), 4))
	},
	0xa3: func() {
		// RES 4, E
		gb.cpu.set8Reg(REG_E, clearBit(gb.cpu.read8Reg(REG_E), 4))
	},
	0xa4: func() {
		// RES 4, H
		gb.cpu.set8Reg(REG_H, clearBit(gb.cpu.read8Reg(REG_H), 4))
	},
	0xa5: func() {
		// RES 4, L
		gb.cpu.set8Reg(REG_L, clearBit(gb.cpu.read8Reg(REG_L), 4))
	},
	0xa6: func() {
		// RES 4, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, clearBit(val, 4))
	},
	0xa7: func() {
		// RES 4, A
		gb.cpu.set8Reg(REG_A, clearBit(gb.cpu.read8Reg(REG_A), 4))
	},
	0xa8: func() {
		// RES 5, B
		gb.cpu.set8Reg(REG_B, clearBit(gb.cpu.read8Reg(REG_B), 5))
	},
	0xa9: func() {
		// RES 5, C
		gb.cpu.set8Reg(REG_C, clearBit(gb.cpu.read8Reg(REG_C), 5))
	},
	0xaa: func() {
		// RES 5, D
		gb.cpu.set8Reg(REG_D, clearBit(gb.cpu.read8Reg(REG_D), 5))
	},
	0xab: func() {
		// RES 5, E
		gb.cpu.set8Reg(REG_E, clearBit(gb.cpu.read8Reg(REG_E), 5))
	},
	0xac: func() {
		// RES 5, H
		gb.cpu.set8Reg(REG_H, clearBit(gb.cpu.read8Reg(REG_H), 5))
	},
	0xad: func() {
		// RES 5, L
		gb.cpu.set8Reg(REG_L, clearBit(gb.cpu.read8Reg(REG_L), 5))
	},
	0xae: func() {
		// RES 5, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, clearBit(val, 5))
	},
	0xaf: func() {
		// RES 5, A
		gb.cpu.set8Reg(REG_A, clearBit(gb.cpu.read8Reg(REG_A), 5))
	},

	0xb0: func() {
		// RES 6, B
		gb.cpu.set8Reg(REG_B, clearBit(gb.cpu.read8Reg(REG_B), 6))
	},
	0xb1: func() {
		// RES 6, C
		gb.cpu.set8Reg(REG_C, clearBit(gb.cpu.read8Reg(REG_C), 6))
	},
	0xb2: func() {
		// RES 6, D
		gb.cpu.set8Reg(REG_D, clearBit(gb.cpu.read8Reg(REG_D), 6))
	},
	0xb3: func() {
		// RES 6, E
		gb.cpu.set8Reg(REG_E, clearBit(gb.cpu.read8Reg(REG_E), 6))
	},
	0xb4: func() {
		// RES 6, H
		gb.cpu.set8Reg(REG_H, clearBit(gb.cpu.read8Reg(REG_H), 6))
	},
	0xb5: func() {
		// RES 6, L
		gb.cpu.set8Reg(REG_L, clearBit(gb.cpu.read8Reg(REG_L), 6))
	},
	0xb6: func() {
		// RES 6, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, clearBit(val, 6))
	},
	0xb7: func() {
		// RES 6, A
		gb.cpu.set8Reg(REG_A, clearBit(gb.cpu.read8Reg(REG_A), 6))
	},
	0xb8: func() {
		// RES 7, B
		gb.cpu.set8Reg(REG_B, clearBit(gb.cpu.read8Reg(REG_B), 7))
	},
	0xb9: func() {
		// RES 7, C
		gb.cpu.set8Reg(REG_C, clearBit(gb.cpu.read8Reg(REG_C), 7))
	},
	0xba: func() {
		// RES 7, D
		gb.cpu.set8Reg(REG_D, clearBit(gb.cpu.read8Reg(REG_D), 7))
	},
	0xbb: func() {
		// RES 7, E
		gb.cpu.set8Reg(REG_E, clearBit(gb.cpu.read8Reg(REG_E), 7))
	},
	0xbc: func() {
		// RES 7, H
		gb.cpu.set8Reg(REG_H, clearBit(gb.cpu.read8Reg(REG_H), 7))
	},
	0xbd: func() {
		// RES 7, L
		gb.cpu.set8Reg(REG_L, clearBit(gb.cpu.read8Reg(REG_L), 7))
	},
	0xbe: func() {
		// RES 7, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, clearBit(val, 7))
	},
	0xbf: func() {
		// RES 7, A
		gb.cpu.set8Reg(REG_A, clearBit(gb.cpu.read8Reg(REG_A), 7))
	},

	0xc0: func() {
		// SET 0, B
		gb.cpu.set8Reg(REG_B, setBit(gb.cpu.read8Reg(REG_B), 0))
	},
	0xc1: func() {
		// SET 0, C
		gb.cpu.set8Reg(REG_C, setBit(gb.cpu.read8Reg(REG_C), 0))
	},
	0xc2: func() {
		// SET 0, D
		gb.cpu.set8Reg(REG_D, setBit(gb.cpu.read8Reg(REG_D), 0))
	},
	0xc3: func() {
		// SET 0, E
		gb.cpu.set8Reg(REG_E, setBit(gb.cpu.read8Reg(REG_E), 0))
	},
	0xc4: func() {
		// SET 0, H
		gb.cpu.set8Reg(REG_H, setBit(gb.cpu.read8Reg(REG_H), 0))
	},
	0xc5: func() {
		// SET 0, L
		gb.cpu.set8Reg(REG_L, setBit(gb.cpu.read8Reg(REG_L), 0))
	},
	0xc6: func() {
		// SET 0, (HL)
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, setBit(val, 0))
	},
	0xc7: func() {
		// SET 0, A
		gb.cpu.set8Reg(REG_A, setBit(gb.cpu.read8Reg(REG_A), 0))
	},
	0xc8: func() {
		// SET 1, B
		gb.cpu.set8Reg(REG_B, setBit(gb.cpu.read8Reg(REG_B), 1))
	},
	0xc9: func() {
		// SET 1, C
		gb.cpu.set8Reg(REG_C, setBit(gb.cpu.read8Reg(REG_C), 1))
	},
	0xca: func() {
		// SET 1, D
		gb.cpu.set8Reg(REG_D, setBit(gb.cpu.read8Reg(REG_D), 1))
	},
	0xcb: func() {
		// SET 1, E
		gb.cpu.set8Reg(REG_E, setBit(gb.cpu.read8Reg(REG_E), 1))
	},
	0xcc: func() {
		// SET 1, H
		gb.cpu.set8Reg(REG_H, setBit(gb.cpu.read8Reg(REG_H), 1))
	},
	0xcd: func() {
		// SET 1, L
		gb.cpu.set8Reg(REG_L, setBit(gb.cpu.read8Reg(REG_L), 1))
	},
	0xce: func() {
		// SET 1, HL
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, setBit(val, 1))
	},
	0xcf: func() {
		// SET 1, A
		gb.cpu.set8Reg(REG_A, setBit(gb.cpu.read8Reg(REG_A), 1))
	},

	0xd0: func() {
		// SET 2, B
		gb.cpu.set8Reg(REG_B, setBit(gb.cpu.read8Reg(REG_B), 2))
	},
	0xd1: func() {
		// SET 2, C
		gb.cpu.set8Reg(REG_C, setBit(gb.cpu.read8Reg(REG_C), 2))
	},
	0xd2: func() {
		// SET 2, D
		gb.cpu.set8Reg(REG_D, setBit(gb.cpu.read8Reg(REG_D), 2))
	},
	0xd3: func() {
		// SET 2, E
		gb.cpu.set8Reg(REG_E, setBit(gb.cpu.read8Reg(REG_E), 2))
	},
	0xd4: func() {
		// SET 2, H
		gb.cpu.set8Reg(REG_H, setBit(gb.cpu.read8Reg(REG_H), 2))
	},
	0xd5: func() {
		// SET 2, L
		gb.cpu.set8Reg(REG_L, setBit(gb.cpu.read8Reg(REG_L), 2))
	},
	0xd6: func() {
		// SET 02 HL
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, setBit(val, 2))
	},
	0xd7: func() {
		// SET 2, A
		gb.cpu.set8Reg(REG_A, setBit(gb.cpu.read8Reg(REG_A), 2))
	},
	0xd8: func() {
		// SET 3, B
		gb.cpu.set8Reg(REG_B, setBit(gb.cpu.read8Reg(REG_B), 3))
	},
	0xd9: func() {
		// SET 3, C
		gb.cpu.set8Reg(REG_C, setBit(gb.cpu.read8Reg(REG_C), 3))
	},
	0xda: func() {
		// SET 3, D
		gb.cpu.set8Reg(REG_D, setBit(gb.cpu.read8Reg(REG_D), 3))
	},
	0xdb: func() {
		// SET 3, E
		gb.cpu.set8Reg(REG_E, setBit(gb.cpu.read8Reg(REG_E), 3))
	},
	0xdc: func() {
		// SET 3, H
		gb.cpu.set8Reg(REG_H, setBit(gb.cpu.read8Reg(REG_H), 3))
	},
	0xdd: func() {
		// SET 3, L
		gb.cpu.set8Reg(REG_L, setBit(gb.cpu.read8Reg(REG_L), 3))
	},
	0xde: func() {
		// SET 3, HL
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, setBit(val, 3))
	},
	0xdf: func() {
		// SET 3, A
		gb.cpu.set8Reg(REG_A, setBit(gb.cpu.read8Reg(REG_A), 3))
	},

	0xe0: func() {
		// SET 4, B
		gb.cpu.set8Reg(REG_B, setBit(gb.cpu.read8Reg(REG_B), 4))
	},
	0xe1: func() {
		// SET 4, C
		gb.cpu.set8Reg(REG_C, setBit(gb.cpu.read8Reg(REG_C), 4))
	},
	0xe2: func() {
		// SET 4, D
		gb.cpu.set8Reg(REG_D, setBit(gb.cpu.read8Reg(REG_D), 4))
	},
	0xe3: func() {
		// SET 4, E
		gb.cpu.set8Reg(REG_E, setBit(gb.cpu.read8Reg(REG_E), 4))
	},
	0xe4: func() {
		// SET 4, H
		gb.cpu.set8Reg(REG_H, setBit(gb.cpu.read8Reg(REG_H), 4))
	},
	0xe5: func() {
		// SET 4, L
		gb.cpu.set8Reg(REG_L, setBit(gb.cpu.read8Reg(REG_L), 4))
	},
	0xe6: func() {
		// SET 04 HL
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, setBit(val, 4))
	},
	0xe7: func() {
		// SET 4, A
		gb.cpu.set8Reg(REG_A, setBit(gb.cpu.read8Reg(REG_A), 4))
	},
	0xe8: func() {
		// SET 5, B
		gb.cpu.set8Reg(REG_B, setBit(gb.cpu.read8Reg(REG_B), 5))
	},
	0xe9: func() {
		// SET 5, C
		gb.cpu.set8Reg(REG_C, setBit(gb.cpu.read8Reg(REG_C), 5))
	},
	0xea: func() {
		// SET 5, D
		gb.cpu.set8Reg(REG_D, setBit(gb.cpu.read8Reg(REG_D), 5))
	},
	0xeb: func() {
		// SET 5, E
		gb.cpu.set8Reg(REG_E, setBit(gb.cpu.read8Reg(REG_E), 5))
	},
	0xec: func() {
		// SET 5, H
		gb.cpu.set8Reg(REG_H, setBit(gb.cpu.read8Reg(REG_H), 5))
	},
	0xed: func() {
		// SET 5, L
		gb.cpu.set8Reg(REG_L, setBit(gb.cpu.read8Reg(REG_L), 5))
	},
	0xee: func() {
		// SET 5, HL
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, setBit(val, 5))
	},
	0xef: func() {
		// SET 5, A
		gb.cpu.set8Reg(REG_A, setBit(gb.cpu.read8Reg(REG_A), 5))
	},

	0xf0: func() {
		// SET 6, B
		gb.cpu.set8Reg(REG_B, setBit(gb.cpu.read8Reg(REG_B), 6))
	},
	0xf1: func() {
		// SET 6, C
		gb.cpu.set8Reg(REG_C, setBit(gb.cpu.read8Reg(REG_C), 6))
	},
	0xf2: func() {
		// SET 6, D
		gb.cpu.set8Reg(REG_D, setBit(gb.cpu.read8Reg(REG_D), 6))
	},
	0xf3: func() {
		// SET 6, E
		gb.cpu.set8Reg(REG_E, setBit(gb.cpu.read8Reg(REG_E), 6))
	},
	0xf4: func() {
		// SET 6, H
		gb.cpu.set8Reg(REG_H, setBit(gb.cpu.read8Reg(REG_H), 6))
	},
	0xf5: func() {
		// SET 6, L
		gb.cpu.set8Reg(REG_L, setBit(gb.cpu.read8Reg(REG_L), 6))
	},
	0xf6: func() {
		// SET 6, HL
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, setBit(val, 6))
	},
	0xf7: func() {
		// SET 6, A
		gb.cpu.set8Reg(REG_A, setBit(gb.cpu.read8Reg(REG_A), 6))
	},
	0xf8: func() {
		// SET 7, B
		gb.cpu.set8Reg(REG_B, setBit(gb.cpu.read8Reg(REG_B), 7))
	},
	0xf9: func() {
		// SET 7, C
		gb.cpu.set8Reg(REG_C, setBit(gb.cpu.read8Reg(REG_C), 7))
	},
	0xfa: func() {
		// SET 7, D
		gb.cpu.set8Reg(REG_D, setBit(gb.cpu.read8Reg(REG_D), 7))
	},
	0xfb: func() {
		// SET 7, E
		gb.cpu.set8Reg(REG_E, setBit(gb.cpu.read8Reg(REG_E), 7))
	},
	0xfc: func() {
		// SET 7, H
		gb.cpu.set8Reg(REG_H, setBit(gb.cpu.read8Reg(REG_H), 7))
	},
	0xfd: func() {
		// SET 7, L
		gb.cpu.set8Reg(REG_L, setBit(gb.cpu.read8Reg(REG_L), 7))
	},
	0xfe: func() {
		// SET 7, HL
		hl := gb.cpu.read16Reg(REG_HL)
		val := gb.bus.read(hl)
		gb.bus.write(hl, setBit(val, 7))
	},
	0xff: func() {
		// SET 7, A
		gb.cpu.set8Reg(REG_A, setBit(gb.cpu.read8Reg(REG_A), 7))
	},
}
