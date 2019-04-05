package set

import (
	"encoding/binary"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/finnan444/utils/math/ints"
)

const (
	addressBitsPerWord = 6
	bitsPerWord        = 1 << addressBitsPerWord
	one                = int64(1)
)

var (
	// ErrOutOfBounds error about index out of bounds
	ErrOutOfBounds = fmt.Errorf("index out of bounds")
)

// BitSet struct
type BitSet struct {
	words      []int64
	wordsInUse int64
	nBits      uint
}

func wordIndex(bitIndex uint) uint {
	return bitIndex >> addressBitsPerWord
}

// New returns new BitSet
func New() *BitSet {
	return &BitSet{
		words: make([]int64, wordIndex(bitsPerWord-1)+1),
	}
}

// NewSticky do smth
func NewSticky(nBits uint) *BitSet {
	return &BitSet{
		words: make([]int64, wordIndex(nBits-1)+1),
		nBits: nBits,
	}
}

// recalculateWordsInUse Sets the field wordsInUse to the logical size in words of the bit set.
// WARNING:This method assumes that the number of words actually in use is
// less than or equal to the current value of wordsInUse!
func (m *BitSet) recalculateWordsInUse() {
	var i int64
	for i = m.wordsInUse - 1; i >= 0; i-- {
		if m.words[i] != 0 {
			break
		}
	}
	m.wordsInUse = i + 1
}

// ensureCapacity Ensures that the BitSet can hold enough words.
// @param wordsRequired the minimum acceptable number of words.
func (m *BitSet) ensureCapacity(wordsRequired int64) {
	if int64(len(m.words)) < wordsRequired {
		newWords := make([]int64, ints.MaxInt64(wordsRequired, 2*int64(len(m.words))))
		copy(newWords, m.words)
		m.words = newWords
		m.nBits = 0
	}
}

// expandTo Ensures that the BitSet can accommodate a given wordIndex,
// temporarily violating the invariants.  The caller must
// restore the invariants before returning to the user,
// possibly using recalculateWordsInUse().
// @param wordIndex the index to be accommodated.
func (m *BitSet) expandTo(wordIndex uint) {
	wordsRequired := int64(wordIndex + 1)
	if m.wordsInUse < wordsRequired {
		m.ensureCapacity(wordsRequired)
		atomic.StoreInt64(&m.wordsInUse, wordsRequired)
	}
}

// FlipOne Sets the bit at the specified index to the complement of its
// current value.
// @param  bitIndex the index of the bit to flip
// @return ErrOutOfBounds if the specified index is negative
func (m *BitSet) FlipOne(bitIndex uint) {
	wIndex := wordIndex(bitIndex)
	m.expandTo(wIndex)
	m.words[wIndex] ^= (one << (bitIndex % 64))
	m.recalculateWordsInUse()
}

// SetOne do smth
func (m *BitSet) SetOne(bitIndex uint) {
	m.set(bitIndex)
}

// set do smth
func (m *BitSet) set(bitIndex uint) {
	wIndex := wordIndex(bitIndex)
	m.expandTo(wIndex)
	m.words[wIndex] |= (one << (bitIndex % 64))
}

// SetValue do smth
func (m *BitSet) SetValue(bitIndex uint, value bool) {
	if value {
		m.SetOne(bitIndex)
	} else {
		m.ClearOne(bitIndex)
	}
}

// ClearOne do smth
func (m *BitSet) ClearOne(bitIndex uint) {
	wIndex := wordIndex(bitIndex)
	if wIndex >= uint(m.wordsInUse) {
		return
	}
	m.words[wIndex] &= ^(one << (bitIndex % 64))
	m.recalculateWordsInUse()
}

// Clear do smth
func (m *BitSet) Clear() {
	for m.wordsInUse >= 0 {
		m.words[m.wordsInUse] = 0
		m.wordsInUse--
	}
}

// GetOne do smth
func (m *BitSet) GetOne(bitIndex uint) bool {
	wIndex := wordIndex(bitIndex)
	result := (wIndex < uint(m.wordsInUse)) && ((m.words[wIndex] & (one << (bitIndex % 64))) != 0)
	return result
}

// ReadFrom do smth
func (m *BitSet) ReadFrom(reader io.Reader) {
	buf := make([]byte, 1)
	ones := []byte{1, 2, 4, 8, 16, 32, 64, 128}
	var err error
	i := uint(0)
	for i < m.nBits {
		_, err = reader.Read(buf)
		if err != nil {
			break
		} else {
			val := buf[0]
			for _, o := range ones {
				if val&o != 0 {
					m.set(i)
				}
				i++
			}
		}
	}
}

// Dump do smth
func (m *BitSet) Dump(writer io.Writer) error {
	for _, w := range m.words {
		if err := binary.Write(writer, binary.LittleEndian, w); err != nil {
			return err
		}
	}
	return nil
}
