package main

import (
	"image"
	"image/color"

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
	imgRGBA := a.createImgRGBA(ppu.pixels[:])
	img := ebiten.NewImageFromImage(imgRGBA)

	drawopts := &ebiten.DrawImageOptions{}
	drawopts.GeoM.Scale(4, 3.33)
	screen.DrawImage(img, drawopts)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (a *App) createImgRGBA(pixels []uint32) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, GB_W, GB_H))

	for y := 0; y < GB_H; y++ {
		for x := 0; x < GB_W; x++ {
			i := y*GB_W + x
			c := color.RGBA{
				R: uint8((pixels[i] >> 16) & 0xff),
				G: uint8((pixels[i] >> 8) & 0xff),
				B: uint8(pixels[i] & 0xff),
				A: uint8((pixels[i] >> 24) & 0xff),
			}

			img.SetRGBA(x, y, c)
		}
	}

	return img
}
