package converter

import (
	"reflect"
	"testing"
)

const testString = "asdadadqqwdqwvefwvewvewwewq   eq tw wqwt twertewrtwwrwerrcq  rg eheberyeeybwcevfenhbebewew bjwhb wub hbqj bkjh vsakjhdb we b v wjhvwejfv e wefwb jhvw f"

var testSlice = []byte(testString)

func BenchmarkStringCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = []byte(testString)
	}
}

func BenchmarkStringSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = String2ByteSlice(testString)
	}
}

func BenchmarkByteCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(testSlice)
	}
}

func BenchmarkByteSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ByteSlice2String(testSlice)
	}
}

func TestString2ByteSlice(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{name: "asdsadasdasd", args: args{s: "asdsadasdasd"}, want: []byte("asdsadasdasd")},
		{name: "test string", args: args{s: testString}, want: []byte(testString)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String2ByteSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String2ByteSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteSlice2String(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "asdsadasdasd", args: args{b: []byte("asdsadasdasd")}, want: "asdsadasdasd"},
		{name: "test string", args: args{b: []byte(testString)}, want: testString},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteSlice2String(tt.args.b); got != tt.want {
				t.Errorf("ByteSlice2String() = %v, want %v", got, tt.want)
			}
		})
	}
}
