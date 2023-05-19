package bit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	assert := assert.New(t)

	t.Run("swapNibbleU16", func(_ *testing.T) {
		assert.Equal(SwapNibbleU16(0xaabb), uint16(0xbbaa))
	})

	t.Run("swapNibbleU8", func(_ *testing.T) {
		assert.Equal(SwapNibbleU8(0xef), uint8(0xfe))
	})

	t.Run("lo", func(_ *testing.T) {
		assert.Equal(Lo(0xaabb), uint8(0xbb))
	})

	t.Run("hi", func(_ *testing.T) {
		assert.Equal(Hi(0xaabb), uint8(0xaa))
	})

	t.Run("joinU8", func(_ *testing.T) {
		assert.Equal(Joinu8(0x12, 0xb8), uint16(0x12b8))
	})

	t.Run("rotateLeft", func(_ *testing.T) {
		assert.Equal(RotateLeft(0b11001010, 4), uint8(0b10101100))
	})

	t.Run("rotateRight", func(_ *testing.T) {
		assert.Equal(RotateRight(0b11001010, 2), uint8(0b10110010))
	})

	t.Run("readBit", func(_ *testing.T) {
		assert.NotEqual(uint8(Read(0b11001010, 7)), uint8(0))
		assert.Equal(uint8(Read(0b11001010, 5)), uint8(0))
	})

	t.Run("clearBit", func(_ *testing.T) {
		x := Clear(0b11001010, 1)
		assert.Equal(uint8(Read(x, 1)), uint8(0))
	})

	t.Run("isBitSet", func(_ *testing.T) {
		assert.True(IsSet(0xf0, 7))
		assert.False(IsSet(0xf0, 0))
	})

	t.Run("toggleBit", func(_ *testing.T) {
		b := Toggle(0xf0, 7, false)
		assert.False(IsSet(b, 7))
	})
}
