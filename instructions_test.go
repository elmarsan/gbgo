package main

import "testing"

func Test0x4(t *testing.T) {
	t.Run("0x40: LD B, B", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_B, val)
		instructions[0x40]()

		if cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x41: LD B, C", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_C, val)
		instructions[0x41]()

		if cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x42: LD B, D", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_D, val)
		instructions[0x42]()

		if cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x43: LD B, E", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_E, val)
		instructions[0x43]()

		if cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x44 -> LD B, H", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_H, val)
		instructions[0x44]()

		if cpu.read8Reg(REG_B) != val {
			t.Error("Wrong execution: 0x44 -> LD B, H")
		}
	})

	t.Run("0x45: LD B, L", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_L, val)
		instructions[0x45]()

		if cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x46 -> LD B, HL", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x47: LD B, A", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_A, val)
		instructions[0x47]()

		if cpu.read8Reg(REG_B) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x48: LD C, B", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_B, val)
		instructions[0x48]()

		if cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x49: LD C, C", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_C, val)
		instructions[0x49]()

		if cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4a: LD C, D", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_D, val)
		instructions[0x4a]()

		if cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4b: LD C, E", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_E, val)
		instructions[0x4b]()

		if cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4c: LD C, H", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_H, val)
		instructions[0x4c]()

		if cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4d :LD C, L", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_L, val)
		instructions[0x4d]()

		if cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x4e: LD C, HL", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x4f: LD C, A", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_A, val)
		instructions[0x4f]()

		if cpu.read8Reg(REG_C) != val {
			t.Error("Unexpected behaviour")
		}
	})
}

func Test0x5(t *testing.T) {
	t.Run("0x50: LD D, B", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_B, val)
		instructions[0x50]()

		if cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x51: LD D, C", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_C, val)
		instructions[0x51]()

		if cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x52: LD D, D", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_D, val)
		instructions[0x52]()

		if cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x53: LD D, E", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_E, val)
		instructions[0x53]()

		if cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x54: LD D, H", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_H, val)
		instructions[0x54]()

		if cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x55: LD D, L", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_L, val)
		instructions[0x55]()

		if cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x56: LD D, HL", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x57: LD D, A", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_A, val)
		instructions[0x57]()

		if cpu.read8Reg(REG_D) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x58: LD E, B", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_B, val)
		instructions[0x58]()

		if cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x59: LD E, C", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_C, val)
		instructions[0x59]()

		if cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5a: LD E, D", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_D, val)
		instructions[0x5a]()

		if cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5b: LD E, E", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_E, val)
		instructions[0x5b]()

		if cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5c: LD E, H", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_H, val)
		instructions[0x5c]()

		if cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5d: LD E, L", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_L, val)
		instructions[0x5d]()

		if cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x5e: LD E, HL", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x5f: LD E, A", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_A, val)
		instructions[0x5f]()

		if cpu.read8Reg(REG_E) != val {
			t.Error("Unexpected behaviour")
		}
	})
}

func Test0x6(t *testing.T) {
	t.Run("0x60: LD H, B", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_B, val)
		instructions[0x60]()

		if cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x61: LD H, C", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_C, val)
		instructions[0x61]()

		if cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x62: LD H, D", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_D, val)
		instructions[0x62]()

		if cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x63: LD H, E", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_E, val)
		instructions[0x63]()

		if cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x64: LD H, H", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_H, val)
		instructions[0x64]()

		if cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x65: LD H, L", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_L, val)
		instructions[0x65]()

		if cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x66: LD H, (HL)", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x67: LD H, A", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_A, val)
		instructions[0x67]()

		if cpu.read8Reg(REG_H) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x68: LD L, B", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_B, val)
		instructions[0x68]()

		if cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x69: LD L, C", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_C, val)
		instructions[0x69]()

		if cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6a: LD L, D", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_D, val)
		instructions[0x6a]()

		if cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6b: LD L, E", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_E, val)
		instructions[0x6b]()

		if cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6c: LD L, H", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_H, val)
		instructions[0x6c]()

		if cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6d: LD L, L", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_L, val)
		instructions[0x6d]()

		if cpu.read8Reg(REG_L) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x6e: LD L, (HL)", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x6f: LD L, A", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_A, val)
		instructions[0x6f]()

		if cpu.read8Reg(REG_L) != val {
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
		cpu.set8Reg(REG_B, val)
		instructions[0x78]()

		if cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x79: LD A, C", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_C, val)
		instructions[0x79]()

		if cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7a: LD A, D", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_D, val)
		instructions[0x7a]()

		if cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7b: LD A, E", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_E, val)
		instructions[0x7b]()

		if cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7c: LD A, H", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_H, val)
		instructions[0x7c]()

		if cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7d: LD A, L", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_L, val)
		instructions[0x7d]()

		if cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})

	t.Run("0x7e: LD A, (HL)", func(t *testing.T) {
		t.Skip()
	})

	t.Run("0x7f: LD A, A", func(t *testing.T) {
		var val uint8 = 0xab
		cpu.set8Reg(REG_A, val)
		instructions[0x7f]()

		if cpu.read8Reg(REG_A) != val {
			t.Error("Unexpected behaviour")
		}
	})
}
