package main

import (
	"image"

	"github.com/elmarsan/gbgo/pkg/emulator"
	"github.com/elmarsan/gbgo/pkg/gameboy"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	// WindowWidth represents the window width of application.
	WindowWidth = 640

	// WindowHeight represents the window height of application.
	WindowHeight = 480
)

// App represents ebiten.Game implementation.
type App struct {
	emu        *emulator.Emulator
	keyHandler map[ebiten.Key]func(down bool)
}

// NewApp returns new instance of App.
func NewApp(emu *emulator.Emulator) *App {
	return &App{
		emu: emu,
		keyHandler: map[ebiten.Key]func(down bool){
			ebiten.KeyArrowUp:    func(down bool) { emu.SetButton(emulator.UpBtn, down) },
			ebiten.KeyArrowDown:  func(down bool) { emu.SetButton(emulator.DownBtn, down) },
			ebiten.KeyArrowLeft:  func(down bool) { emu.SetButton(emulator.LeftBtn, down) },
			ebiten.KeyArrowRight: func(down bool) { emu.SetButton(emulator.RightBtn, down) },
			ebiten.KeySpace:      func(down bool) { emu.SetButton(emulator.SelectBtn, down) },
			ebiten.KeyK:          func(down bool) { emu.SetButton(emulator.StartBtn, down) },
			ebiten.KeyA:          func(down bool) { emu.SetButton(emulator.ABtn, down) },
			ebiten.KeyB:          func(down bool) { emu.SetButton(emulator.BBtn, down) },
		},
	}
}

func (a *App) Run() error {
	ebiten.SetWindowSize(WindowWidth, WindowHeight)
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
	img := a.pixelsToImage()
	drawopts := &ebiten.DrawImageOptions{}
	drawopts.GeoM.Scale(4, 3.33)
	screen.DrawImage(img, drawopts)
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (a *App) pixelsToImage() *ebiten.Image {
	imgRGBA := image.NewRGBA(image.Rect(0, 0, gameboy.ScreenWidth, gameboy.ScreenHeight))
	pixels := a.emu.GetRGBAPixels()

	for y := 0; y < gameboy.ScreenHeight; y++ {
		for x := 0; x < gameboy.ScreenWidth; x++ {
			i := y*gameboy.ScreenWidth + x
			imgRGBA.SetRGBA(x, y, pixels[i])
		}
	}

	return ebiten.NewImageFromImage(imgRGBA)
}
