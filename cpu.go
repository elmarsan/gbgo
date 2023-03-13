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

	ime bool  // Interrup master enable flag
	ie  uint8 // Interrup enable
	iF  uint8 // Interrup flag

	running bool
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
	REG_PC
	REG_SP
)

type CPUFlag int

const (
	C CPUFlag = iota // bit 4
	H                // bit 5
	N                // bit 6
	Z                // bit 7
)

var cpu = &CPU{}

func (cpu *CPU) init() {
	cpu.a = 0x01
	cpu.f = 0xb0
	cpu.b = 0x00
	cpu.c = 0x13
	cpu.d = 0x00
	cpu.e = 0xd8
	cpu.h = 0x01
	cpu.l = 0x4d
	cpu.sp = 0xfffe
	cpu.pc = 0x0100
}

// execute executes next instruction.
func (cpu *CPU) execute() {
	logState()
	pc := cpu.readPc()
	opcode := memory.read(pc)
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
		// log.Fatalf("Unknown register %s", reg)
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
	case REG_PC:
		return cpu.pc
	case REG_SP:
		return cpu.sp
	default:
		// log.Fatalf("Unknown register %s", reg)
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
		// log.Fatalf("Unknown register %s", reg)
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
	case REG_SP:
		cpu.sp = val
		break
	case REG_PC:
		cpu.pc = val
		break
	default:
		// log.Fatalf("Unknown register %s", reg)
		break
	}
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

	cpu.setFlag(N, false)
	cpu.setFlag(H, (reg&0xf)+1 > 0xf)
	cpu.setFlag(Z, inc == 0)
}

// dec8Reg decrements 8 bit register.
// It stores in a register a (a - 1) and sets flags.
func (cpu *CPU) dec8Reg(a CPU8Register) {
	reg := cpu.read8Reg(a)
	dec := reg - 1
	cpu.set8Reg(a, dec)

	cpu.setFlag(N, true)
	cpu.setFlag(H, (reg&0xf) == 0)
	cpu.setFlag(Z, dec == 0)
}

// inc16reg increments 16 bit register.
// It stores in a register a (a + 1) and sets flags.
func (cpu *CPU) inc16reg(reg CPU16Register, setFlags bool) {
	val := cpu.read16Reg(reg)
	inc := val + 1
	cpu.set16Reg(reg, inc)

	if setFlags {
		cpu.setFlag(N, false)
		cpu.setFlag(H, (val&0x00ff)+1 > 0xff)
		cpu.setFlag(Z, inc == 0)
	}
}

// dec16Reg decrements 16 bit register.
// It stores in a register a (a - 1) and sets flags.
func (cpu *CPU) dec16Reg(a CPU16Register) {
	reg := cpu.read16Reg(a)
	dec := reg - 1
	cpu.set16Reg(a, dec)

	cpu.setFlag(N, true)
	cpu.setFlag(H, (reg&0x00ff) == 0)
	cpu.setFlag(Z, dec == 0)
}

// add8Reg adds val to register a.
// It stores in a register a (a + val) and sets flags.
func (cpu *CPU) add8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	add := uint16(reg) + uint16(val)
	cpu.set8Reg(a, uint8(add))

	cpu.setFlag(C, add > 0xff)
	cpu.setFlag(N, false)
	cpu.setFlag(H, (val&0xF)+(reg&0xF) > 0xF)
	cpu.setFlag(Z, add == 0)
}

// add16Reg add val to register a.
// It stores in a register a (a + val) and sets flags.
func (cpu *CPU) add16Reg(a CPU16Register, val uint16) {
	reg := cpu.read16Reg(a)
	add := reg + val
	cpu.set16Reg(a, add)

	cpu.setFlag(C, add > 0xffff)
	cpu.setFlag(N, false)
	cpu.setFlag(H, (reg&0xFFF) > (add&0xFFF))
	cpu.setFlag(Z, add == 1)
}

// adc8Reg add register a, val and carry flag.
// It stores in a register a (a + val + carry flag) and sets flags.
func (cpu *CPU) adc8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	carry := readBit(cpu.f, 4)
	add := reg + val + carry
	cpu.set8Reg(a, add)

	cpu.setFlag(C, add > 0xff)
	cpu.setFlag(N, false)
	cpu.setFlag(H, ((reg&0x0f)+(val&0x0f)+carry) > 0x0f)
	cpu.setFlag(Z, add == 0)
}

// sub8Reg subtracts register a and val.
// It stores in register a (a - val) and sets flags.
func (cpu *CPU) sub8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	sub := reg - val
	cpu.set8Reg(a, sub)

	cpu.setFlag(C, sub > 0xff)
	cpu.setFlag(N, true)
	cpu.setFlag(H, (sub&0x0f) == 0)
	cpu.setFlag(Z, sub == 0)
}

// sbc8Reg subtracts register a, val and carry flag.
// It stores in register a (a - val - carry) and sets flags.
func (cpu *CPU) sbc8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	carry := readBit(cpu.f, 4)
	sub := reg - val - carry
	cpu.set8Reg(a, sub)

	cpu.setFlag(C, sub > 0xff)
	cpu.setFlag(N, true)
	cpu.setFlag(H, (sub&0x0f) == 0)
	cpu.setFlag(Z, sub == 0)
}

// and8Reg performs bitwise AND between register  and val.
// It stores in register a (a & val) and set flags.
func (cpu *CPU) and8Reg(a CPU8Register, val uint8) {
	and := cpu.read8Reg(a) & val
	cpu.set8Reg(a, and)

	cpu.setFlag(C, false)
	cpu.setFlag(N, false)
	cpu.setFlag(H, true)
	cpu.setFlag(Z, and == 0)
}

// xor8Reg performs bitwise XOR between register a and val.
// It stores in register a (a ^ val) and set flags.
func (cpu *CPU) xor8Reg(a CPU8Register, val uint8) {
	xor := cpu.read8Reg(a) ^ val
	cpu.set8Reg(a, xor)

	cpu.setFlag(C, false)
	cpu.setFlag(N, false)
	cpu.setFlag(H, false)
	cpu.setFlag(Z, xor == 0)
}

// or8RegD8 performs bitwise OR between register a and val.
// It stores in register a (a | val) and set flags.
func (cpu *CPU) or8Reg(a CPU8Register, val uint8) {
	or := cpu.read8Reg(a) | val
	cpu.set8Reg(a, or)

	cpu.setFlag(C, false)
	cpu.setFlag(N, false)
	cpu.setFlag(H, false)
	cpu.setFlag(Z, or == 0)
}

// cp8Reg compares the values of register a and val.
// It set flags.
func (cpu *CPU) cp8Reg(a CPU8Register, val uint8) {
	reg := cpu.read8Reg(a)
	sub := val - reg

	cpu.setFlag(C, reg > val)
	cpu.setFlag(N, true)
	cpu.setFlag(H, (reg&0x0f) > (val&0x0f))
	cpu.setFlag(Z, sub == 0)
}

// rlc8Reg rotate A left
// It rotates a register 1 bit to the left and set carry flag.
func (cpu *CPU) rlc8Reg(a CPU8Register) {
	val := cpu.read8Reg(a)
	cpu.set8Reg(a, rotateLeft(val, 1))
	// cpu.cFlag = readBit(val, 7) == 1
}

// rl8Reg rotate A left through carry
// It rotates a register 1 bit to the left and set carry flag.
// The bit rotated is replaced by carry flag value.
func (cpu *CPU) rl8Reg(a CPU8Register) {
	val := cpu.read8Reg(a)
	rotated := rotateLeft(val, 1)

	if cpu.readFlag(C) == 1 {
		setBit(rotated, 7)
	} else {
		clearBit(rotated, 7)
	}

	cpu.set8Reg(a, rotated)
	cpu.setFlag(C, readBit(rotated, 7) == 1)
}

// rrc8Reg rotate A RIGHT
// It rotates a register 1 bit to the right and set carry flag.
func (cpu *CPU) rrc8Reg(a CPU8Register) {
	val := cpu.read8Reg(a)
	cpu.set8Reg(a, rotateRight(val, 1))
	cpu.setFlag(C, readBit(val, 0) == 1)
}

// rr8Reg rotate A right through carry
// It rotates a register 1 bit to the right and set carry flag.
// The bit rotated is replaced by carry flag value.
func (cpu *CPU) rr8Reg(a CPU8Register) {
	val := cpu.read8Reg(a)
	rotated := rotateRight(val, 1)

	if cpu.readFlag(C) == 1 {
		setBit(rotated, 0)
	} else {
		clearBit(rotated, 0)
	}

	cpu.set8Reg(a, rotated)
	cpu.setFlag(C, readBit(rotated, 0) == 1)
}

// pushSp pushes a register on top of the stack pointer.
func (cpu *CPU) pushSp(a CPU16Register) {
	val := cpu.read16Reg(a)
	memory.write(cpu.sp-1, hi(val))
	memory.write(cpu.sp-2, lo(val))
	cpu.sp -= 2
}

// popSp pops memory address from top of the stack pointer.
// It reads the value of the address and stores in a register.
func (cpu *CPU) popSp(a CPU16Register) {
	lsb := memory.read(cpu.sp)
	msb := memory.read(cpu.sp + 1)
	val := joinUint8(msb, lsb)

	cpu.set16Reg(a, val)
	cpu.sp += 2
}

// setIME sets IME flag.
func (cpu *CPU) setIME(enabled bool) {
	cpu.ime = enabled
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

func (cpu *CPU) setFlag(f CPUFlag, val bool) {
	var bit uint8 = 0

	switch f {
	case C:
		bit = 4
		break
	case H:
		bit = 5
		break
	case N:
		bit = 6
		break
	case Z:
		bit = 7
		break
	}

	if val {
		cpu.f = setBit(cpu.f, bit)
	} else {
		cpu.f = clearBit(cpu.f, bit)
	}
}

func (cpu *CPU) readFlag(f CPUFlag) uint8 {
	switch f {
	case C:
		return readBit(cpu.f, 4)
	case H:
		return readBit(cpu.f, 5)
	case N:
		return readBit(cpu.f, 6)
	case Z:
		return readBit(cpu.f, 7)
	}

	return 0
}
