package gameboy

import "github.com/elmarsan/gbgo/pkg/bit"

// Joypad represents Gameboy input controller.
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

// joyp represents Joypad register
const joyp = 0xff00

// Get returns uint8 value indicating which keys are currently pressed.
func (j *Joypad) Get() uint8 {
	var state uint8 = 0xcf

	if j.selectActionBtn {
		if j.Start {
			state = bit.Clear(state, 3)
		}

		if j.Select {
			state = bit.Clear(state, 2)
		}

		if j.B {
			state = bit.Clear(state, 1)
		}

		if j.A {
			state = bit.Clear(state, 0)
		}
	}

	if j.selectDirectionBtn {
		if j.Down {
			state = bit.Clear(state, 3)
		}

		if j.Up {
			state = bit.Clear(state, 2)
		}

		if j.Left {
			state = bit.Clear(state, 1)
		}

		if j.Right {
			state = bit.Clear(state, 0)
		}
	}

	return state
}

// Set indicates whether an action or direction key is pressed.
func (j *Joypad) Set(val uint8) {
	j.selectActionBtn = !bit.IsSet(val, 5)
	j.selectDirectionBtn = !bit.IsSet(val, 4)
}
