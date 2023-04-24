package main

// Joypad represents game boy input controller.
type Joypad struct{}

// Gameboy buttons
const (
	BTN_A = iota
	BTN_B
	BTN_UP
	BTN_DOWN
	BTN_RIGHT
	BTN_LEFT
	BTN_SELECT
	BTN_START
)

// JOYP represents Joypad register
const JOYP = 0xff00

var joypadBtnMap = map[uint8]uint8{
	BTN_A:      0,
	BTN_RIGHT:  0,
	BTN_B:      1,
	BTN_LEFT:   1,
	BTN_UP:     2,
	BTN_SELECT: 2,
	BTN_DOWN:   3,
	BTN_START:  3,
}

// pressBtn simulates button press action.
func (j *Joypad) pressBtn(btn uint8) {
	joyp := memory.read(JOYP)
	bit, _ := joypadBtnMap[btn]

	memory.write(JOYP, setBit(joyp, bit))
	gameboy.reqInterrupt(IT_JOYPAD)
}

// releaseBtn simulates button release action.
func (j *Joypad) releaseBtn(btn uint8) {
	joyp := memory.read(JOYP)
	bit, _ := joypadBtnMap[btn]
	memory.write(JOYP, clearBit(joyp, bit))
	gameboy.reqInterrupt(IT_JOYPAD)
}

func (j *Joypad) setActionBtns() {
	joyp := memory.read(JOYP)

	if isBitSet(joyp, 5) {
		memory.write(JOYP, clearBit(joyp, 5))
	} else {
		memory.write(JOYP, setBit(joyp, 5))
	}
}
