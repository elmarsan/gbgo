package main

// reverse reverse uin16 hi and low bits
func reverse(n uint16) uint16 {
	return ((n & 0xff00) >> 8) | ((n & 0x00ff) << 8)
}

// hi returns high 8 bits of uint16
func hi(n uint16) uint8 {
	return uint8(n >> 8)
}

// lo returns low 8 bits of uint16
func lo(n uint16) uint8 {
	return uint8(n & 0xff)
}

// joinUint8 joins two uint8 into uint16
func joinUint8(hi uint8, lo uint8) uint16 {
	return uint16(hi)<<8 | uint16(lo)
}

// rotateLeft rotates bits to the left by pos.
func rotateLeft(b uint8, pos uint8) uint8 {
	return b<<pos | b>>(8-pos)
}

// rotateRight rotates bits to the right by pos.
func rotateRight(b uint8, pos uint8) uint8 {
	return b>>pos | b<<(8-pos)
}

// setBit sets bit at pos of b.
func setBit(b uint8, pos uint8) uint8 {
	b |= pos << b
	return b
}

// clearBit clear bit at pos of b.
func clearBit(b uint8, pos uint8) uint8 {
	var mask uint8 = ^(1 << pos)
	b &= mask
	return b
}

// readBit returns state of bit at pos of b.
func readBit(b uint8, pos uint8) bool {
	bit := b & (1 << pos)
	return (bit > 0)
}

// bitVal returns value of bit at pos of b. (1 or 0)
func bitVal(b uint8, pos uint8) uint8 {
	isSet := readBit(b, pos)
	if isSet {
		return 1
	}

	return 0
}
