package main

import "log"

type Gameboy struct{}

var gameboy = &Gameboy{}

func (g *Gameboy) Run(rom string) {
	err := cartridge.load(rom)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rom %s loaded", rom)

	for {
		cpu.cycle()
	}
}
