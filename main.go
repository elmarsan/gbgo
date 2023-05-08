package main

import (
	"log"
	"os"
)

// Create a new Gameboy instance
var gb = NewGameboy()

func main() {
	// debug stuff
	debug.init()

	args := os.Args
	if len(args) != 2 {
		log.Fatal("Missing rom arg")
	}

	rom := args[1]
	err := gb.LoadRom(rom)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		gb.Run()
	}()

	app := NewApp()

	if err := app.Run(); err != nil {
		log.Fatal(err)
		gb.Stop()
	}
}
