package bit

// SwapNibbleU16 swaps low and hi nibbles of uint16.
func SwapNibbleU16(n uint16) uint16 {
	return ((n & 0xff00) >> 8) | ((n & 0x00ff) << 8)
}

// SwapNibbleU8 swaps low and hi nibbles of uint8.
func SwapNibbleU8(b uint8) uint8 {
	return ((b & 0xf0) >> 4) | ((b & 0x0f) << 4)
}

// Hi returns high 8 bits of uint16
func Hi(n uint16) uint8 {
	return uint8(n >> 8)
}

// Lo returns low 8 bits of uint16
func Lo(n uint16) uint8 {
	return uint8(n & 0xff)
}

// Joinu8 joins two uint8 into uint16
func Joinu8(hi uint8, lo uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}

// RotateLeft rotates bits to the left by pos.
func RotateLeft(b uint8, pos uint8) uint8 {
	return b<<pos | b>>(8-pos)
}

// RotateRight rotates bits to the right by pos.
func RotateRight(b uint8, pos uint8) uint8 {
	return b>>pos | b<<(8-pos)
}

// Set sets bit at pos of b.
func Set(b uint8, pos uint8) uint8 {
	return b | (1 << pos)
}

// Clear clear bit at pos of b.
func Clear(b uint8, pos uint8) uint8 {
	return b & ^(1 << pos)
}

// Read returns state of bit at pos of b.
func Read(b uint8, pos uint8) uint8 {
	return b & (1 << pos)
}

// IsSet returns boolean if bit = 1 otherwise 0
func IsSet(b uint8, pos uint8) bool {
	return (b & (1 << pos)) != 0
}

// Toggle set or clear bit at pos of b.
func Toggle(b uint8, pos uint8, on bool) uint8 {
	if on {
		return Set(b, pos)
	}

	return Clear(b, pos)
}
