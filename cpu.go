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

// load8Reg loads value of src 8 bits register into dst 98 bits register.
func (cpu *CPU) load8Reg(dst CPU8Register, src CPU8Register) {
	var val uint8 = cpu.read8Reg(src)
	cpu.set8Reg(dst, val)
}

// inc8Reg increments value of 8 bits register and sets flags.
func (cpu *CPU) inc8Reg(reg CPU8Register) {
	val := cpu.read8Reg(reg)
	val += 1

	cpu.set8Reg(reg, val)

	cpu.zFlag = val == 0
	cpu.nFlag = false
	cpu.hFlag = (val & 0x0f) == 0
}

// dec8Reg decrements value of 8 bits register and sets flags.
func (cpu *CPU) dec8Reg(reg CPU8Register) {
	val := cpu.read8Reg(reg)
	val -= 1

	cpu.set8Reg(reg, val)

	cpu.zFlag = val == 0
	cpu.nFlag = true
	cpu.hFlag = (val & 0x0f) == 0
}

// inc16reg increments value of 16 bits register and sets flags if required.
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

// dec16Reg decrements value of 8 bits register and sets flags.
func (cpu *CPU) dec16Reg(reg CPU16Register) {
	val := cpu.read16Reg(reg)
	val -= 1

	cpu.set16Reg(reg, val)

	cpu.zFlag = val == 0
	cpu.nFlag = true
	cpu.hFlag = (val & 0x0f) > 0
}

// add8Reg adds a 8 bit register b 8 bit register and
// stores result back in a register. It also sets cpu flags.
func (cpu *CPU) add8Reg(a CPU8Register, b CPU8Register) {
	add := cpu.read8Reg(a) + cpu.read8Reg(b)
	cpu.set8Reg(a, add)

	cpu.zFlag = add == 0
	cpu.nFlag = false
	cpu.cFlag = add > 0xff
	cpu.hFlag = add > 0xf
}

func (cpu *CPU) add8RegValue(a CPU8Register, val uint8) {
	add := cpu.read8Reg(a) + val
	cpu.set8Reg(a, add)

	cpu.zFlag = add == 0
	cpu.nFlag = false
	cpu.cFlag = add > 0xff
	cpu.hFlag = (add & 0x0f) > 0xf
}

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

func (cpu *CPU) adc8RegValue(a CPU8Register, val uint8) {
	add := cpu.read8Reg(a) + val
	if cpu.cFlag {
		add += 1
	}
	cpu.set8Reg(a, add)

	cpu.zFlag = add == 0
	cpu.nFlag = false
	cpu.cFlag = add > 0xff
	cpu.hFlag = (add & 0x0f) > 0xf
}
