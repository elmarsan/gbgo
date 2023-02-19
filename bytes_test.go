package main

import "testing"

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
