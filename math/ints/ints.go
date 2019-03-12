package ints

// MinByte returns minimum between two byte
func MinByte(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

// MinInt returns minimum between two int
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinInt8 returns minimum between two int8
func MinInt8(a, b int8) int8 {
	if a < b {
		return a
	}
	return b
}

// MinInt16 returns minimum between two int16
func MinInt16(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

// MinInt32 returns minimum between two int32
func MinInt32(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

// MinInt64 returns minimum between two int64
func MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// MinUint returns minimum between two uint
func MinUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

// MinUint8 returns minimum between two uint8
func MinUint8(a, b uint8) uint8 {
	if a < b {
		return a
	}
	return b
}

// MinUint16 returns minimum between two uint16
func MinUint16(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

// MinUint32 returns minimum between two uint32
func MinUint32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

// MinUint64 returns minimum between two uint64
func MinUint64(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

// MaxByte returns maximum between two byte
func MaxByte(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

// MaxInt returns maximum between two int
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxInt8 returns maximum between two int8
func MaxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

// MaxInt16 returns maximum between two int16
func MaxInt16(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

// MaxInt32 returns maximum between two int32
func MaxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

// MaxInt64 returns maximum between two int64
func MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// MaxUint returns maximum between two uint
func MaxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

// MaxUint8 returns maximum between two uint8
func MaxUint8(a, b uint8) uint8 {
	if a > b {
		return a
	}
	return b
}

// MaxUint16 returns maximum between two uint16
func MaxUint16(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

// MaxUint32 returns maximum between two uint32
func MaxUint32(a, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

// MaxUint64 returns maximum between two uint64
func MaxUint64(a, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// AbsInt returns absolute value
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// AbsInt8 returns absolute value
func AbsInt8(a int8) int8 {
	if a < 0 {
		return -a
	}
	return a
}

// AbsInt16 returns absolute value
func AbsInt16(a int16) int16 {
	if a < 0 {
		return -a
	}
	return a
}

// AbsInt32 returns absolute value
func AbsInt32(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}

// AbsInt64 returns absolute value
func AbsInt64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

// ContainsInt returns true if e is in slice s, otherwise returns false
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsInt8 returns true if e is in slice s, otherwise returns false
func ContainsInt8(s []int8, e int8) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsInt16 returns true if e is in slice s, otherwise returns false
func ContainsInt16(s []int16, e int16) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsInt32 returns true if e is in slice s, otherwise returns false
func ContainsInt32(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsInt64 returns true if e is in slice s, otherwise returns false
func ContainsInt64(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsUint returns true if e is in slice s, otherwise returns false
func ContainsUint(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsUint8 returns true if e is in slice s, otherwise returns false
func ContainsUint8(s []uint8, e uint8) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsUint16 returns true if e is in slice s, otherwise returns false
func ContainsUint16(s []uint16, e uint16) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsUint32 returns true if e is in slice s, otherwise returns false
func ContainsUint32(s []uint32, e uint32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsUint64 returns true if e is in slice s, otherwise returns false
func ContainsUint64(s []uint64, e uint64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
