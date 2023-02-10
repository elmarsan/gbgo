package main

// register represent a cpu register of 16 bytes.
type register struct {
	val uint16
}

// Hi returns higher 8 bytes of the register.
func (r *register) Hi() uint8 {
	return byte(r.val >> 8)
}

// Low returns lower 8 bytes of the register.
func (r *register) Low() uint8 {
	return byte(r.val & 0xff)
}

// Set modifies register value.
func (r *register) Set(val uint16) {
	r.val = val
}

// SetHi modifies higher 8 bytes of the register.
func (r *register) SetHi(val uint8) {
	r.val = uint16(val)<<8 | (uint16(r.val) & 0xff)
}

// SetLow modifies lower 8 bytes of the register.
func (r *register) SetLow(val uint8) {
	r.val = uint16(val) | (uint16(r.val) & 0xff00)
}

// CPU represents gameboy central processing unit.
type CPU struct {
	af register
	bc register
	de register
	hl register

	sp register
	pc uint16

	memory [16384]uint8
}

// nextPC8 returns next 8bit val of pc and increment it.
func (cpu *CPU) nextPC8() uint8 {
	opcode := cpu.memory[cpu.pc]
	cpu.pc++
	return opcode
}

// nextPC16 returns next 16bit val of pc.
func (cpu *CPU) nextPC16() uint16 {
	b1 := cpu.nextPC8()
	b2 := cpu.nextPC8()
	return uint16(b2<<8) | uint16(b1)
}
