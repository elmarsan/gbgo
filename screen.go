package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

// const (
// 	WindowTitle  = "GB-GO"
// 	WindowWidth  = 800
// 	WindowHeight = 600
// 	FrameRate    = 60
// )

// func run() int {
// 	var window *sdl.Window
// 	var renderer *sdl.Renderer
// 	var err error
// 	mu := &sync.Mutex{}

// 	sdl.Do(func() {
// 		window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WindowWidth, WindowHeight, sdl.WINDOW_OPENGL)
// 	})

// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
// 		return 1
// 	}

// 	defer func() {
// 		sdl.Do(func() {
// 			window.Destroy()
// 		})
// 	}()

// 	sdl.Do(func() {
// 		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
// 	})
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
// 		return 2
// 	}
// 	defer func() {
// 		sdl.Do(func() {
// 			renderer.Destroy()
// 		})
// 	}()

// 	sdl.Do(func() {
// 		renderer.Clear()
// 	})

// 	running := true
// 	for running {
// 		sdl.Do(func() {
// 			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
// 				switch event.(type) {
// 				case *sdl.QuitEvent:
// 					mu.Lock()
// 					running = false
// 					mu.Unlock()
// 				}
// 			}

// 			renderer.Clear()
// 		})

// 		// wg := sync.WaitGroup{}
// 		// 	wg.Add(1)
// 		// 	go func() {
// 		// 		Drawing
// 		// 		wg.Done()
// 		// 	}()
// 		// wg.Wait()

// 		sdl.Do(func() {
// 			renderer.Present()
// 			sdl.Delay(1000 / FrameRate)
// 		})
// 	}

// 	return 0
// }

const (
	SCREEN_WIDTH  = 1024
	SCREEN_HEIGHT = 768
	SCALE         = 4
)

var (
	window   *sdl.Window
	renderer *sdl.Renderer
	texture  *sdl.Texture
	surface  *sdl.Surface

	tileColors [4]uint32 = [4]uint32{0xFFFFFFFF, 0xFFAAAAAA, 0xFF555555, 0xFF000000}
)

func initScreen() error {
	sdl.Init(sdl.INIT_VIDEO)
	fmt.Println("SDL INIT")

	var width int32 = 16 * 8 * SCALE
	var height int32 = 32 * 8 * SCALE

	win, ren, err := sdl.CreateWindowAndRenderer(width, height, 0)
	if err != nil {
		return err
	}

	window = win
	renderer = ren

	surface, err = sdl.CreateRGBSurface(0, (16*8*SCALE)+(16*SCALE), (32*8*SCALE)+(64*SCALE), 32, 0x00FF0000, 0x0000FF00, 0x000000FF, 0xFF000000)
	if err != nil {
		return err
	}

	texture, err = renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, (16*8*SCALE)+(16*SCALE), (32*8*SCALE)+(64*SCALE))
	if err != nil {
		return err
	}

	return nil
}

func delayScreen(ms uint32) {
	sdl.Delay(ms)
}

func displayTile(startLocation uint16, tileNum uint16, x, y int32) {
	rc := &sdl.Rect{}

	for tileY := 0; tileY < 16; tileY += 2 {
		b1Addr := startLocation + (tileNum * 16) + uint16(tileY)
		b2Addr := startLocation + (tileNum * 16) + uint16(tileY) + 1

		var b1 uint8 = memory.read(b1Addr)
		var b2 uint8 = memory.read(b2Addr)

		for bit := 7; bit >= 0; bit-- {
			hi := (b1 & (1 << bit)) << 1
			lo := (b2 & (1 << bit))

			color := hi | lo

			rc.X = x + int32((7-bit)*SCALE)
			rc.Y = y + int32(tileY/2*SCALE)
			rc.W = SCALE
			rc.H = SCALE

			surface.FillRect(rc, tileColors[color])
		}
	}
}

func updateScreen() {
	var xDraw int32 = 0
	var yDraw int32 = 0
	var tileNum uint16 = 0

	rc := &sdl.Rect{}
	rc.X = 0
	rc.Y = 0
	rc.W = surface.W
	rc.H = surface.H

	surface.FillRect(rc, 0xFF111111)
	var addr uint16 = 0x8000

	//384 tiles, 24 x 16
	for y := 0; y < 24; y++ {
		for x := 0; x < 16; x++ {
			displayTile(addr, tileNum, xDraw+(int32(x)*SCALE), yDraw+(int32(y)*SCALE))
			xDraw += (8 * SCALE)
			tileNum++
		}

		yDraw += (8 * SCALE)
		xDraw = 0
	}

	pixels := surface.Pixels()

	p := &pixels[0]

	texture.Update(rc, unsafe.Pointer(p), int(surface.Pitch))
	renderer.Clear()
	renderer.Copy(texture, nil, nil)
	renderer.Present()
}

func handleEventsScreen() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.(type) {
		case *sdl.QuitEvent:
			// kill gameboy
			fmt.Println("SDL QuitEvent")
			os.Exit(0)
		}
	}
}
