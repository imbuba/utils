package ints

import (
	"testing"
)

func TestMinByte(t *testing.T) {
	type args struct {
		a byte
		b byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinByte(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt8(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt8(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt16(t *testing.T) {
	type args struct {
		a int16
		b int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt16(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt32(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt32(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt64(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt64(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint(t *testing.T) {
	type args struct {
		a uint
		b uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint8(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint8(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint16(t *testing.T) {
	type args struct {
		a uint16
		b uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint16(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint32(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint32(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinUint64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "first less", args: args{a: 1, b: 10}, want: 1},
		{name: "both equals", args: args{a: 1, b: 1}, want: 1},
		{name: "second less", args: args{a: 10, b: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinUint64(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MinUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxByte(t *testing.T) {
	type args struct {
		a byte
		b byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxByte(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt8(t *testing.T) {
	type args struct {
		a int8
		b int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt8(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt16(t *testing.T) {
	type args struct {
		a int16
		b int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt16(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt32(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt32(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInt64(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt64(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint(t *testing.T) {
	type args struct {
		a uint
		b uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint8(t *testing.T) {
	type args struct {
		a uint8
		b uint8
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint8(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint16(t *testing.T) {
	type args struct {
		a uint16
		b uint16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint16(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint32(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint32(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "first greater", args: args{a: 10, b: 1}, want: 10},
		{name: "both equals", args: args{a: 10, b: 10}, want: 10},
		{name: "second greater", args: args{a: 1, b: 10}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint64(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("MaxUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "positive", args: args{a: 10}, want: 10},
		{name: "negative", args: args{a: -10}, want: 10},
		{name: "zero", args: args{a: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt(tt.args.a); got != tt.want {
				t.Errorf("AbsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt8(t *testing.T) {
	type args struct {
		a int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{name: "positive", args: args{a: 10}, want: 10},
		{name: "negative", args: args{a: -10}, want: 10},
		{name: "zero", args: args{a: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt8(tt.args.a); got != tt.want {
				t.Errorf("AbsInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt16(t *testing.T) {
	type args struct {
		a int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{name: "positive", args: args{a: 10}, want: 10},
		{name: "negative", args: args{a: -10}, want: 10},
		{name: "zero", args: args{a: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt16(tt.args.a); got != tt.want {
				t.Errorf("AbsInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt32(t *testing.T) {
	type args struct {
		a int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "positive", args: args{a: 10}, want: 10},
		{name: "negative", args: args{a: -10}, want: 10},
		{name: "zero", args: args{a: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt32(tt.args.a); got != tt.want {
				t.Errorf("AbsInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbsInt64(t *testing.T) {
	type args struct {
		a int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "positive", args: args{a: 10}, want: 10},
		{name: "negative", args: args{a: -10}, want: 10},
		{name: "zero", args: args{a: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AbsInt64(tt.args.a); got != tt.want {
				t.Errorf("AbsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	type args struct {
		s []int
		e int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt8(t *testing.T) {
	type args struct {
		s []int8
		e int8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt8(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt16(t *testing.T) {
	type args struct {
		s []int16
		e int16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt16(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt32(t *testing.T) {
	type args struct {
		s []int32
		e int32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt32(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt64(t *testing.T) {
	type args struct {
		s []int64
		e int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt64(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint(t *testing.T) {
	type args struct {
		s []uint
		e uint
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint8(t *testing.T) {
	type args struct {
		s []uint8
		e uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint8(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint16(t *testing.T) {
	type args struct {
		s []uint16
		e uint16
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint16(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint32(t *testing.T) {
	type args struct {
		s []uint32
		e uint32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint32(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsUint64(t *testing.T) {
	type args struct {
		s []uint64
		e uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "empty slice", args: args{e: 10}, want: false},
		{name: "has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 10, 12)}, want: true},
		{name: "not has", args: args{e: 10, s: append(args{}.s, 0, 2, 4, 6, 8, 12)}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsUint64(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
