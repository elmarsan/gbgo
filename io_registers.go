package main

const (
	JOYP     = 0xff00 // Joypad
	SB       = 0xff01 // Serial transfer data
	SC       = 0xff02 // Serial transfer control
	DIV      = 0xff04 // Divider register
	TIMA     = 0xff05 // Timer counter
	TMA      = 0xff06 // Timer modulo
	TAC      = 0xff07 // Timer control
	IF       = 0xff0f // Interrupt flag
	NR10     = 0xff10 // Sound channel 1 sweep
	NR11     = 0xff11 // Sound channel 1 length timer & duty cycle
	NR12     = 0xff12 // Sound channel 1 volume & envelope
	NR13     = 0xff13 // Sound channel 1 wavelength low
	NR14     = 0xff14 // Sound channel 1 wavelength high & control
	NR21     = 0xff16 // Sound channel 2 length timer & duty cycle
	NR22     = 0xff17 // Sound channel 2 volume & envelope
	NR23     = 0xff18 // Sound channel 2 wavelength low
	NR24     = 0xff19 // Sound channel 2 wavelength high & control
	NR30     = 0xff1a // Sound channel 3 DAC enable
	NR31     = 0xff1b // Sound channel 3 length timer
	NR32     = 0xff1c // Sound channel 3 output level
	NR33     = 0xff1d // Sound channel 3 wavelength low
	NR34     = 0xff1e // Sound channel 3 wavelength high & control
	NR41     = 0xff20 // Sound channel 4 length timer
	NR42     = 0xff21 // Sound channel 4 volume & envelope
	NR43     = 0xff22 // Sound channel 4 frequency & randomness
	NR44     = 0xff23 // Sound channel 4 control
	NR50     = 0xff24 // Master volume & VIN panning
	NR51     = 0xff25 // Sound panning
	NR52     = 0xff26 // Sound on/off
	WAVE_RAM = 0xff30 // (0xff30-ff3f) Storage for one of the sound channels’ waveform
	LCDC     = 0xff40 // LCD control
	STAT     = 0xff41 // LCD status	Mixed
	SCY      = 0xff42 // Viewport Y position
	SCX      = 0xff43 // Viewport X position
	LY       = 0xff44 // LCD Y coordinate
	LYC      = 0xff45 // LY compare
	DMA      = 0xff46 // OAM DMA source address & start
	BGP      = 0xff47 // BG palette data
	OBP0     = 0xff48 // OBJ palette 0 data
	OBP1     = 0xff49 // OBJ palette 1 data
	WY       = 0xff4a // Window Y position
	WX       = 0xff4b // Window X position plus 7
	IE       = 0xffff // Interrupt enable
)
