package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		log.Fatal("Missing rom arg")
	}

	rom := args[1]

	gameboy.Run(rom)
}
