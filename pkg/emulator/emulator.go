package emulator

import (
	"image/color"

	"github.com/elmarsan/gbgo/pkg/gameboy"
)

// Btn represents emulator button.
type Btn uint8

const (
	UpBtn Btn = iota
	DownBtn
	LeftBtn
	RightBtn
	SelectBtn
	StartBtn
	ABtn
	BBtn
)

// Emulator represents an emulator of Game boy console.
type Emulator struct {
	// Gb represents Gb console.
	Gb *gameboy.Gameboy
}

func New() *Emulator {
	return &Emulator{
		Gb: gameboy.New(),
	}
}

// SetButton press or release gameboy joypad button.
func (e *Emulator) SetButton(btn Btn, down bool) {
	switch btn {
	case UpBtn:
		e.Gb.Joypad.Up = down
		break
	case DownBtn:
		e.Gb.Joypad.Down = down
		break
	case LeftBtn:
		e.Gb.Joypad.Left = down
		break
	case RightBtn:
		e.Gb.Joypad.Right = down
		break
	case SelectBtn:
		e.Gb.Joypad.Select = down
		break
	case StartBtn:
		e.Gb.Joypad.Start = down
		break
	case ABtn:
		e.Gb.Joypad.A = down
		break
	case BBtn:
		e.Gb.Joypad.B = down
		break
	}
}

// GetRGBAPixels returns an slice of RGBA colors.
// RBGA pixels are generated from ppu video buffer,
// Each videoBuf pixel contains an index of palette color.
func (e *Emulator) GetRGBAPixels() []color.RGBA {
	pixels := [gameboy.ScreenWidth * gameboy.ScreenHeight]color.RGBA{}
	videoBuffer := e.Gb.GetVideoBuffer()

	for y := 0; y < gameboy.ScreenHeight; y++ {
		for x := 0; x < gameboy.ScreenWidth; x++ {
			i := y*gameboy.ScreenWidth + x
			colorIndex := videoBuffer[i]
			rgba := gameboy.OriginalPalette[colorIndex]
			pixels[i] = rgba
		}
	}

	return pixels[:]
}
