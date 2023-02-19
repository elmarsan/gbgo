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

	0x40: func() {},
	0x41: func() {}, // LD B, C
	0x42: func() {}, // LD B, E
	0x43: func() {}, // LD B, E
	0x44: func() {}, // LD B, H
	0x45: func() {}, // LD B, L
	0x46: func() {}, // LD B, HL
	0x47: func() {},
	0x48: func() {},
	0x49: func() {},
	0x4a: func() {},
	0x4b: func() {},
	0x4c: func() {},
	0x4d: func() {},
	0x4e: func() {},
	0x4f: func() {},

	0x50: func() {},
	0x51: func() {},
	0x52: func() {},
	0x53: func() {},
	0x54: func() {},
	0x55: func() {},
	0x56: func() {},
	0x57: func() {},
	0x58: func() {},
	0x59: func() {},
	0x5a: func() {},
	0x5b: func() {},
	0x5c: func() {},
	0x5d: func() {},
	0x5e: func() {},
	0x5f: func() {},

	0x60: func() {},
	0x61: func() {},
	0x62: func() {},
	0x63: func() {},
	0x64: func() {},
	0x65: func() {},
	0x66: func() {},
	0x67: func() {},
	0x68: func() {},
	0x69: func() {},
	0x6a: func() {},
	0x6b: func() {},
	0x6c: func() {},
	0x6d: func() {},
	0x6e: func() {},
	0x6f: func() {},

	0x70: func() {},
	0x71: func() {},
	0x72: func() {},
	0x73: func() {},
	0x74: func() {},
	0x75: func() {},
	0x76: func() {},
	0x77: func() {},
	0x78: func() {},
	0x79: func() {},
	0x7a: func() {},
	0x7b: func() {},
	0x7c: func() {},
	0x7d: func() {},
	0x7e: func() {},
	0x7f: func() {},

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
