package main

var instructions = [0x100]func(){
	0x00: func() {}, // NOP
	0x01: func() {}, // LD BC, d16
	0x02: func() {}, // LD (BC), A
	0x03: func() {}, // INC BC
	0x04: func() {}, // INC BW
	0x05: func() {}, // DEC B
	0x06: func() {}, // LD B, d8
	0x07: func() {}, // RLCA
	0x08: func() {}, // LD (a16), SP
	0x09: func() {}, // ADD HL, BC
	0x0a: func() {}, // LD A, (BC)
	0x0b: func() {}, // DEC BC
	0x0c: func() {}, // INC C
	0x0d: func() {}, // DEC C
	0x0e: func() {}, // LD C, d8
	0x0f: func() {}, // RRCA

	0x10: func() {}, // STOP d8
	0x11: func() {}, // LD DE, d16
	0x12: func() {}, // LD (DE), A
	0x13: func() {}, // INC DE
	0x14: func() {}, // INC D
	0x15: func() {}, // DEC D
	0x16: func() {}, // LD D, d8
	0x17: func() {}, // RLA
	0x18: func() {}, // JR r8
	0x19: func() {}, // ADD HL, DE
	0x1a: func() {}, // LD A, (DE)
	0x1b: func() {}, // DEC DE
	0x1c: func() {}, // INC E
	0x1d: func() {}, // DEC E
	0x1e: func() {}, // LD E, d8
	0x1f: func() {}, // RRA

	0x20: func() {}, // JR NZ, r8
	0x21: func() {}, // LD HL, d16
	0x22: func() {}, // LD (HL+), A
	0x23: func() {}, // INC HL
	0x24: func() {}, // INC H
	0x25: func() {}, // DEC H
	0x26: func() {}, // LD H, d8
	0x27: func() {}, // DAA
	0x28: func() {}, // JR Z, r8
	0x29: func() {}, // ADD HL, HL
	0x2a: func() {}, // LD A, (HL+)
	0x2b: func() {}, // DEC HL
	0x2c: func() {}, // INC L
	0x2d: func() {}, // DEC L
	0x2e: func() {}, // LD L, d8
	0x2f: func() {}, // CPL

	0x30: func() {}, // JR NC, r8
	0x31: func() {}, // LD SP, d16
	0x32: func() {}, // LD (HL-), A
	0x33: func() {}, // INC SP
	0x34: func() {}, // INC (HL)
	0x35: func() {}, // DEC (HL)
	0x36: func() {}, // LD (HL), d8
	0x37: func() {}, // SCF
	0x38: func() {}, // JR C, r8
	0x39: func() {}, // ADD HL, SP
	0x3a: func() {}, // LD A, (HL-)
	0x3b: func() {}, // DEC SP
	0x3c: func() {}, // INC A
	0x3d: func() {}, // DEC A
	0x3e: func() {}, // LD A, d8
	0x3f: func() {}, // CCF

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
	},
	0x6f: func() {
		// LD L, A
		cpu.load8Reg(REG_L, REG_A)
	},

	0x70: func() {
		// LD (HL), B
	},
	0x71: func() {
		// LD (HL), C
	},
	0x72: func() {
		// LD (HL), D
	},
	0x73: func() {
		// LD (HL), E
	},
	0x74: func() {
		// LD (HL), H
	},
	0x75: func() {
		// LD (HL), L
	},
	0x76: func() {
		// HALT
	},
	0x77: func() {
		// LD (HL), A
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
	},
	0x7f: func() {
		// LD A, A
		cpu.load8Reg(REG_A, REG_A)
	},

	0x80: func() {},
	0x81: func() {},
	0x82: func() {},
	0x83: func() {},
	0x84: func() {},
	0x85: func() {},
	0x86: func() {},
	0x87: func() {},
	0x88: func() {},
	0x89: func() {},
	0x8a: func() {},
	0x8b: func() {},
	0x8c: func() {},
	0x8d: func() {},
	0x8e: func() {},
	0x8f: func() {},

	0x90: func() {},
	0x91: func() {},
	0x92: func() {},
	0x93: func() {},
	0x94: func() {},
	0x95: func() {},
	0x96: func() {},
	0x97: func() {},
	0x98: func() {},
	0x99: func() {},
	0x9a: func() {},
	0x9b: func() {},
	0x9c: func() {},
	0x9d: func() {},
	0x9e: func() {},
	0x9f: func() {},

	0xa0: func() {},
	0xa1: func() {},
	0xa2: func() {},
	0xa3: func() {},
	0xa4: func() {},
	0xa5: func() {},
	0xa6: func() {},
	0xa7: func() {},
	0xa8: func() {},
	0xa9: func() {},
	0xaa: func() {},
	0xab: func() {},
	0xac: func() {},
	0xad: func() {},
	0xae: func() {},
	0xaf: func() {},

	0xb0: func() {},
	0xb1: func() {},
	0xb2: func() {},
	0xb3: func() {},
	0xb4: func() {},
	0xb5: func() {},
	0xb6: func() {},
	0xb7: func() {},
	0xb8: func() {},
	0xb9: func() {},
	0xba: func() {},
	0xbb: func() {},
	0xbc: func() {},
	0xbd: func() {},
	0xbe: func() {},
	0xbf: func() {},

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
	0xce: func() {},
	0xcf: func() {},

	0xd0: func() {},
	0xd1: func() {},
	0xd2: func() {},
	0xd3: func() {},
	0xd4: func() {},
	0xd5: func() {},
	0xd6: func() {},
	0xd7: func() {},
	0xd8: func() {},
	0xd9: func() {},
	0xda: func() {},
	0xdb: func() {},
	0xdc: func() {},
	0xdd: func() {},
	0xde: func() {},
	0xdf: func() {},

	0xe0: func() {},
	0xe1: func() {},
	0xe2: func() {},
	0xe3: func() {},
	0xe4: func() {},
	0xe5: func() {},
	0xe6: func() {},
	0xe7: func() {},
	0xe8: func() {},
	0xe9: func() {},
	0xea: func() {},
	0xeb: func() {},
	0xec: func() {},
	0xed: func() {},
	0xee: func() {},
	0xef: func() {},

	0xf0: func() {},
	0xf1: func() {},
	0xf2: func() {},
	0xf3: func() {},
	0xf4: func() {},
	0xf5: func() {},
	0xf6: func() {},
	0xf7: func() {},
	0xf8: func() {},
	0xf9: func() {},
	0xfa: func() {},
	0xfb: func() {},
	0xfc: func() {},
	0xfd: func() {},
	0xfe: func() {},
	0xff: func() {},
}
