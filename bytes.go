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
func joinUint8(n1 uint8, n2 uint8) uint16 {
	return uint16(n1)<<8 | uint16(n2)
}
