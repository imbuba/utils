package data

import (
	"testing"
)

func TestComparableInt(t *testing.T) {
	a := Int(1)
	b := Int(2)
	d := Int(2)

	if a.CompareTo(b) >= 0 {
		t.Errorf("Expecting %v is less than %v", a, b)
	}
	if b.CompareTo(a) <= 0 {
		t.Errorf("Expecting %v is greater than %v", b, a)
	}
	if b.CompareTo(d) != 0 {
		t.Errorf("Expecting %v is equals %v", b, d)
	}
	if a.Equals(b) {
		t.Errorf("Expecting %v is not equal %v", a, b)
	}
	if !d.Equals(b) {
		t.Errorf("Expecting %v is equal %v", d, b)
	}
}

func TestComparableString(t *testing.T) {
	a := String("a")
	b := String("b")
	d := String("b")

	if a.CompareTo(b) >= 0 {
		t.Errorf("Expecting %v is less than %v", a, b)
	}
	if b.CompareTo(a) <= 0 {
		t.Errorf("Expecting %v is greater than %v", b, a)
	}
	if b.CompareTo(d) != 0 {
		t.Errorf("Expecting %v is equals %v", b, d)
	}
	if a.Equals(b) {
		t.Errorf("Expecting %v is not equal %v", a, b)
	}
	if !d.Equals(b) {
		t.Errorf("Expecting %v is equal %v", d, b)
	}
}
