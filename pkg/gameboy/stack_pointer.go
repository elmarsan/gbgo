package gameboy

import "github.com/elmarsan/gbgo/pkg/bit"

// pushSp pushes a register on top of the stack pointer.
func (gb *Gameboy) pushSp(val uint16) {
	gb.memoryBus.write(gb.cpu.sp-1, bit.Hi(val))
	gb.memoryBus.write(gb.cpu.sp-2, bit.Lo(val))
	gb.cpu.sp -= 2
}

// popSp pops cpu.bus address from top of the stack pointer.
// It reads the value of the address and stores in a register.func (gb *Gameboy) popSp() uint16 {
func (gb *Gameboy) popSp() uint16 {
	lsb := gb.memoryBus.read(gb.cpu.sp)
	msb := gb.memoryBus.read(gb.cpu.sp + 1)
	gb.cpu.sp += 2
	return bit.Joinu8(msb, lsb)
}
