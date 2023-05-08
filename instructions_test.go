package main

import "testing"

func Test0x0(t *testing.T) {
	t.Run("0x02: LD (BC), A", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x06: LD B, d8", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x0a: LD A, (BC)", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x0e: LD C, d8", func(t *testing.T) {
		t.Skip()
	})
}

func Test0x1(t *testing.T) {
	t.Run("0x12: LD (DE), A", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x16: LD D, d8", func(t *testing.T) {
		t.Skip()
	})
}

func Test0x2(t *testing.T) {
	t.Run("0x2a: LD A, (HL+)", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x26: LD H, d8", func(t *testing.T) {
		t.Skip()
	})
}

func Test0x3(t *testing.T) {
	t.Run("0x32: LD (HL-), A", func(t *testing.T) {
		t.Skip()
	})
	t.Run("0x36: LD (HL), d8", func(t *testing.T) {
		t.Skip()
	})
}

func Test0x4(t *testing.T) {
	t.Run("0x40: LD B, B", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_B, val)
		instructions[0x40]()

		if gb.cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x41: LD B, C", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_C, val)
		instructions[0x41]()

		if gb.cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x42: LD B, D", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_D, val)
		instructions[0x42]()

		if gb.cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x43: LD B, E", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_E, val)
		instructions[0x43]()

		if gb.cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x44 -> LD B, H", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_H, val)
		instructions[0x44]()

		if gb.cpu.read8Reg(REG_B) != val {
			t.Error("Wrong execution: 0x44 -> LD B, H")
		}
	})

	t.Run("0x45: LD B, L", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_L, val)
		instructions[0x45]()

		if gb.cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x46 -> LD B, HL", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x47: LD B, A", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_A, val)
		instructions[0x47]()

		if gb.cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x48: LD C, B", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_B, val)
		instructions[0x48]()

		if gb.cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x49: LD C, C", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_C, val)
		instructions[0x49]()

		if gb.cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4a: LD C, D", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_D, val)
		instructions[0x4a]()

		if gb.cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4b: LD C, E", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_E, val)
		instructions[0x4b]()

		if gb.cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4c: LD C, H", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_H, val)
		instructions[0x4c]()

		if gb.cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4d :LD C, L", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_L, val)
		instructions[0x4d]()

		if gb.cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4e: LD C, HL", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x4f: LD C, A", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_A, val)
		instructions[0x4f]()

		if gb.cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})
}

func Test0x5(t *testing.T) {
	t.Run("0x50: LD D, B", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_B, val)
		instructions[0x50]()

		if gb.cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x51: LD D, C", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_C, val)
		instructions[0x51]()

		if gb.cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x52: LD D, D", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_D, val)
		instructions[0x52]()

		if gb.cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x53: LD D, E", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_E, val)
		instructions[0x53]()

		if gb.cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x54: LD D, H", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_H, val)
		instructions[0x54]()

		if gb.cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x55: LD D, L", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_L, val)
		instructions[0x55]()

		if gb.cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x56: LD D, HL", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x57: LD D, A", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_A, val)
		instructions[0x57]()

		if gb.cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x58: LD E, B", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_B, val)
		instructions[0x58]()

		if gb.cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x59: LD E, C", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_C, val)
		instructions[0x59]()

		if gb.cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5a: LD E, D", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_D, val)
		instructions[0x5a]()

		if gb.cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5b: LD E, E", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_E, val)
		instructions[0x5b]()

		if gb.cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5c: LD E, H", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_H, val)
		instructions[0x5c]()

		if gb.cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5d: LD E, L", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_L, val)
		instructions[0x5d]()

		if gb.cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5e: LD E, HL", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x5f: LD E, A", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_A, val)
		instructions[0x5f]()

		if gb.cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})
}

func Test0x6(t *testing.T) {
	t.Run("0x60: LD H, B", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_B, val)
		instructions[0x60]()

		if gb.cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x61: LD H, C", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_C, val)
		instructions[0x61]()

		if gb.cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x62: LD H, D", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_D, val)
		instructions[0x62]()

		if gb.cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x63: LD H, E", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_E, val)
		instructions[0x63]()

		if gb.cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x64: LD H, H", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_H, val)
		instructions[0x64]()

		if gb.cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x65: LD H, L", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_L, val)
		instructions[0x65]()

		if gb.cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x66: LD H, (HL)", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x67: LD H, A", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_A, val)
		instructions[0x67]()

		if gb.cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x68: LD L, B", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_B, val)
		instructions[0x68]()

		if gb.cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x69: LD L, C", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_C, val)
		instructions[0x69]()

		if gb.cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6a: LD L, D", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_D, val)
		instructions[0x6a]()

		if gb.cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6b: LD L, E", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_E, val)
		instructions[0x6b]()

		if gb.cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6c: LD L, H", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_H, val)
		instructions[0x6c]()

		if gb.cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6d: LD L, L", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_L, val)
		instructions[0x6d]()

		if gb.cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6e: LD L, (HL)", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x6f: LD L, A", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_A, val)
		instructions[0x6f]()

		if gb.cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})
}

func Test0x7(t *testing.T) {
	t.Run("0x70: LD (HL), B", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x71: LD (HL), C", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x72: LD (HL), D", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x73: LD (HL), E", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x74: LD (HL), H", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x75: LD (HL), L", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x76: HALT", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x77: LD (HL), A", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x78: LD A, B", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_B, val)
		instructions[0x78]()

		if gb.cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x79: LD A, C", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_C, val)
		instructions[0x79]()

		if gb.cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7a: LD A, D", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_D, val)
		instructions[0x7a]()

		if gb.cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7b: LD A, E", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_E, val)
		instructions[0x7b]()

		if gb.cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7c: LD A, H", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_H, val)
		instructions[0x7c]()

		if gb.cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7d: LD A, L", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_L, val)
		instructions[0x7d]()

		if gb.cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7e: LD A, (HL)", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x7f: LD A, A", func(t *testing.T) {
		var val uint8 = 0xab
		gb.cpu.set8Reg(REG_A, val)
		instructions[0x7f]()

		if gb.cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})
}
