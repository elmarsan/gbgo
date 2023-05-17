package main

// CpuRegister represent 16 bits cpu register.
type CpuRegister struct {
	// val holds register value.
	val uint16
}

// Get returns register value.
func (r *CpuRegister) Get() uint16 {
	return r.val
}

// Set modifies register value.
func (r *CpuRegister) Set(val uint16) {
	r.val = val
}

// Hi returns higher 8 bits of register value.
func (r *CpuRegister) Hi() uint8 {
	return uint8(r.val >> 8)
}

// SetHi modifies higher 8 bits of register value.
func (r *CpuRegister) SetHi(val uint8) {
	r.val = uint16(val)<<8 | (uint16(r.val) & 0xff)
}

// Lo returns lower 8 bits of register value.
func (r *CpuRegister) Lo() uint8 {
	return uint8(r.val & 0xff)
}

// SetLo modifies lower 8 bits of register value.
func (r *CpuRegister) SetLo(val uint8) {
	r.val = uint16(val) | (uint16(r.val) & 0xff00)
}
