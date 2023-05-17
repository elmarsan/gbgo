package main

import (
	"log"
	"os"
)

func main() {
	// Create a new Gameboy instance
	gb := NewGameboy()

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

	app := NewApp(gb)

	if err := app.Run(); err != nil {
		log.Fatal(err)
		gb.Stop()
	}
}
