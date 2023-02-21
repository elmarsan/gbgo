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
}

// CPU8Register represents 8 bit register.
type CPU8Register string

// CPU16Register represents 16 bit register.
type CPU16Register string

const (
	REG_A CPU8Register = "REG_A"
	REG_F CPU8Register = "REG_F"
	REG_B CPU8Register = "REG_B"
	REG_C CPU8Register = "REG_C"
	REG_D CPU8Register = "REG_D"
	REG_E CPU8Register = "REG_E"
	REG_H CPU8Register = "REG_H"
	REG_L CPU8Register = "REG_L"

	REG_AF CPU16Register = "REG_AF"
	REG_BC CPU16Register = "REG_BC"
	REG_DE CPU16Register = "REG_DE"
	REG_HL CPU16Register = "REG_HL"
)

var cpu = &CPU{
	pc: 0x100,
}

// cycle executes next instruction.
func (cpu *CPU) cycle() {
	opcode := memory.read(cpu.pc)
	log.Printf("Executing opcode 0x%x", opcode)
	instructions[opcode]()
	cpu.incPc()
}

// incPc increments pc.
func (cpu *CPU) incPc() {
	cpu.pc++
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
