package main

import "log"

// CPU represents gameboy central processing unit.
type CPU struct {
	a uint8
	f uint8
	b uint8
	c uint8
	d uint8
	e uint8
	h uint8
	l uint8

	sp uint16
	pc uint16

	zFlag bool
	nFlag bool
	hFlag bool
	cFlag bool
}

// CPU8Register represents 8 bit register.
type CPU8Register int

// 8 bits registers
const (
	REG_A CPU8Register = iota
	REG_F
	REG_B
	REG_C
	REG_D
	REG_E
	REG_H
	REG_L
)

// CPU16Register represents 16 bit register.
type CPU16Register int

// 16 bits registers
const (
	REG_AF = iota
	REG_BC
	REG_DE
	REG_HL
)

var cpu = &CPU{
	pc: 0x100,
}

// cycle executes next instruction.
func (cpu *CPU) cycle() {
	opcode := memory.read(cpu.readPc())
	log.Printf("Executing opcode 0x%x", opcode)
	instructions[opcode]()
}

// readPc returns current value of pc and increments it.
func (cpu *CPU) readPc() uint16 {
	pc := cpu.pc
	cpu.pc++

	return pc
}

// read8Reg reads and returns 8 bit register value.
func (cpu *CPU) read8Reg(reg CPU8Register) uint8 {
	switch reg {
	case REG_A:
		return cpu.a
	case REG_F:
		return cpu.f
	case REG_B:
		return cpu.b
	case REG_C:
		return cpu.c
	case REG_D:
		return cpu.d
	case REG_E:
		return cpu.e
	case REG_H:
		return cpu.h
	case REG_L:
		return cpu.l
	default:
		log.Fatalf("Unknown register %s", reg)
		return 0
	}
}

// read16Reg reads and returns 16 bit register value.
func (cpu *CPU) read16Reg(reg CPU16Register) uint16 {
	switch reg {
	case REG_AF:
		return joinUint8(cpu.a, cpu.f)
	case REG_BC:
		return joinUint8(cpu.b, cpu.c)
	case REG_DE:
		return joinUint8(cpu.d, cpu.e)
	case REG_HL:
		return joinUint8(cpu.h, cpu.l)
	default:
		log.Fatalf("Unknown register %s", reg)
		return 0
	}
}

// set8Reg sets the value of 8 bit register
func (cpu *CPU) set8Reg(reg CPU8Register, val uint8) {
	switch reg {
	case REG_A:
		cpu.a = val
		break
	case REG_F:
		cpu.f = val
		break
	case REG_B:
		cpu.b = val
		break
	case REG_C:
		cpu.c = val
		break
	case REG_D:
		cpu.d = val
		break
	case REG_E:
		cpu.e = val
		break
	case REG_H:
		cpu.h = val
		break
	case REG_L:
		cpu.l = val
		break
	default:
		log.Fatalf("Unknown register %s", reg)
		break
	}
}

// set16Reg sets the value of 16 bit register
func (cpu *CPU) set16Reg(reg CPU16Register, val uint16) {
	switch reg {
	case REG_AF:
		cpu.a = hi(val)
		cpu.f = lo(val)
		break
	case REG_BC:
		cpu.b = hi(val)
		cpu.c = lo(val)
		break
	case REG_DE:
		cpu.d = hi(val)
		cpu.e = lo(val)
		break
	case REG_HL:
		cpu.h = hi(val)
		cpu.l = lo(val)
		break
	default:
		log.Fatalf("Unknown register %s", reg)
		break
	}
}

// load8Reg loads b into a.
func (cpu *CPU) load8Reg(a CPU8Register, b CPU8Register) {
	var val uint8 = cpu.read8Reg(b)
	cpu.set8Reg(a, val)
}

// inc8Reg increments 8 bit register.
// It stores in a register a (a + 1) and sets flags.
func (cpu *CPU) inc8Reg(a CPU8Register) {
	val := cpu.read8Reg(a)
	inc := val + 1

	cpu.set8Reg(a, inc)

	cpu.zFlag = val == 0
	cpu.nFlag = false
	cpu.hFlag = (val & 0x0f) == 0
}

// dec8Reg decrements 8 bit register.
// It stores in a register a (a - 1) and sets flags.
func (cpu *CPU) dec8Reg(reg CPU8Register) {
	val := cpu.read8Reg(reg)
	val -= 1

	cpu.set8Reg(reg, val)

	cpu.zFlag = val == 0
	cpu.nFlag = true
	cpu.hFlag = (val & 0x0f) == 0
}

// inc16reg increments 16 bit register.
// It stores in a register a (a + 1) and sets flags.
func (cpu *CPU) inc16reg(reg CPU16Register, setFlags bool) {
	val := cpu.read16Reg(reg)
	val += 1
	cpu.set16Reg(reg, val)

	if setFlags {
		cpu.zFlag = val == 0
		cpu.nFlag = true
		cpu.hFlag = (val & 0x0f) == 0
	}
}

// dec16Reg decrements 16 bit register.
// It stores in a register a (a - 1) and sets flags.
func (cpu *CPU) dec16Reg(a CPU16Register) {
	val := cpu.read16Reg(a)
	dec := val - 1

	cpu.set16Reg(a, dec)

	cpu.zFlag = dec == 0
	cpu.nFlag = true
	cpu.hFlag = (dec & 0x0f) > 0
}

// add8RegD8 adds b to a.
// It stores in a register a (a + b) and sets flags.
func (cpu *CPU) add8Reg(a CPU8Register, b CPU8Register) {
	add := cpu.read8Reg(a) + cpu.read8Reg(b)
	cpu.set8Reg(a, add)

	cpu.zFlag = add == 0
	cpu.nFlag = false
	cpu.cFlag = add > 0xff
	cpu.hFlag = add > 0xf
}

// add8RegD8 add d8 into a.
// It stores in a register a (a + d8) and sets flags.
func (cpu *CPU) add8RegD8(a CPU8Register, d8 uint8) {
	add := cpu.read8Reg(a) + d8
	cpu.set8Reg(a, add)

	cpu.zFlag = add == 0
	cpu.nFlag = false
	cpu.cFlag = add > 0xff
	cpu.hFlag = (add & 0x0f) > 0xf
}

// adc8Reg add b + carry flag into a.
// It stores in a register a (a + b + carry flag) and sets flags.
func (cpu *CPU) adc8Reg(a CPU8Register, b CPU8Register) {
	add := cpu.read8Reg(a) + cpu.read8Reg(b)
	if cpu.cFlag {
		add += 1
	}
	cpu.set8Reg(a, add)

	cpu.zFlag = add == 0
	cpu.nFlag = false
	cpu.cFlag = add > 0xff
	cpu.hFlag = (add & 0x0f) > 0xf
}

// adc8RegD8 add d8 + carry flag into a.
// It stores in a register a (a + d8 + carry flag) and sets flags.
func (cpu *CPU) adc8RegD8(a CPU8Register, d8 uint8) {
	add := cpu.read8Reg(a) + d8
	if cpu.cFlag {
		add += 1
	}
	cpu.set8Reg(a, add)

	cpu.zFlag = add == 0
	cpu.nFlag = false
	cpu.cFlag = add > 0xff
	cpu.hFlag = (add & 0x0f) > 0xf
}

// sub8Reg subtracts b to a.
// It stores in register a (a - b) and sets flags.
func (cpu *CPU) sub8Reg(a CPU8Register, b CPU8Register) {
	sub := cpu.read8Reg(a) - cpu.read8Reg(b)
	cpu.set8Reg(a, sub)

	cpu.zFlag = sub == 0
	cpu.nFlag = true
	cpu.cFlag = sub > 0xff
	cpu.hFlag = (sub & 0x0f) == 0
}

// sub8RegD8 subtracts d8 to a.
// It stores in register a (a - d8) and sets flags.
func (cpu *CPU) sub8RegD8(a CPU8Register, d8 uint8) {
	val := cpu.read8Reg(a)
	sub := val - d8
	cpu.set8Reg(a, sub)

	cpu.zFlag = sub == 0
	cpu.nFlag = true
	cpu.cFlag = sub > 0xff
	cpu.hFlag = (sub & 0x0f) == 0
}

// sbc8Reg subtracts d8 and carry flag to a.
// It stores in register a (a - d8 - carry flag) and sets flags.
func (cpu *CPU) sbc8Reg(a CPU8Register, b CPU8Register) {
	sub := cpu.read8Reg(a) - cpu.read8Reg(b)
	if cpu.cFlag {
		sub -= 1
	}
	cpu.set8Reg(a, sub)

	cpu.zFlag = sub == 0
	cpu.nFlag = true
	cpu.cFlag = sub > 0xff
	cpu.hFlag = (sub & 0x0f) == 0
}

// sbc8RegD8 subtracts d8 and carry flag to a.
// It stores in register a (a - d8) and sets flags.
func (cpu *CPU) sbc8RegD8(a CPU8Register, d8 uint8) {
	val := cpu.read8Reg(a)
	sub := val - d8
	if cpu.cFlag {
		sub -= 1
	}
	cpu.set8Reg(a, sub)

	cpu.zFlag = sub == 0
	cpu.nFlag = true
	cpu.cFlag = sub > 0xff
	cpu.hFlag = (sub & 0x0f) == 0
}

// and8Reg performs bitwise AND between a and b.
// It stores in register a (a & b) and set flags.
func (cpu *CPU) and8Reg(a CPU8Register, b CPU8Register) {
	and := cpu.read8Reg(a) & cpu.read8Reg(b)
	cpu.set8Reg(a, and)

	cpu.zFlag = and == 0
	cpu.nFlag = false
	cpu.cFlag = false
	cpu.hFlag = true
}

// and8RegD8 performs bitwise AND between a and d8.
// It stores in register a (a & d8) and set flags.
func (cpu *CPU) and8RegD8(a CPU8Register, d8 uint8) {
	and := cpu.read8Reg(a) & d8
	cpu.set8Reg(a, and)

	cpu.zFlag = and == 0
	cpu.nFlag = false
	cpu.cFlag = false
	cpu.hFlag = true
}

// xor8Reg performs bitwise XOR between a and b.
// It stores in register a (a ^ b) and set flags.
func (cpu *CPU) xor8Reg(a CPU8Register, b CPU8Register) {
	xorg := cpu.read8Reg(a) ^ cpu.read8Reg(b)
	cpu.set8Reg(a, xorg)

	cpu.zFlag = xorg == 0
	cpu.nFlag = false
	cpu.cFlag = false
	cpu.hFlag = false
}

// xor8RegD8 performs bitwise XOR between a and d8.
// It stores in register a (a ^ d8) and set flags.
func (cpu *CPU) xor8RegD8(a CPU8Register, d8 uint8) {
	xorg := cpu.read8Reg(a) ^ d8
	cpu.set8Reg(a, xorg)

	cpu.zFlag = xorg == 0
	cpu.nFlag = false
	cpu.cFlag = false
	cpu.hFlag = false
}

// or8Reg performs bitwise OR between a and b.
// It stores in register a (a | b) and set flags.
func (cpu *CPU) or8Reg(a CPU8Register, b CPU8Register) {
	xorg := cpu.read8Reg(a) | cpu.read8Reg(b)
	cpu.set8Reg(a, xorg)

	cpu.zFlag = xorg == 0
	cpu.nFlag = false
	cpu.cFlag = false
	cpu.hFlag = false
}

// or8RegD8 performs bitwise OR between a and d8.
// It stores in register a (a | d8) and set flags.
func (cpu *CPU) or8RegD8(a CPU8Register, d8 uint8) {
	xorg := cpu.read8Reg(a) | d8
	cpu.set8Reg(a, xorg)

	cpu.zFlag = xorg == 0
	cpu.nFlag = false
	cpu.cFlag = false
	cpu.hFlag = false
}

// cp8Reg compares the values of register a and b.
// It set flags.
func (cpu *CPU) cp8Reg(a CPU8Register, b CPU8Register) {
	sub := cpu.read8Reg(a) - cpu.read8Reg(b)

	cpu.zFlag = sub == 0
	cpu.nFlag = true
	cpu.cFlag = sub > 0xff
	cpu.hFlag = (sub & 0x0f) == 0
}

// cp8RegD8 compares the values of register a and d8.
// It set flags.
func (cpu *CPU) cp8RegD8(a CPU8Register, d8 uint8) {
	sub := cpu.read8Reg(a) - d8

	cpu.zFlag = sub == 0
	cpu.nFlag = true
	cpu.cFlag = sub > 0xff
	cpu.hFlag = (sub & 0x0f) == 0
}
