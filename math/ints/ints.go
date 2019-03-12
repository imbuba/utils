package ints

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func MinInt32(a int32, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func MinInt64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func MinUint(a uint, b uint) uint {
	if a < b {
		return a
	}
	return b
}

func MinUint32(a uint32, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

func MinUint64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MaxInt32(a int32, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func MaxInt64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MaxUint(a uint, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func MaxUint32(a uint32, b uint32) uint32 {
	if a > b {
		return a
	}
	return b
}

func MaxUint64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func AbsInt32(a int32) int32 {
	if a < 0 {
		return -a
	}
	return a
}

func AbsInt64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsInt32(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsInt64(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsUint(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsUint32(s []uint32, e uint32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsUint64(s []uint64, e uint64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
