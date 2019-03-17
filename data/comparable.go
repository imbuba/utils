package data

import (
	"strings"
)

// Comparable interface
type Comparable interface {
	CompareTo(other Comparable) int
	Equals(other Comparable) bool
	Default() Comparable
}

// Int wrapper
type Int int

// CompareTo impl Comparable
func (v Int) CompareTo(other Comparable) int {
	return int(v - other.(Int))
}

// Equals impl Comparable
func (v Int) Equals(other Comparable) bool {
	return v == other.(Int)
}

// Default impl Comparable
func (v Int) Default() Comparable {
	return Int(0)
}

// String wrapper
type String string

// CompareTo impl Comparable
func (v String) CompareTo(other Comparable) int {
	return strings.Compare(string(v), string(other.(String)))
}

// Equals impl Comparable
func (v String) Equals(other Comparable) bool {
	return v == other.(String)
}

// Default impl Comparable
func (v String) Default() Comparable {
	return String("")
}
