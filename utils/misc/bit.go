package misc

func GetBit64(bits uint64, index uint) bool {
	return bits>>index&1 != 0
}

func SetBit64(bits uint64, index uint) uint64 {
	bits |= uint64(1 << index)
	return bits
}
func GetBit32(bits uint32, index uint) bool {
	return bits>>index&1 != 0
}

func SetBit32(bits uint32, index uint) uint32 {
	bits |= uint32(1 << index)
	return bits
}

func ClearBit32(bits uint32, index uint) uint32 {
	bits &= ^uint32(1 << index)
	return bits
}

func GetBit8(bits uint8, index uint) bool {
	return bits>>index&1 != 0
}

func SetBit8(bits uint8, index uint) uint8 {
	bits |= uint8(1 << index)
	return bits
}

func ClearBit8(bits uint8, index uint) uint8 {
	bits &= ^uint8(1 << index)
	return bits
}

func CountBit8(bits uint8) int {
	result := 0
	for i := uint(0); i < 8; i++ {
		if bits>>i&1 != 0 {
			result++
		}
	}
	return result
}
