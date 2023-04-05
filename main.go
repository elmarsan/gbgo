package main

import (
	"fmt"
	"log"
	"os"
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
		fmt.Println(err)
		os.Exit(1)
	}

	err = initScreen()
	if err != nil {
		log.Fatal(err)
	}

	gameboy.Init()

	go func() {
		gameboy.Run()
	}()

	for {
		handleEventsScreen()
		updateScreen()
	}
}
