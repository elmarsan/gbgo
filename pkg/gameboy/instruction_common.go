package gameboy

import "github.com/elmarsan/gbgo/pkg/bit"

// INC increments 8 bit register.
// It stores in a register a (a + 1) and sets flags.
func inc(gb *Gameboy, val uint8) uint8 {
	inc := val + 1

	gb.cpu.SetN(false)
	gb.cpu.SetH((val&0xf)+1 > 0xf)
	gb.cpu.SetZ(inc == 0)

	return inc
}

// DEC decrements 8 bit register.
// It stores in a register a (a - 1) and sets flags.
func dec(gb *Gameboy, val uint8) uint8 {
	dec := val - 1

	gb.cpu.SetN(true)
	gb.cpu.SetH((val & 0xf) == 0)
	gb.cpu.SetZ(dec == 0)

	return dec
}

// ADN performs bitwise AND between register  ADN val.
// It stores in register a (a & val) ADN set flags.
func and(gb *Gameboy, val uint8) {
	and := gb.cpu.af.Hi() & val
	gb.cpu.af.SetHi(and)

	gb.cpu.SetC(false)
	gb.cpu.SetN(false)
	gb.cpu.SetH(true)
	gb.cpu.SetZ(and == 0)
}

// XOR performs bitwise XOR between register a and val.
// It stores in register a (a ^ val) and set flags.
func xor(gb *Gameboy, val uint8) {
	xor := gb.cpu.af.Hi() ^ val
	gb.cpu.af.SetHi(xor)

	gb.cpu.SetC(false)
	gb.cpu.SetN(false)
	gb.cpu.SetH(false)
	gb.cpu.SetZ(xor == 0)
}

// OR performs bitwise OR between register a and val.
// It stores in register a (a | val) and set flags.
func or(gb *Gameboy, val uint8) {
	or := gb.cpu.af.Hi() | val
	gb.cpu.af.SetHi(or)

	gb.cpu.SetC(false)
	gb.cpu.SetN(false)
	gb.cpu.SetH(false)
	gb.cpu.SetZ(or == 0)
}

// CP compares the values of register a and val.
// It set flags.
func cp(gb *Gameboy, val uint8) {
	reg := gb.cpu.af.Hi()
	sub := val - reg

	gb.cpu.SetC((int16(reg) - int16(val)) < 0)
	gb.cpu.SetN(true)
	gb.cpu.SetH((int16(reg&0xf) - int16(val&0xf)) < 0)
	gb.cpu.SetZ(sub == 0)
}

// JUMP jumps to the next instruction located in addr.
// It sets pc = addr
func jump(gb *Gameboy, addr uint16) {
	gb.cpu.pc = addr
}

// CALL calls function located in addr and push pc into sp.
func call(gb *Gameboy, addr uint16) {
	gb.pushSp(gb.cpu.pc)
	gb.cpu.pc = addr
}

// RET returns from function.
func ret(gb *Gameboy) {
	gb.cpu.pc = gb.popSp()
}

// RRCA rotate A RIGHT
// It rotates a register 1 bit to the right and set carry flag.
func rrca(gb *Gameboy) {
	val := gb.cpu.af.Hi()
	rotation := bit.RotateRight(val, 1)
	gb.cpu.af.SetHi(rotation)

	gb.cpu.SetC(rotation > 0x7f)
	gb.cpu.SetN(false)
	gb.cpu.SetH(false)
	gb.cpu.SetZ(false)
}

// RLA rotate A left through carry
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

	gb.cpu.SetC(reg > 0x7f)
	gb.cpu.SetN(false)
	gb.cpu.SetH(false)
	gb.cpu.SetZ(false)
}

// RLCA rotate A left
// It rotates a register 1 bit to the left and set carry flag.
func rlca(gb *Gameboy) {
	reg := gb.cpu.af.Hi()
	rot := bit.RotateLeft(reg, 1)
	gb.cpu.af.SetHi(rot)

	gb.cpu.SetC(reg > 0x7f)
	gb.cpu.SetN(false)
	gb.cpu.SetH(false)
	gb.cpu.SetZ(false)
}

// SUB subtracts register a and val.
// It stores in register a (a - val) and sets flags.
func sub(gb *Gameboy, reg uint8, val uint8) uint8 {
	sub := reg - val

	gb.cpu.SetC((int16(reg) - int16(val)) < 0)
	gb.cpu.SetN(true)
	gb.cpu.SetH((int16(reg&0xf) - int16(val&0xf)) < 0)
	gb.cpu.SetZ(sub == 0)

	return sub
}

// SBC subtracts register a, val and carry flag.
// It stores in register a (a - val - carry) and sets flags.
func sbc(gb *Gameboy, reg uint8, val uint8) uint8 {
	var c uint8 = 1
	if !gb.cpu.C() {
		c = 0
	}
	sub := int16(reg) - int16(val) - int16(c)
	gb.cpu.af.SetHi(uint8(sub))

	gb.cpu.SetC(sub < 0)
	gb.cpu.SetN(true)
	gb.cpu.SetH(int16(reg&0xf)-int16(val&0xf)-int16(c) < 0)
	gb.cpu.SetZ(uint8(sub) == 0)

	return uint8(sub)
}

// bit performs BIT instruction in register a.
// It set flags depending on result.
func bitInstr(gb *Gameboy, val uint8, pos uint8) {
	bitVal := bit.Read(val, pos)

	gb.cpu.SetH(true)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(!bit.IsSet(bitVal, pos))
}

// SWAP performs SWAP instruction in register a.
// It set flags depending on result.
func swap(gb *Gameboy, val uint8) uint8 {
	swap := bit.SwapNibbleU8(val)

	gb.cpu.SetC(false)
	gb.cpu.SetH(false)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(swap == 0)

	return swap
}

// SRL performs SRL instruction in register a.
// It set flags depending on result.
func srl(gb *Gameboy, val uint8) uint8 {
	carry := bit.Read(val, 0)
	rot := val >> 1

	gb.cpu.SetC(carry == 1)
	gb.cpu.SetH(false)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(rot == 0)

	return rot
}

// SRA performs SRA instruction in register a.
// It set flags depending on result.
func sra(gb *Gameboy, val uint8) uint8 {
	rot := (val & 128) | (val >> 1)

	gb.cpu.SetC((val & 1) == 1)
	gb.cpu.SetH(false)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(rot == 0)

	return rot
}

// SLA performs SLA instruction in register a.
// It set flags depending on result.
func sla(gb *Gameboy, val uint8) uint8 {
	carry := val >> 7
	rot := (val << 1) & 0xff

	gb.cpu.SetC(carry == 1)
	gb.cpu.SetH(false)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(rot == 0)

	return rot
}

// RL performs RL instruction in register a.
// It set flags depending on result.
func rl(gb *Gameboy, val uint8) uint8 {
	var c uint8 = 1
	if !gb.cpu.C() {
		c = 0
	}
	rot := (val<<1)&0xff | c

	gb.cpu.SetC(bit.Read(val, 7) > 0)
	gb.cpu.SetH(false)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(rot == 0)
	return rot
}

// RR performs RR instruction in register a.
// It set flags depending on result.
func rr(gb *Gameboy, val uint8) uint8 {
	rot := bit.RotateRight(val, 1)
	rot = bit.Toggle(rot, 7, gb.cpu.C())

	gb.cpu.SetC(bit.IsSet(val, 0))
	gb.cpu.SetH(false)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(rot == 0)

	return rot
}

// RLC performs RLC instruction in register a.
// It set flags depending on result.
func rlc(gb *Gameboy, val uint8) uint8 {
	carry := val >> 7
	rot := (val<<1)&0xff | carry

	gb.cpu.SetC(carry == 1)
	gb.cpu.SetH(false)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(rot == 0)

	return rot
}

// RRC performs RRC instruction in register a.
// It set flags depending on result.
func rrc(gb *Gameboy, val uint8) uint8 {
	carry := val & 1
	rot := (val >> 1) | (carry << 7)

	gb.cpu.SetC(carry == 1)
	gb.cpu.SetH(false)
	gb.cpu.SetN(false)
	gb.cpu.SetZ(rot == 0)

	return rot
}

// ADD adds val to register a.
// It stores in a register a (a + val) and sets flags.
func add(gb *Gameboy, reg uint8, val uint8) uint8 {
	add := reg + val

	gb.cpu.SetC((uint16(reg) + uint16(val)) > 0xff)
	gb.cpu.SetN(false)
	gb.cpu.SetH((val&0xF)+(reg&0xF) > 0xF)
	gb.cpu.SetZ(add == 0)

	return add
}

// ADD16 add val to register a.
// It stores in a register a (a + val) and sets flags.
func add16(gb *Gameboy, reg uint16, val uint16) uint16 {
	add := int32(reg) + int32(val)

	gb.cpu.SetC(add > 0xffff)
	gb.cpu.SetN(false)
	gb.cpu.SetH(int32(reg&0xfff) > (add & 0xfff))

	return uint16(add)
}

// ADC add register a, val and carry flag.
// It stores in a register a (a + val + carry flag) and sets flags.
func adc(gb *Gameboy, reg uint8, val uint8) uint8 {
	var carry uint8 = 1
	if !gb.cpu.C() {
		carry = 0
	}

	add := reg + val + carry

	gb.cpu.SetC(uint16(reg)+uint16(val)+uint16(carry) > 0xff)
	gb.cpu.SetN(false)
	gb.cpu.SetH(((reg & 0x0f) + (val & 0x0f) + carry) > 0x0f)
	gb.cpu.SetZ(add == 0)

	return add
}
