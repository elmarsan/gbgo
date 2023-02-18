package main

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

var cpu = &CPU{}

func (cpu *CPU) cycle() {
	opcode := read(cpu.pc)
	opcodes[opcode]()
	cpu.incPc()
}

func (cpu *CPU) incPc() {
	cpu.pc += 1
}
