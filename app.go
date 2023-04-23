package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIN_X = 640
	WIN_H = 480
)

type App struct{}

func (a *App) Run() error {
	ebiten.SetWindowSize(WIN_X, WIN_H)
	ebiten.SetWindowTitle("GBGO")

	return ebiten.RunGame(a)
}

func (a *App) Update() error {
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
