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
	keys        []ebiten.Key
	gameboyBtns []uint8
}

func NewApp() *App {
	return &App{
		keys: []ebiten.Key{
			ebiten.KeyW,
			ebiten.KeyS,
			ebiten.KeyA,
			ebiten.KeyD,
			ebiten.KeyK,
			ebiten.KeyL,
			ebiten.KeyI,
			ebiten.KeyO,
		},
		gameboyBtns: []uint8{
			BTN_UP,
			BTN_DOWN,
			BTN_LEFT,
			BTN_RIGHT,
			BTN_A,
			BTN_B,
			BTN_START,
			BTN_SELECT,
		},
	}
}

func (a *App) Run() error {
	ebiten.SetWindowSize(WIN_X, WIN_H)
	ebiten.SetWindowTitle("GBGO")

	return ebiten.RunGame(a)
}

func (a *App) Update() error {
	for i, key := range a.keys {
		if ebiten.IsKeyPressed(key) {
			joypad.pressBtn(a.gameboyBtns[i])
		}
	}

	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	imgRGBA := a.createImgRGBA(ppu.videoBuf[:])
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
