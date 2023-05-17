package main

import (
	"fmt"
	"log"
	"os"
)

type Debug struct {
	fileLogger   *log.Logger
	stdOutLogger *log.Logger
	msg          string
}

var debug = &Debug{}

func (d *Debug) init() {
	f, err := os.Create("debug")
	if err != nil {
		log.Fatal(err)
	}

	fileLogger := &log.Logger{}
	fileLogger.SetOutput(f)
	fileLogger.SetFlags(0)
	d.fileLogger = fileLogger

	stdOutLogger := &log.Logger{}
	stdOutLogger.SetOutput(os.Stdout)
	stdOutLogger.SetFlags(0)
	d.stdOutLogger = stdOutLogger
}

func (d *Debug) logState(gb *Gameboy) {
	state := fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X",
		gb.cpu.af.Hi(), gb.cpu.af.Lo(), gb.cpu.bc.Hi(), gb.cpu.bc.Lo(), gb.cpu.de.Hi(), gb.cpu.de.Lo(), gb.cpu.hl.Hi(),
		gb.cpu.hl.Lo(), gb.cpu.sp, gb.cpu.pc, gb.memoryBus.read(gb.cpu.pc), gb.memoryBus.read(gb.cpu.pc+1), gb.memoryBus.read(gb.cpu.pc+2), gb.memoryBus.read(gb.cpu.pc+3))

	if gb.memoryBus.read(0xff02) == 0x81 {
		char := gb.memoryBus.read(0xff01)
		d.msg += string(char)

		gb.memoryBus.write(0xff02, 0)
	}

	if len(d.msg) > 0 {
		d.stdOutLogger.Printf("Message: %s\n", d.msg)
	}

	d.fileLogger.Println(state)
}
