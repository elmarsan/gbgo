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

	// sp represents stack pointer register.
	sp uint16

	// pc represents program counter register.
	pc uint16

	// halted indicates if cpu is in low-power state.
	halted bool

	// ime keeps a bool indicating whether interrupts are enabled or not.
	ime bool

	// enablingIme is set as previous step before enable/disable IME.
	enablingIme bool

	// clockCycles holds clock cycles elapsed during previous instruction execution.
	clockCycles int

	// bus represents memory bus used by Gameboy.
	bus *MemoryBus
}

// NewCPU creates and returns a new CPU instance.
// Register values are set to defaults of DMG boot sequence.
func NewCPU(bus *MemoryBus) *CPU {
	return &CPU{
		a:   0x01,
		f:   0xb0,
		b:   0x00,
		c:   0x13,
		d:   0x00,
		e:   0xd8,
		h:   0x01,
		l:   0x4d,
		sp:  0xfffe,
		pc:  0x0100,
		bus: bus,
	}
}

// execute fetchs next opcode and executes the corresponding instruction.
func (cpu *CPU) execute() {
	// debug.logState()

	cpu.clockCycles = 0
	pc := cpu.readPc()
	opcode := cpu.bus.read(pc)

	if opcode == 0xcb {
		pc := cpu.readPc()
		opcode := cpu.bus.read(pc)
		prefixedInstructions[opcode]()
		cpu.clockCycles += cbInstructionCycles[opcode] * 4
	} else {
		instructions[opcode]()
		cpu.clockCycles += instructionCycles[opcode] * 4
	}
}

// readPc returns current value of pc and increments it.
func (cpu *CPU) readPc() uint16 {
	pc := cpu.pc
	cpu.pc++

	return pc
}

// load8Reg loads b into a.
func (cpu *CPU) load8Reg(a CPU8Register, b CPU8Register) {
	var val uint8 = cpu.read8Reg(b)
	cpu.set8Reg(a, val)
}

// load16RegD16 loads D16 into a.
func (cpu *CPU) load16RegD16(a CPU16Register, d16 uint16) {
	cpu.set16Reg(a, d16)
}

// inc8Reg increments 8 bit register.
// It stores in a register a (a + 1) and sets flags.
func (cpu *CPU) inc8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	inc := reg + 1
	cpu.set8Reg(a, inc)

	cpu.setN(false)
	cpu.setH((reg&0xf)+1 > 0xf)
	cpu.setZ(inc == 0)
}

// dec8Reg decrements 8 bit register.
// It stores in a register a (a - 1) and sets flags.
func (cpu *CPU) dec8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	dec := reg - 1
	cpu.set8Reg(a, dec)

	cpu.setN(true)
	cpu.setH((reg & 0xf) == 0)
	cpu.setZ(dec == 0)
}

// inc16reg increments 16 bit register.
// It stores in a register a (a + 1).
func (cpu *CPU) inc16reg(reg CPU16Register) {
	val := cpu.read16Reg(reg)
	cpu.set16Reg(reg, val+1)
}

// dec16Reg decrements 16 bit register.
// It stores in a register a (a - 1) and sets flags.
func (cpu *CPU) dec16Reg(a CPU16Register) {
	reg := cpu.read16Reg(a)
	cpu.set16Reg(a, reg-1)
}

// add8Reg adds val to register a.
// It stores in a register a (a + val) and sets flags.
func (cpu *CPU) add8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	add := reg + val
	cpu.set8Reg(a, add)

	cpu.setC((uint16(reg) + uint16(val)) > 0xff)
	cpu.setN(false)
	cpu.setH((val&0xF)+(reg&0xF) > 0xF)
	cpu.setZ(add == 0)
}

// add16Reg add val to register a.
// It stores in a register a (a + val) and sets flags.
func (cpu *CPU) add16Reg(a CPU16Register, val uint16) {
	reg := cpu.read16Reg(a)
	add := int32(reg) + int32(val)
	cpu.set16Reg(a, uint16(add))

	cpu.setC(add > 0xffff)
	cpu.setN(false)
	cpu.setH(int32(reg&0xfff) > (add & 0xfff))
}

// adc8Reg add register a, val and carry flag.
// It stores in a register a (a + val + carry flag) and sets flags.
func (cpu *CPU) adc8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	var carry uint8 = 1
	if !cpu.C() {
		carry = 0
	}

	add := reg + val + carry
	cpu.set8Reg(a, add)

	cpu.setC(uint16(reg)+uint16(val)+uint16(carry) > 0xff)
	cpu.setN(false)
	cpu.setH(((reg & 0x0f) + (val & 0x0f) + carry) > 0x0f)
	cpu.setZ(add == 0)
}

// sub8Reg subtracts register a and val.
// It stores in register a (a - val) and sets flags.
func (cpu *CPU) sub8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	sub := reg - val
	cpu.set8Reg(a, sub)

	cpu.setC((int16(reg) - int16(val)) < 0)
	cpu.setN(true)
	cpu.setH((int16(reg&0xf) - int16(val&0xf)) < 0)
	cpu.setZ(sub == 0)
}

// sbc8Reg subtracts register a, val and carry flag.
// It stores in register a (a - val - carry) and sets flags.
func (cpu *CPU) sbc8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	var c uint8 = 1
	if !cpu.C() {
		c = 0
	}
	sub := int16(reg) - int16(val) - int16(c)
	cpu.set8Reg(a, uint8(sub))

	cpu.setC(sub < 0)
	cpu.setN(true)
	cpu.setH(int16(reg&0xf)-int16(val&0xf)-int16(c) < 0)
	cpu.setZ(uint8(sub) == 0)
}

// and8Reg performs bitwise AND between register  and val.
// It stores in register a (a & val) and set flags.
func (cpu *CPU) and8Reg(a CPU8Register, val uint8) {
	and := cpu.read8Reg(a) & val
	cpu.set8Reg(a, and)

	cpu.setC(false)
	cpu.setN(false)
	cpu.setH(true)
	cpu.setZ(and == 0)
}

// xor8Reg performs bitwise XOR between register a and val.
// It stores in register a (a ^ val) and set flags.
func (cpu *CPU) xor8Reg(a CPU8Register, val uint8) {
	xor := cpu.read8Reg(a) ^ val
	cpu.set8Reg(a, xor)

	cpu.setC(false)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setZ(xor == 0)
}

// or8RegD8 performs bitwise OR between register a and val.
// It stores in register a (a | val) and set flags.
func (cpu *CPU) or8Reg(a CPU8Register, val uint8) {
	or := cpu.read8Reg(a) | val
	cpu.set8Reg(a, or)

	cpu.setC(false)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setZ(or == 0)
}

// cp8Reg compares the values of register a and val.
// It set flags.
func (cpu *CPU) cp8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	sub := val - reg

	cpu.setC((int16(reg) - int16(val)) < 0)
	cpu.setN(true)
	cpu.setH((int16(reg&0xf) - int16(val&0xf)) < 0)
	cpu.setZ(sub == 0)
}

// rlca8Reg rotate A left
// It rotates a register 1 bit to the left and set carry flag.
func (cpu *CPU) rlca8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	rotation := rotateLeft(reg, 1)
	cpu.set8Reg(a, rotation)

	cpu.setC(reg > 0x7f)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setZ(false)
}

// rla8Reg rotate A left through carry
// It rotates a register 1 bit to the left and set carry flag.
// The bit rotated is replaced by carry flag value.
func (cpu *CPU) rla8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	var c uint8 = 1
	if !cpu.C() {
		c = 0
	}

	rotation := (reg << 1) + c
	cpu.set8Reg(a, rotation)
	cpu.setC(reg > 0x7f)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setZ(false)
}

// rrca8Reg rotate A RIGHT
// It rotates a register 1 bit to the right and set carry flag.
func (cpu *CPU) rrca8Reg(a CPU8Register) {
	val := cpu.read8Reg(a)
	rotation := rotateRight(val, 1)
	cpu.set8Reg(a, rotation)

	cpu.setC(rotation > 0x7f)
	cpu.setN(false)
	cpu.setH(false)
	cpu.setZ(false)
}

// pushSp pushes a register on top of the stack pointer.
func (cpu *CPU) pushSp(a CPU16Register) {
	val := cpu.read16Reg(a)
	cpu.bus.write(cpu.sp-1, hi(val))
	cpu.bus.write(cpu.sp-2, lo(val))
	cpu.sp -= 2
}

// popSp pops cpu.bus address from top of the stack pointer.
// It reads the value of the address and stores in a register.
func (cpu *CPU) popSp(a CPU16Register) {
	lsb := cpu.bus.read(cpu.sp)
	msb := cpu.bus.read(cpu.sp + 1)
	val := joinu8(msb, lsb)

	cpu.set16Reg(a, val)
	cpu.sp += 2
}

// jump jumps to the next instruction located in addr.
// It sets pc = addr
func (cpu *CPU) jump(addr uint16) {
	cpu.pc = addr
}

// call calls function located in addr and push pc into sp.
func (cpu *CPU) call(addr uint16) {
	cpu.pushSp(REG_PC)
	cpu.pc = addr
}

// ret returns from function.
func (cpu *CPU) ret() {
	cpu.popSp(REG_PC)
}

// C returns boolean that indicates whether the C flag is turned on or not.
func (cpu *CPU) C() bool {
	return (cpu.f >> 4 & 1) == 1
}

// setC on/off C flag.
func (cpu *CPU) setC(on bool) {
	cpu.f = toggleBit(cpu.f, 4, on)
}

// H returns boolean that indicates whether the H flag is turned on or not.
func (cpu *CPU) H() bool {
	return (cpu.f >> 5 & 1) == 1
}

// setH on/off H flag.
func (cpu *CPU) setH(on bool) {
	cpu.f = toggleBit(cpu.f, 5, on)
}

// N returns boolean that indicates whether the N flag is turned on or not.
func (cpu *CPU) N() bool {
	return (cpu.f >> 6 & 1) == 1
}

// setN on/off N flag.
func (cpu *CPU) setN(on bool) {
	cpu.f = toggleBit(cpu.f, 6, on)
}

// Z returns boolean that indicates whether the Z flag is turned on or not.
func (cpu *CPU) Z() bool {
	return (cpu.f >> 7 & 1) == 1
}

// setZ on/off Z flag.
func (cpu *CPU) setZ(on bool) {
	cpu.f = toggleBit(cpu.f, 7, on)
}

// bit8Reg performs BIT instruction in register a.
// It set flags depending on result.
func (cpu *CPU) bit8Reg(a CPU8Register, bit uint8) {
	val := readBit(cpu.read8Reg(a), bit)

	cpu.setH(true)
	cpu.setN(false)
	cpu.setZ(!isBitSet(val, bit))
}

// bitHL performs BIT instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) bitHL(bit uint8) {
	hl := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(hl)

	cpu.setH(true)
	cpu.setN(false)
	cpu.setZ(!isBitSet(val, bit))
}

// swap8Reg performs SWAP instruction in register a.
// It set flags depending on result.
func (cpu *CPU) swap8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	swap := swapNibbleU8(reg)
	cpu.set8Reg(a, swap)

	cpu.setC(false)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(swap == 0)

}

// swapHL performs SWAP instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) swapHL() {
	addr := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(addr)
	swap := swapNibbleU8(val)
	cpu.bus.write(addr, swap)

	cpu.setC(false)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(swap == 0)

}

// srl8Reg performs SRL instruction in register a.
// It set flags depending on result.
func (cpu *CPU) srl8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	carry := readBit(reg, 0)
	shift := reg >> 1
	cpu.set8Reg(a, shift)

	cpu.setC(carry == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(shift == 0)
}

// srlHL performs SRL instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) srlHL() {
	addr := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(addr)
	carry := readBit(val, 0)
	shift := val >> 1
	cpu.bus.write(addr, shift)

	cpu.setC(carry == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(shift == 0)
}

// sra8Reg performs SRA instruction in register a.
// It set flags depending on result.
func (cpu *CPU) sra8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	rotation := (reg & 128) | (reg >> 1)
	cpu.set8Reg(a, rotation)

	cpu.setC((reg & 1) == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// sraHL performs SRA instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) sraHL() {
	addr := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(addr)
	rotation := (val & 128) | (val >> 1)
	cpu.bus.write(addr, rotation)

	cpu.setC((val & 1) == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// sla8Reg performs SLA instruction in register a.
// It set flags depending on result.
func (cpu *CPU) sla8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	carry := reg >> 7
	rotation := (reg << 1) & 0xff
	cpu.set8Reg(a, rotation)

	cpu.setC(carry == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// slaHL performs SLA instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) slaHL() {
	addr := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(addr)
	rotation := (val << 1) & 0xff
	cpu.bus.write(addr, rotation)

	cpu.setC(isBitSet(val, 7))
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// rl8Reg performs RL instruction in register a.
// It set flags depending on result.
func (cpu *CPU) rl8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	var c uint8 = 1
	if !cpu.C() {
		c = 0
	}
	shift := (reg<<1)&0xff | c
	cpu.set8Reg(a, shift)

	cpu.setC(readBit(reg, 7) > 0)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(shift == 0)
}

// rlHL performs RL instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) rlHL() {
	addr := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(addr)
	var c uint8 = 1
	if !cpu.C() {
		c = 0
	}

	rotation := (val<<1)&0xff | c
	cpu.bus.write(addr, rotation)

	cpu.setC(isBitSet(val, 7))
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// rr8Reg performs RR instruction in register a.
// It set flags depending on result.
func (cpu *CPU) rr8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	rotation := rotateRight(reg, 1)
	rotation = toggleBit(rotation, 7, cpu.C())
	cpu.set8Reg(a, rotation)

	cpu.setC(isBitSet(reg, 0))
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// rrHL performs RR instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) rrHL() {
	addr := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(addr)
	var c uint8 = 1
	if !cpu.C() {
		c = 0
	}

	rotation := (val >> 1) | (c << 7)
	cpu.bus.write(addr, rotation)

	cpu.setC((val & 1) == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// rlc8Reg performs RLC instruction in register a.
// It set flags depending on result.
func (cpu *CPU) rlc8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	carry := reg >> 7
	rotation := (reg<<1)&0xff | carry
	cpu.set8Reg(a, rotation)

	cpu.setC(carry == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// rlcHL performs RLC instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) rlcHL() {
	addr := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(addr)
	carry := val >> 7
	rotation := (val<<1)&0xff | carry
	cpu.bus.write(addr, rotation)

	cpu.setC(carry == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// rrc8Reg performs RRC instruction in register a.
// It set flags depending on result.
func (cpu *CPU) rrc8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	carry := reg & 1
	rotation := (reg >> 1) | (carry << 7)
	cpu.set8Reg(a, rotation)

	cpu.setC(carry == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}

// rrcHL performs RRC instruction in register HL.
// It set flags depending on result.
func (cpu *CPU) rrcHL() {
	addr := cpu.read16Reg(REG_HL)
	val := cpu.bus.read(addr)
	carry := val & 1
	rotation := val>>1 | (carry << 7)
	cpu.bus.write(addr, rotation)

	cpu.setC(carry == 1)
	cpu.setH(false)
	cpu.setN(false)
	cpu.setZ(rotation == 0)
}
