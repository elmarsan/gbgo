package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIN_X = 640
	WIN_H = 480
)

type App struct {
	keyHandler map[ebiten.Key]func(pressed bool)
}

func NewApp() *App {
	return &App{
		keyHandler: map[ebiten.Key]func(pressed bool){
			ebiten.KeyW: func(pressed bool) { gb.joypad.Up = pressed },
			ebiten.KeyS: func(pressed bool) { gb.joypad.Down = pressed },
			ebiten.KeyA: func(pressed bool) { gb.joypad.Left = pressed },
			ebiten.KeyD: func(pressed bool) { gb.joypad.Right = pressed },
			ebiten.KeyK: func(pressed bool) { gb.joypad.A = pressed },
			ebiten.KeyL: func(pressed bool) { gb.joypad.B = pressed },
			ebiten.KeyI: func(pressed bool) { gb.joypad.Start = pressed },
			ebiten.KeyO: func(pressed bool) { gb.joypad.Select = pressed },
		},
	}
}

func (a *App) Run() error {
	ebiten.SetWindowSize(WIN_X, WIN_H)
	ebiten.SetWindowTitle("GBGO")

	return ebiten.RunGame(a)
}

func (a *App) Update() error {
	for key, handler := range a.keyHandler {
		handler(ebiten.IsKeyPressed(key))
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	imgRGBA := a.createImgRGBA(gb.ppu.videoBuf[:])
	img := ebiten.NewImageFromImage(imgRGBA)

	drawopts := &ebiten.DrawImageOptions{}
	drawopts.GeoM.Scale(4, 3.33)
	screen.DrawImage(img, drawopts)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (a *App) createImgRGBA(pixels []uint8) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, GB_W, GB_H))

	for y := 0; y < GB_H; y++ {
		for x := 0; x < GB_W; x++ {
			i := y*GB_W + x
			img.SetRGBA(x, y, palette[pixels[i]])
		}
	}

	return img
}
