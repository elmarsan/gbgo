package main

// CpuRegister represent 16 bits cpu register.
type CpuRegister struct {
	// val holds register value.
	val uint16
}

// Get returns register value.
func (r *CpuRegister) Get() uint16 {
	return r.val
}

// Set modifies register value.
func (r *CpuRegister) Set(val uint16) {
	r.val = val
}

// Hi returns higher 8 bits of register value.
func (r *CpuRegister) Hi() uint8 {
	return uint8(r.val >> 8)
}

// SetHi modifies higher 8 bits of register value.
func (r *CpuRegister) SetHi(val uint8) {
	r.val = uint16(val)<<8 | (uint16(r.val) & 0xff)
}

// Lo returns lower 8 bits of register value.
func (r *CpuRegister) Lo() uint8 {
	return uint8(r.val & 0xff)
}

// SetLo modifies lower 8 bits of register value.
func (r *CpuRegister) SetLo(val uint8) {
	r.val = uint16(val) | (uint16(r.val) & 0xff00)
}

// CPU represents gameboy central processing unit.
type CPU struct {
	af CpuRegister
	bc CpuRegister
	de CpuRegister
	hl CpuRegister

	// sp represents stack pointer register.
	sp uint16

	// pc represents program counter register.
	pc uint16

	// halted indicates if cpu is in low-power state.
	halted bool

	// clockCycles holds clock cycles elapsed during previous instruction execution.
	clockCycles int
}

// NewCPU creates and returns a new CPU instance.
// Register values are set to defaults of DMG boot sequence.
func NewCPU(bus *MemoryBus) *CPU {
	return &CPU{
		af: CpuRegister{val: 0x01b0},
		bc: CpuRegister{val: 0x0013},
		de: CpuRegister{val: 0x00d8},
		hl: CpuRegister{val: 0x014d},
		sp: 0xfffe,
		pc: 0x0100,
	}
}

// readPc returns current value of pc and increments it.
func (cpu *CPU) readPc() uint16 {
	pc := cpu.pc
	cpu.pc++

	return pc
}

// C returns boolean that indicates whether the C flag is turned on or not.
func (cpu *CPU) C() bool {
	return (cpu.af.Lo() >> 4 & 1) == 1
}

// setC on/off C flag.
func (cpu *CPU) setC(on bool) {
	cpu.af.SetLo(toggleBit(cpu.af.Lo(), 4, on))
}

// H returns boolean that indicates whether the H flag is turned on or not.
func (cpu *CPU) H() bool {
	return (cpu.af.Lo() >> 5 & 1) == 1
}

// setH on/off H flag.
func (cpu *CPU) setH(on bool) {
	cpu.af.SetLo(toggleBit(cpu.af.Lo(), 5, on))
}

// N returns boolean that indicates whether the N flag is turned on or not.
func (cpu *CPU) N() bool {
	return (cpu.af.Lo() >> 6 & 1) == 1
}

// setN on/off N flag.
func (cpu *CPU) setN(on bool) {
	cpu.af.SetLo(toggleBit(cpu.af.Lo(), 6, on))
}

// Z returns boolean that indicates whether the Z flag is turned on or not.
func (cpu *CPU) Z() bool {
	return (cpu.af.Lo() >> 7 & 1) == 1
}

// setZ on/off Z flag.
func (cpu *CPU) setZ(on bool) {
	cpu.af.SetLo(toggleBit(cpu.af.Lo(), 7, on))
}
