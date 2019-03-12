package hash

// Murmur calculates murmur hash
func Murmur(data []byte, seed int32) int32 {
	m, r, length := int32(0x5bd1e995), int32(24), int32(len(data))
	h, len4 := seed^length, length>>2

	for i := int32(0); i < len4; i++ {
		i4 := i << 2
		k := int32(data[i4+3])
		k = k << 8
		k = k | int32(data[i4+2]&0xff)
		k = k << 8
		k = k | int32(data[i4+1]&0xff)
		k = k << 8
		k = k | int32(data[i4+0]&0xff)
		k *= m
		k ^= int32(uint32(k) >> uint32(r))
		k *= m
		h *= m
		h ^= k
	}

	lenm := len4 << 2
	left := length - lenm

	if left != 0 {
		if left >= 3 {
			h ^= int32(data[length-3]) << 16
		}
		if left >= 2 {
			h ^= int32(data[length-2]) << 8
		}
		if left >= 1 {
			h ^= int32(data[length-1])
		}

		h *= m
	}

	h ^= int32(uint32(h) >> 13)
	h *= m
	h ^= int32(uint32(h) >> 15)

	return h
}
