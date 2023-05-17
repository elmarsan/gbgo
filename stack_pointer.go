package main

// pushSp pushes a register on top of the stack pointer.
func (gb *Gameboy) pushSp(val uint16) {
	gb.memoryBus.write(gb.cpu.sp-1, hi(val))
	gb.memoryBus.write(gb.cpu.sp-2, lo(val))
	gb.cpu.sp -= 2
}

// popSp pops cpu.bus address from top of the stack pointer.
// It reads the value of the address and stores in a register.
func (gb *Gameboy) popSp() uint16 {
	lsb := gb.memoryBus.read(gb.cpu.sp)
	msb := gb.memoryBus.read(gb.cpu.sp + 1)

	gb.cpu.sp += 2
	return joinu8(msb, lsb)
}
