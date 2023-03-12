package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	initLogger()

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Missing rom arg")
		os.Exit(1)
	}

	rom := args[1]

	err := cartridge.load(rom)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = initScreen()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		cpu.init()
		cpu.running = true

		for cpu.running {
			cpu.execute()
			// time.Sleep(time.Second / 4)
		}
	}()

	for {
		handleEventsScreen()
		updateScreen()
	}
}
