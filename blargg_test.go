package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type blarggTest struct {
	name     string
	targetPc uint16
	seconds  float64
}

func (bt *blarggTest) run(t *testing.T) {
	gameboy = &Gameboy{}
	memory = &Memory{}
	cpu = &CPU{}
	ppu = NewPPU()
	cartridge = &Cartridge{}
	timer = &Timer{}
	app = &App{}
	joypad = &Joypad{}

	cpu.init()
	memory.init()

	t.Run(bt.name, func(t *testing.T) {
		t.Logf("Running: %s\n", bt.name)

		path := fmt.Sprintf("rom/%s", bt.name)
		err := gameboy.LoadRom(path)
		if err != nil {
			t.Fatalf("Unable to load rom %s\n", path)
		}

		start := time.Now()
		errChan := make(chan error, 1)

		go func() {
			for {
				if time.Since(start).Seconds() > bt.seconds {
					msg := fmt.Sprintf("ERROR: %s\n", bt.name)
					errChan <- errors.New(msg)
					break
				}

				if cpu.pc == bt.targetPc {
					errChan <- nil
					break
				}
			}
		}()

		go func() {
			gameboy.Run()
		}()

		for {
			err := <-errChan
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("PASSED: %s\n", bt.name)
			close(errChan)
			break
		}
	})
}

func TestBlargg(t *testing.T) {
	tests := []blarggTest{
		{
			name:     "01-special.gb",
			targetPc: 0xc7d2,
			seconds:  60,
		},
		{
			name:     "02-interrupts.gb",
			targetPc: 0xc7f4,
			seconds:  60,
		},
		{
			name:     "03-op sp,hl.gb",
			targetPc: 0xcb44,
			seconds:  60,
		},
		{
			name:     "04-op sp,hl.gb",
			targetPc: 0xcb35,
			seconds:  60,
		},
		{
			name:     "05-op rp.gb",
			targetPc: 0xcb31,
			seconds:  60,
		},
		{
			name:     "06-ld r,r.gb",
			targetPc: 0xcc5f,
			seconds:  60,
		},
		{
			name:     "07-jr,jp,call,ret,rst.gb",
			targetPc: 0xcbb0,
			seconds:  60,
		},
		{
			name:     "08-misc instrs.gb",
			targetPc: 0xcb91,
			seconds:  60,
		},
		{
			name:     "09-op r,r.gb",
			targetPc: 0xce67,
			seconds:  120,
		},
		{
			name:     "10-bit ops.gb",
			targetPc: 0xcf58,
			seconds:  120,
		},
		{
			name:     "11-op a,(hl).gb",
			targetPc: 0xcc62,
			seconds:  120,
		},
	}

	for _, test := range tests {
		test.run(t)
	}
}
