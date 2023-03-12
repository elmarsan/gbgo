package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var val uint16 = 0xaabb
	reversed := reverse(val)

	if reversed != 0xbbaa {
		t.Error("Wrong reverse")
	}
}

func TestLo(t *testing.T) {
	var val uint16 = 0xaabb
	lo := lo(val)

	if lo != 0xbb {
		t.Error("Wrong lo")
	}
}

func TestHi(t *testing.T) {
	var val uint16 = 0xaabb
	hi := hi(val)

	if hi != 0xaa {
		t.Error("Wrong hi")
	}
}

func TestJoinUint8(t *testing.T) {
	var n1 uint8 = 0x12
	var n2 uint8 = 0xb8

	var val = joinUint8(n1, n2)

	if val != 0x12b8 {
		t.Error("Wrong uint8 joining")
	}
}

func TestRotateLeft(t *testing.T) {
	var x uint8 = 0b11001010

	rotated := rotateLeft(x, 4)

	if rotated != 0b10101100 {
		t.Error("Wrong left rotation")
	}
}

func TestRotateRight(t *testing.T) {
	var x uint8 = 0b11001010

	rotated := rotateRight(x, 2)

	if rotated != 0b10110010 {
		t.Error("Wrong right rotation")
	}
}

func TestReadBit(t *testing.T) {
	t.Run("Bit enabled", func(t *testing.T) {
		var x uint8 = 0b11001010

		enabled := readBit(x, 7)

		if enabled == 0 {
			t.Error("Wrong bit reading")
		}
	})

	t.Run("Bit disabled", func(t *testing.T) {
		var x uint8 = 0b11001010

		enabled := readBit(x, 5)

		if enabled == 1 {
			t.Error("Wrong bit reading")
		}
	})
}

func TestClearBit(t *testing.T) {
	x := clearBit(0b11001010, 1)

	enabled := readBit(x, 1)
	if enabled == 1 {
		t.Error("Wrong bit reading")
	}
}
