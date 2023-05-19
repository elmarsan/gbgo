package gameboy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPU(t *testing.T) {
	assert := assert.New(t)

	cpu := &cpu{}

	t.Run("Read flags", func(_ *testing.T) {
		cpu.af.val = 0b10010010

		assert.True(cpu.C(), "Error reading C flag")
		assert.False(cpu.H(), "Error reading H flag")
		assert.False(cpu.N(), "Error reading N flag")
		assert.True(cpu.Z(), "Error reading Z flag")

		cpu.af.val = 0b1100010

		assert.False(cpu.C(), "Error reading C flag")
		assert.True(cpu.H(), "Error reading H flag")
		assert.True(cpu.N(), "Error reading N flag")
		assert.False(cpu.Z(), "Error reading Z flag")
	})

	t.Run("Set flags", func(_ *testing.T) {
		cpu.af.val = 0
		cpu.SetC(true)
		cpu.SetH(true)
		cpu.SetN(true)
		cpu.SetZ(true)

		assert.True(cpu.C(), "Unable to turn on C flag")
		assert.True(cpu.H(), "Unable to turn on H flag")
		assert.True(cpu.N(), "Unable to turn on N flag")
		assert.True(cpu.Z(), "Unable to turn on Z flag")

		cpu.af.val = 0xff
		cpu.SetC(false)
		cpu.SetH(false)
		cpu.SetN(false)
		cpu.SetZ(false)

		assert.False(cpu.C(), "Unable to turn off C flag")
		assert.False(cpu.H(), "Unable to turn off H flag")
		assert.False(cpu.N(), "Unable to turn off N flag")
		assert.False(cpu.Z(), "Unable to turn off Z flag")
	})
}
