package bloom

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"sync"

	"github.com/imbuba/utils/data/structures/set"
	"github.com/imbuba/utils/hash"
	"github.com/imbuba/utils/math/ints"
)

// ErrEmptyKey error for empty key
var ErrEmptyKey = fmt.Errorf("empty key")

// MurmurBloom struct
type MurmurBloom struct {
	sync.RWMutex
	vectorSize int32
	nbHash     uint
	bits       *set.BitSet
}

// New returns new MurmurBloom
func New(vectorSize int32, nbHash uint) *MurmurBloom {
	return &MurmurBloom{
		vectorSize: vectorSize,
		nbHash:     nbHash,
		bits:       set.NewSticky(uint(vectorSize)),
	}
}

// Add do smth
func (m *MurmurBloom) Add(key []byte) error {
	if len(key) == 0 {
		return ErrEmptyKey
	}
	hashes := m.hash(key)
	m.Lock()
	for _, h := range hashes {
		m.bits.SetOne(h)
	}
	m.Unlock()
	return nil
}

// Clear do smth
func (m *MurmurBloom) Clear() {
	m.Lock()
	m.bits.Clear()
	m.Unlock()
}

// Contains do smth
func (m *MurmurBloom) Contains(key []byte) (bool, error) {
	if len(key) == 0 {
		return false, ErrEmptyKey
	}
	hashes := m.hash(key)
	result := true
	m.RLock()
	for _, h := range hashes {
		if !m.bits.GetOne(h) {
			result = false
			break
		}
	}
	m.RUnlock()
	return result, nil
}

// Dump do smth
func (m *MurmurBloom) Dump(writer io.Writer) error {
	m.RLock()
	defer m.RUnlock()
	if err := binary.Write(writer, binary.BigEndian, m.vectorSize); err != nil {
		return err
	}
	if err := binary.Write(writer, binary.BigEndian, int32(m.nbHash)); err != nil {
		return err
	}
	if err := m.bits.Dump(writer); err != nil {
		return err
	}
	return nil
}

// Bytes do smth
func (m *MurmurBloom) Bytes() []byte {
	buf := new(bytes.Buffer)
	m.Dump(buf)
	return buf.Bytes()
}

// hash do smth
func (m *MurmurBloom) hash(key []byte) []uint {
	result := make([]uint, m.nbHash)
	for i, initVal := uint(0), int32(0); i < m.nbHash; i++ {
		initVal = hash.Murmur(key, initVal)
		result[i] = uint(ints.AbsInt32(initVal % m.vectorSize))
	}
	return result
}

// CreateFromJava do smth
func CreateFromJava(stream io.Reader) (*MurmurBloom, error) {
	var ver int32
	if err := binary.Read(stream, binary.BigEndian, &ver); err != nil {
		return nil, err
	}
	if ver > 0 {
		return nil, fmt.Errorf("not supported version %d", ver)
	}
	if ver != -1 {
		return nil, fmt.Errorf("not supported version %d", ver)
	}
	var m, k uint32
	var ht byte
	if err := binary.Read(stream, binary.BigEndian, &k); err != nil {
		return nil, err
	}
	if err := binary.Read(stream, binary.BigEndian, &ht); err != nil {
		return nil, err
	}
	if ht != 1 {
		return nil, fmt.Errorf("not supported hash function")
	}
	if err := binary.Read(stream, binary.BigEndian, &m); err != nil {
		return nil, err
	}
	result := New(int32(m), uint(k))
	result.bits.ReadFrom(stream)
	return result, nil
}

// Create do smth
func Create(stream io.Reader) (*MurmurBloom, error) {
	var vSize, h int32
	if err := binary.Read(stream, binary.BigEndian, &vSize); err != nil {
		return nil, err
	}
	if err := binary.Read(stream, binary.BigEndian, &h); err != nil {
		return nil, err
	}
	result := New(vSize, uint(h))
	result.bits.ReadFrom(stream)
	return result, nil
}
