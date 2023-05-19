package gameboy

import "github.com/elmarsan/gbgo/pkg/bit"

// register represent 16 bits cpu register.
type register struct {
	// val holds register value.
	val uint16
}

// Hi returns higher 8 bits of register value.
func (r *register) Hi() uint8 {
	return bit.Hi(r.val)
}

// SetHi modifies higher 8 bits of register value.
func (r *register) SetHi(val uint8) {
	r.val = uint16(val)<<8 | (uint16(r.val) & 0xff)
}

// Lo returns lower 8 bits of register value.
func (r *register) Lo() uint8 {
	return bit.Lo(r.val)
}

// SetLo modifies lower 8 bits of register value.
func (r *register) SetLo(val uint8) {
	r.val = uint16(val) | (uint16(r.val) & 0xff00)
}

// cpu represents gameboy central processing unit.
type cpu struct {
	af register
	bc register
	de register
	hl register

	// sp represents stack pointer register.
	sp uint16

	// pc represents program counter register.
	pc uint16

	// halted indicates if cpu is in low-power state.
	halted bool

	// clockCycles holds clock cycles elapsed during previous instruction execution.
	clockCycles int
}

// newCpu creates and returns a new cpu instance.
// Register values are set to defaults of DMG boot sequence.
func newCpu(bus *memoryBus) *cpu {
	return &cpu{
		af: register{val: 0x01b0},
		bc: register{val: 0x0013},
		de: register{val: 0x00d8},
		hl: register{val: 0x014d},
		sp: 0xfffe,
		pc: 0x0100,
	}
}

// readPc returns current value of pc and increments it.
func (cpu *cpu) ReadPc() uint16 {
	pc := cpu.pc
	cpu.pc++

	return pc
}

// C returns boolean that indicates whether the C flag is turned on or not.
func (cpu *cpu) C() bool {
	return (cpu.af.Lo() >> 4 & 1) == 1
}

// setC on/off C flag.
func (cpu *cpu) SetC(on bool) {
	cpu.af.SetLo(bit.Toggle(cpu.af.Lo(), 4, on))
}

// H returns boolean that indicates whether the H flag is turned on or not.
func (cpu *cpu) H() bool {
	return (cpu.af.Lo() >> 5 & 1) == 1
}

// setH on/off H flag.
func (cpu *cpu) SetH(on bool) {
	cpu.af.SetLo(bit.Toggle(cpu.af.Lo(), 5, on))
}

// N returns boolean that indicates whether the N flag is turned on or not.
func (cpu *cpu) N() bool {
	return (cpu.af.Lo() >> 6 & 1) == 1
}

// setN on/off N flag.
func (cpu *cpu) SetN(on bool) {
	cpu.af.SetLo(bit.Toggle(cpu.af.Lo(), 6, on))
}

// Z returns boolean that indicates whether the Z flag is turned on or not.
func (cpu *cpu) Z() bool {
	return (cpu.af.Lo() >> 7 & 1) == 1
}

// setZ on/off Z flag.
func (cpu *cpu) SetZ(on bool) {
	cpu.af.SetLo(bit.Toggle(cpu.af.Lo(), 7, on))
}
