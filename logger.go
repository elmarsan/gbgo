package main

import (
	"fmt"
	"log"
	"os"
)

func initLogger() {
	f, err := os.Create("debug")
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.SetFlags(0)
}

func logState() {
	state := fmt.Sprintf("A:%02X F:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X SP:%04X PC:%04X PCMEM:%02X,%02X,%02X,%02X",
		cpu.a, cpu.f, cpu.b, cpu.c, cpu.d, cpu.e, cpu.h, cpu.l, cpu.sp, cpu.pc,
		memory.read(cpu.pc), memory.read(cpu.pc+1), memory.read(cpu.pc+2), memory.read(cpu.pc+3))

	log.Println(state)
}
