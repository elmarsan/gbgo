package main

// Joypad represents game boy input controller.
// It holds which keys are down/up.
type Joypad struct {
	Start  bool
	Select bool
	B      bool
	A      bool
	Down   bool
	Up     bool
	Left   bool
	Right  bool

	selectActionBtn    bool
	selectDirectionBtn bool
}

// JOYP represents Joypad register
const JOYP = 0xff00

// Get returns uint8 value indicating which keys are currently pressed.
func (j *Joypad) Get() uint8 {
	var state uint8 = 0xcf

	if j.selectActionBtn {
		if j.Start {
			state = clearBit(state, 3)
		}

		if j.Select {
			state = clearBit(state, 2)
		}

		if j.B {
			state = clearBit(state, 1)
		}

		if j.A {
			state = clearBit(state, 0)
		}
	}

	if j.selectDirectionBtn {
		if j.Down {
			state = clearBit(state, 3)
		}

		if j.Up {
			state = clearBit(state, 2)
		}

		if j.Left {
			state = clearBit(state, 1)
		}

		if j.Right {
			state = clearBit(state, 0)
		}
	}

	return state
}

// Set indicates whether an action or direction key is pressed.
func (j *Joypad) Set(val uint8) {
	j.selectActionBtn = !isBitSet(val, 5)
	j.selectDirectionBtn = !isBitSet(val, 4)
}
