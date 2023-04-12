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
	ppu       = &PPU{}
	cartridge = &Cartridge{}
	lcd       = &LCD{}
	timer     = &Timer{}
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

	err = initScreen()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		gameboy.Run()
	}()

	for {
		handleEventsScreen()
		updateScreen()
	}
}
