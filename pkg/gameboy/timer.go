package gameboy

import "github.com/elmarsan/gbgo/pkg/bit"

const (
	// Divider register
	div = 0xff04

	// Timer counter
	tima = 0xff05

	// Timer module
	tma = 0xff06

	// Timer control
	tac = 0xff07
)

// timer represents the timer hardware in the Game Boy console.
type timer struct {
	clockCycles uint

	// memoryBus represents memory memoryBus used by Gameboy.
	memoryBus *memoryBus

	// InterruptBus represents the interruption system of game boy.
	interruptBus *interruptBus
}

// newTimer returns new instance of Game boy timer.
func newTimer(memoryBus *memoryBus, irBus *interruptBus) *timer {
	return &timer{
		memoryBus:    memoryBus,
		interruptBus: irBus,
	}
}

// Tick updates the timer with the given number of clock cycles.
func (t *timer) tick(clockCycles int) {
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
func (t *timer) incDiv(cycles int) {
	divVal := t.memoryBus.read(div)
	inc := int(divVal) + cycles

	if inc > 0xff {
		t.memoryBus.io[div-ioStart] = 1
	} else {
		t.memoryBus.io[div-ioStart] = uint8(inc)
	}
}

// incTIMA increments the timer counter and handles overflow.
func (t *timer) incTIMA() {
	inc := uint16(t.memoryBus.read(tima)) + 1

	if inc > 0xff {
		t.interruptBus.request(timerInterrupt)
		tma := t.memoryBus.read(tma)
		t.memoryBus.write(tima, tma)
	} else {
		t.memoryBus.write(tima, uint8(inc))
	}
}

// ResetDIV resets the divider register to 0.
func (t *timer) resetDIV() {
	t.memoryBus.write(div, 0)
}

// tacEnabled returns whether the timer is enabled.
func (t *timer) tacEnabled() bool {
	tac := t.memoryBus.read(tac)
	return bit.IsSet(tac, 2)
}

// clockFreq returns the clock frequency of the timer.
func (t *timer) clockFreq() int {
	tac := t.memoryBus.read(tac)

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
