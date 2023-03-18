package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytes(t *testing.T) {
	assert := assert.New(t)

	t.Run("swapNibbleU16", func(t *testing.T) {
		assert.Equal(swapNibbleU16(0xaabb), uint16(0xbbaa))
	})

	t.Run("swapNibbleU8", func(t *testing.T) {
		assert.Equal(swapNibbleU8(0xef), uint8(0xfe))
	})

	t.Run("lo", func(t *testing.T) {
		assert.Equal(lo(0xaabb), uint8(0xbb))
	})

	t.Run("hi", func(t *testing.T) {
		assert.Equal(hi(0xaabb), uint8(0xaa))
	})

	t.Run("joinU8", func(t *testing.T) {
		assert.Equal(joinu8(0x12, 0xb8), uint16(0x12b8))
	})

	t.Run("rotateLeft", func(t *testing.T) {
		assert.Equal(rotateLeft(0b11001010, 4), uint8(0b10101100))
	})

	t.Run("rotateRight", func(t *testing.T) {
		assert.Equal(rotateRight(0b11001010, 2), uint8(0b10110010))
	})

	t.Run("readBit", func(t *testing.T) {
		assert.NotEqual(uint8(readBit(0b11001010, 7)), uint8(0))
		assert.Equal(uint8(readBit(0b11001010, 5)), uint8(0))
	})

	t.Run("clearBit", func(t *testing.T) {
		x := clearBit(0b11001010, 1)
		assert.Equal(uint8(readBit(x, 1)), uint8(0))
	})

	t.Run("isBitSet", func(t *testing.T) {
		assert.True(isBitSet(0xf0, 7))
		assert.False(isBitSet(0xf0, 0))
	})

	t.Run("toggleBit", func(t *testing.T) {
		b := toggleBit(0xf0, 7, false)
		assert.False(isBitSet(b, 7))
	})
}
