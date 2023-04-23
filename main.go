package main

import (
	"fmt"
	"log"
	"os"
)

var (
	gameboy = &Gameboy{}
	memory  = &Memory{}
	cpu     = &CPU{}
	ppu     = &PPU{
		videoBuf: [GB_W * GB_H]uint8{},
	}
	cartridge = &Cartridge{}
	timer     = &Timer{}
	app       = &App{}
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

	memory.init()

	ppu.clearScreen()

	go func() {
		gameboy.Run()
	}()

	for {
		if err := app.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
