package main

import (
	"fmt"
	"log"
	"os"
)

var (
	gameboy   = &Gameboy{}
	memory    = &Memory{}
	cpu       = &CPU{}
	ppu       = NewPPU()
	cartridge = &Cartridge{}
	timer     = &Timer{}
	app       = &App{}
	joypad    = &Joypad{}
)

func main() {
	debug.init()

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Missing rom arg")
		os.Exit(1)
	}

	rom := args[1]
	err := gameboy.LoadRom(rom)
	if err != nil {
		log.Fatal(err)
	}

	cpu.init()
	memory.init()

	go func() {
		gameboy.Run()
	}()

	app := NewApp()

	for {
		if err := app.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
