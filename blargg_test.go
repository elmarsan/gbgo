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
}

func (bt *blarggTest) run() error {
	path := fmt.Sprintf("rom/%s", bt.name)
	err := gb.LoadRom(path)
	if err != nil {
		return err
	}

	start := time.Now()
	errChan := make(chan error, 1)

	go func() {
		for {
			// If pc does not have targetPc value in 30 seconds, stop and return error.
			if time.Since(start).Seconds() > 30 {
				msg := fmt.Sprintf("Rom %s unable to reached target pc: %04x", bt.name, bt.targetPc)
				errChan <- errors.New(msg)
				break
			}

			if gb.cpu.pc == bt.targetPc {
				errChan <- nil
				break
			}
		}

		// After timeout of success stop Gameboy loop.
		gb.Stop()
	}()

	// goroutine that execute gameboy steps
	// it stops when target pc or timeout is reached
	go func() {
		gb.Run()
	}()

	err = <-errChan
	close(errChan)
	return err
}

func TestBlargg(t *testing.T) {
	tests := []blarggTest{
		{
			name:     "01-special.gb",
			targetPc: 0xc7d2,
		},
		{
			name:     "02-interrupts.gb",
			targetPc: 0xc7f4,
		},
		{
			name:     "03-op sp,hl.gb",
			targetPc: 0xcb44,
		},
		{
			name:     "04-op r,imm.gb",
			targetPc: 0xcb35,
		},
		{
			name:     "05-op rp.gb",
			targetPc: 0xcb31,
		},
		{
			name:     "06-ld r,r.gb",
			targetPc: 0xcc5f,
		},
		{
			name:     "07-jr,jp,call,ret,rst.gb",
			targetPc: 0xcbb0,
		},
		{
			name:     "08-misc instrs.gb",
			targetPc: 0xcb91,
		},
		{
			name:     "09-op r,r.gb",
			targetPc: 0xce67,
		},
		{
			name:     "10-bit ops.gb",
			targetPc: 0xcf58,
		},
		{
			name:     "11-op a,(hl).gb",
			targetPc: 0xcc62,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.run()
			if err != nil {
				t.Error(err)
			}
		})
	}
}
