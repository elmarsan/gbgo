package main

var opcodes = [0x100]func(){
	0x00: func() {}, // NOP
	0x01: func() {}, // LD BC, d16
	0x02: func() {}, // LD (BC), A
	0x03: func() {}, // INC BC
	0x04: func() {}, // INC B
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
}
