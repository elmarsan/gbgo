package main

import (
	"fmt"
	"log"
	"os"
)

type Gameboy struct{}

var gameboy = &Gameboy{}

func (g *Gameboy) Run(rom string) {
	err := cartridge.load(rom)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Printf("Rom %s loaded", rom)

	for {
		cpu.execute()
	}
}
