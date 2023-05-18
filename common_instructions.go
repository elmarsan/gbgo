package main

// inc increments 8 bit register.
// It stores in a register a (a + 1) and sets flags.
func inc(gb *Gameboy, val uint8) uint8 {
	inc := val + 1

	gb.cpu.setN(false)
	gb.cpu.setH((val&0xf)+1 > 0xf)
	gb.cpu.setZ(inc == 0)

	return inc
}

// dec decrements 8 bit register.
// It stores in a register a (a - 1) and sets flags.
func dec(gb *Gameboy, val uint8) uint8 {
	dec := val - 1

	gb.cpu.setN(true)
	gb.cpu.setH((val & 0xf) == 0)
	gb.cpu.setZ(dec == 0)

	return dec
}

// and performs bitwise AND between register  and val.
// It stores in register a (a & val) and set flags.
func and(gb *Gameboy, val uint8) {
	and := gb.cpu.af.Hi() & val
	gb.cpu.af.SetHi(and)

	gb.cpu.setC(false)
	gb.cpu.setN(false)
	gb.cpu.setH(true)
	gb.cpu.setZ(and == 0)
}

// xor performs bitwise XOR between register a and val.
// It stores in register a (a ^ val) and set flags.
func xor(gb *Gameboy, val uint8) {
	xor := gb.cpu.af.Hi() ^ val
	gb.cpu.af.SetHi(xor)

	gb.cpu.setC(false)
	gb.cpu.setN(false)
	gb.cpu.setH(false)
	gb.cpu.setZ(xor == 0)
}

// or performs bitwise OR between register a and val.
// It stores in register a (a | val) and set flags.
func or(gb *Gameboy, val uint8) {
	or := gb.cpu.af.Hi() | val
	gb.cpu.af.SetHi(or)

	gb.cpu.setC(false)
	gb.cpu.setN(false)
	gb.cpu.setH(false)
	gb.cpu.setZ(or == 0)
}

// cp compares the values of register a and val.
// It set flags.
func cp(gb *Gameboy, val uint8) {
	reg := gb.cpu.af.Hi()
	sub := val - reg

	gb.cpu.setC((int16(reg) - int16(val)) < 0)
	gb.cpu.setN(true)
	gb.cpu.setH((int16(reg&0xf) - int16(val&0xf)) < 0)
	gb.cpu.setZ(sub == 0)
}

// jump jumps to the next instruction located in addr.
// It sets pc = addr
func jump(gb *Gameboy, addr uint16) {
	gb.cpu.pc = addr
}

// call calls function located in addr and push pc into sp.
func call(gb *Gameboy, addr uint16) {
	gb.pushSp(gb.cpu.pc)
	gb.cpu.pc = addr
}

// ret returns from function.
func ret(gb *Gameboy) {
	gb.cpu.pc = gb.popSp()
}

// rrca rotate A RIGHT
// It rotates a register 1 bit to the right and set carry flag.
func rrca(gb *Gameboy) {
	val := gb.cpu.af.Hi()
	rotation := rotateRight(val, 1)
	gb.cpu.af.SetHi(rotation)

	gb.cpu.setC(rotation > 0x7f)
	gb.cpu.setN(false)
	gb.cpu.setH(false)
	gb.cpu.setZ(false)
}

// rla rotate A left through carry
// It rotates a register 1 bit to the left and set carry flag.
// The bit rotated is replaced by carry flag value.
func rla(gb *Gameboy) {
	reg := gb.cpu.af.Hi()
	var c uint8 = 1
	if !gb.cpu.C() {
		c = 0
	}

	rot := (reg << 1) + c
	gb.cpu.af.SetHi(rot)

	gb.cpu.setC(reg > 0x7f)
	gb.cpu.setN(false)
	gb.cpu.setH(false)
	gb.cpu.setZ(false)
}

// rlca rotate A left
// It rotates a register 1 bit to the left and set carry flag.
func rlca(gb *Gameboy) {
	reg := gb.cpu.af.Hi()
	rot := rotateLeft(reg, 1)
	gb.cpu.af.SetHi(rot)

	gb.cpu.setC(reg > 0x7f)
	gb.cpu.setN(false)
	gb.cpu.setH(false)
	gb.cpu.setZ(false)
}

// sub subtracts register a and val.
// It stores in register a (a - val) and sets flags.
func sub(gb *Gameboy, reg uint8, val uint8) uint8 {
	sub := reg - val

	gb.cpu.setC((int16(reg) - int16(val)) < 0)
	gb.cpu.setN(true)
	gb.cpu.setH((int16(reg&0xf) - int16(val&0xf)) < 0)
	gb.cpu.setZ(sub == 0)

	return sub
}

// sbc subtracts register a, val and carry flag.
// It stores in register a (a - val - carry) and sets flags.
func sbc(gb *Gameboy, reg uint8, val uint8) uint8 {
	var c uint8 = 1
	if !gb.cpu.C() {
		c = 0
	}
	sub := int16(reg) - int16(val) - int16(c)
	gb.cpu.af.SetHi(uint8(sub))

	gb.cpu.setC(sub < 0)
	gb.cpu.setN(true)
	gb.cpu.setH(int16(reg&0xf)-int16(val&0xf)-int16(c) < 0)
	gb.cpu.setZ(uint8(sub) == 0)

	return uint8(sub)
}

// bit performs BIT instruction in register a.
// It set flags depending on result.
func bit(gb *Gameboy, val uint8, pos uint8) {
	bit := readBit(val, pos)

	gb.cpu.setH(true)
	gb.cpu.setN(false)
	gb.cpu.setZ(!isBitSet(bit, pos))
}

// swap performs SWAP instruction in register a.
// It set flags depending on result.
func swap(gb *Gameboy, val uint8) uint8 {
	swap := swapNibbleU8(val)

	gb.cpu.setC(false)
	gb.cpu.setH(false)
	gb.cpu.setN(false)
	gb.cpu.setZ(swap == 0)

	return swap
}

// srl performs SRL instruction in register a.
// It set flags depending on result.
func srl(gb *Gameboy, val uint8) uint8 {
	carry := readBit(val, 0)
	rot := val >> 1

	gb.cpu.setC(carry == 1)
	gb.cpu.setH(false)
	gb.cpu.setN(false)
	gb.cpu.setZ(rot == 0)

	return rot
}

// sra performs SRA instruction in register a.
// It set flags depending on result.
func sra(gb *Gameboy, val uint8) uint8 {
	rot := (val & 128) | (val >> 1)

	gb.cpu.setC((val & 1) == 1)
	gb.cpu.setH(false)
	gb.cpu.setN(false)
	gb.cpu.setZ(rot == 0)

	return rot
}

// sla performs SLA instruction in register a.
// It set flags depending on result.
func sla(gb *Gameboy, val uint8) uint8 {
	carry := val >> 7
	rot := (val << 1) & 0xff

	gb.cpu.setC(carry == 1)
	gb.cpu.setH(false)
	gb.cpu.setN(false)
	gb.cpu.setZ(rot == 0)

	return rot
}

// rl performs RL instruction in register a.
// It set flags depending on result.
func rl(gb *Gameboy, val uint8) uint8 {
	var c uint8 = 1
	if !gb.cpu.C() {
		c = 0
	}
	rot := (val<<1)&0xff | c

	gb.cpu.setC(readBit(val, 7) > 0)
	gb.cpu.setH(false)
	gb.cpu.setN(false)
	gb.cpu.setZ(rot == 0)
	return rot
}

// rr performs RR instruction in register a.
// It set flags depending on result.
func rr(gb *Gameboy, val uint8) uint8 {
	rot := rotateRight(val, 1)
	rot = toggleBit(rot, 7, gb.cpu.C())

	gb.cpu.setC(isBitSet(val, 0))
	gb.cpu.setH(false)
	gb.cpu.setN(false)
	gb.cpu.setZ(rot == 0)

	return rot
}

// rlc performs RLC instruction in register a.
// It set flags depending on result.
func rlc(gb *Gameboy, val uint8) uint8 {
	carry := val >> 7
	rot := (val<<1)&0xff | carry

	gb.cpu.setC(carry == 1)
	gb.cpu.setH(false)
	gb.cpu.setN(false)
	gb.cpu.setZ(rot == 0)

	return rot
}

// rrc performs RRC instruction in register a.
// It set flags depending on result.
func rrc(gb *Gameboy, val uint8) uint8 {
	carry := val & 1
	rot := (val >> 1) | (carry << 7)

	gb.cpu.setC(carry == 1)
	gb.cpu.setH(false)
	gb.cpu.setN(false)
	gb.cpu.setZ(rot == 0)

	return rot
}

// add adds val to register a.
// It stores in a register a (a + val) and sets flags.
func add(gb *Gameboy, reg uint8, val uint8) uint8 {
	add := reg + val

	gb.cpu.setC((uint16(reg) + uint16(val)) > 0xff)
	gb.cpu.setN(false)
	gb.cpu.setH((val&0xF)+(reg&0xF) > 0xF)
	gb.cpu.setZ(add == 0)

	return add
}

// add16 add val to register a.
// It stores in a register a (a + val) and sets flags.
func add16(gb *Gameboy, reg uint16, val uint16) uint16 {
	add := int32(reg) + int32(val)

	gb.cpu.setC(add > 0xffff)
	gb.cpu.setN(false)
	gb.cpu.setH(int32(reg&0xfff) > (add & 0xfff))

	return uint16(add)
}

// adc add register a, val and carry flag.
// It stores in a register a (a + val + carry flag) and sets flags.
func adc(gb *Gameboy, reg uint8, val uint8) uint8 {
	var carry uint8 = 1
	if !gb.cpu.C() {
		carry = 0
	}

	add := reg + val + carry

	gb.cpu.setC(uint16(reg)+uint16(val)+uint16(carry) > 0xff)
	gb.cpu.setN(false)
	gb.cpu.setH(((reg & 0x0f) + (val & 0x0f) + carry) > 0x0f)
	gb.cpu.setZ(add == 0)

	return add
}
