package main

import (
	"log"
	"os"

	"github.com/elmarsan/gbgo/pkg/emulator"
)

func main() {
	// Create a new Emulator instance
	emu := emulator.New()

	args := os.Args
	if len(args) != 2 {
		log.Fatal("Missing rom arg")
	}

	romPath := args[1]
	err := emu.Gb.LoadRom(romPath)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		emu.Gb.Run()
	}()

	app := NewApp(emu)

	if err := app.Run(); err != nil {
		emu.Gb.Stop()
		log.Fatal(err)
	}
}
