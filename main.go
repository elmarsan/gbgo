package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.SetFlags(0)

	args := os.Args

	if len(args) != 2 {
		log.Fatal("Missing rom arg")
	}

	rom := args[1]

	err := cartridge.load(rom)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rom %s loaded", rom)

	// fmt.Printf("0x100: 0x%x\n", cartridge.read(0x100))
	// fmt.Printf("0x101: 0x%x\n", cartridge.read(0x101))
	// fmt.Printf("0x102: 0x%x\n", cartridge.read(0x102))
	// fmt.Printf("0x103: 0x%x\n", cartridge.read(0x103))
	// fmt.Printf("0x104: 0x%x\n", cartridge.read(0x104))
	// fmt.Printf("0x105: 0x%x\n", cartridge.read(0x105))
	// fmt.Printf("0x106: 0x%x\n", cartridge.read(0x106))
	// fmt.Printf("0x107: 0x%x\n", cartridge.read(0x107))
	// fmt.Printf("0x108: 0x%x\n", cartridge.read(0x108))
	// fmt.Printf("0x109: 0x%x\n", cartridge.read(0x109))
	// fmt.Printf("0x10a: 0x%x\n", cartridge.read(0x10a))
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()
	// os.Exit(1)

	err = initScreen()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		cpu.init()
		cpu.running = true

		for cpu.running {
			cpu.execute()
			time.Sleep(time.Second / 2)
		}
	}()

	for {
		handleEventsScreen()
		updateScreen()
	}
}
