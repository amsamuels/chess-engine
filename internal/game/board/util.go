package board

// Calculate the union of integers.
func Union(i ...uint64) uint64 {
	var u uint64
	for _, v := range i {
		u = u | v
	}
	return u
}

// GetBit returns the value of the bit at position p.
func GetBit(i *uint64, p int) int {
	return int((*i >> uint(p)) & 1)
}

// SetBit sets (sets to 1) the bit at position p.
func SetBit(i *uint64, p int) {
	var mask uint64
	mask = (1 << uint(p))
	*i |= mask
}

// ClearBit clears (sets to 0) the bit at position p.
func ClearBit(i *uint64, p int) {
	var mask uint64
	mask = ^(1 << uint(p))
	*i &= mask
}

// ToggleBit toggles the value of the bit at position p.
func ToggleBit(i *uint64, p int) {
	var mask uint64
	mask = (1 << uint(p))
	*i ^= mask
}

func IsBitSet(i uint64, p int) bool {
	var mask uint64
	mask = 1 << uint64(p)
	return (i & mask) != 0
}
