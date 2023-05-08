package main

// Timer represents the timer hardware in the Game Boy console.
type Timer struct {
	clockCycles uint
}

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

// Tick updates the timer with the given number of clock cycles.
func (t *Timer) Tick(clockCycles int) {
	// increment the divider register based on the number of cycles passed
	t.incDiv(clockCycles)

	// if the timer is enabled, increment the timer counter based on the clock frequency
	if t.tacEnabled() {
		t.clockCycles += uint(clockCycles)
		freq := uint(t.clockFreq())

		for t.clockCycles >= freq {
			t.clockCycles -= freq

			t.incTIMA()
		}
	}
}

// incDiv increments the divider register based on the given number of cycles.
func (t *Timer) incDiv(cycles int) {
	div := gb.bus.read(DIV)
	inc := int(div) + cycles

	if inc > 0xff {
		gb.bus.io[DIV-IO_START] = 1
	} else {
		gb.bus.io[DIV-IO_START] = uint8(inc)
	}
}

// incTIMA increments the timer counter and handles overflow.
func (t *Timer) incTIMA() {
	inc := uint16(gb.bus.read(TIMA)) + 1

	if inc > 0xff {
		gb.reqInterrupt(IT_TIMER)
		tma := gb.bus.read(TMA)
		gb.bus.write(TIMA, tma)
	} else {
		gb.bus.write(TIMA, uint8(inc))
	}
}

// resetDIV resets the divider register to 0.
func (t *Timer) resetDIV() {
	gb.bus.write(DIV, 0)
}

// tacEnabled returns whether the timer is enabled.
func (t *Timer) tacEnabled() bool {
	tac := gb.bus.read(TAC)
	return isBitSet(tac, 2)
}

// clockFreq returns the clock frequency of the timer.
func (t *Timer) clockFreq() int {
	tac := gb.bus.read(TAC)

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
