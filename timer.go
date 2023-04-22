package main

type Timer struct{}

const (
	// Divider register
	DIV = 0xff04
	// Timer counter
	TIMA = 0xff05
	// Timer module
	TMA = 0xff06
	// Timer control
	TAC = 0xff07
)

func (t *Timer) incDiv(cycles int) {
	div := memory.read(DIV)
	inc := int(div) + cycles

	if inc > 0xff {
		memory.io[DIV-IO_START] = 1
	} else {
		memory.io[DIV-IO_START] = uint8(inc)
	}
}

func (t *Timer) incTIMA() {
	inc := uint16(memory.read(TIMA)) + 1

	if inc > 0xff {
		gameboy.reqInterrupt(IT_TIMER)
		tma := memory.read(TMA)
		memory.write(TIMA, tma)
	} else {
		memory.write(TIMA, uint8(inc))
	}
}

func (t *Timer) resetDIV() {
	memory.write(DIV, 0)
}

func (t *Timer) tacEnabled() bool {
	tac := memory.read(TAC)
	return isBitSet(tac, 2)
}

func (t *Timer) clockFreq() int {
	tac := memory.read(TAC)

	clock := tac & 0x03
	switch clock {
	case 0:
		return 1024
	case 1:
		return 16
	case 2:
		return 64
	}

	return 256
}

func (t *Timer) update(cycles int) {
	t.incDiv(cycles)

	if t.tacEnabled() {
		timerClockSum := int(cycles)
		clockSpeed := t.clockFreq()

		for timerClockSum >= clockSpeed {
			t.incTIMA()
			timerClockSum -= clockSpeed
		}
	}
}
