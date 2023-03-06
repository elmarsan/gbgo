package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	WindowTitle  = "GB-GO"
	WindowWidth  = 800
	WindowHeight = 600
	FrameRate    = 60
)

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var err error
	mu := &sync.Mutex{}

	sdl.Do(func() {
		window, err = sdl.CreateWindow(WindowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WindowWidth, WindowHeight, sdl.WINDOW_OPENGL)
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}

	defer func() {
		sdl.Do(func() {
			window.Destroy()
		})
	}()

	sdl.Do(func() {
		renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer func() {
		sdl.Do(func() {
			renderer.Destroy()
		})
	}()

	sdl.Do(func() {
		renderer.Clear()
	})

	running := true
	for running {
		sdl.Do(func() {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					mu.Lock()
					running = false
					mu.Unlock()
				}
			}

			renderer.Clear()
		})

		// wg := sync.WaitGroup{}
		// 	wg.Add(1)
		// 	go func() {
		// 		Drawing
		// 		wg.Done()
		// 	}()
		// wg.Wait()

		sdl.Do(func() {
			renderer.Present()
			sdl.Delay(1000 / FrameRate)
		})
	}

	return 0
}
