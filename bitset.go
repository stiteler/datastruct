package datastruct

// bitset data structure in go
// by Steaz
import (
	"fmt"
)

type Bitset struct {
	Max   int
	Bytes []byte
}

// make a new bitset
func NewBitset(size int) (*Bitset, error) {
	if size < 0 {
		return nil, fmt.Errorf("invalid size")
	}

	num := (size + 7) / 8
	bytes := make([]byte, num)
	return &Bitset{size, bytes}, nil
}

// sets nth bit in b.Bytes to 1
func (b *Bitset) SetBit(n int) error {
	if n > b.Max {
		return fmt.Errorf("out of range")
	}
	_byte := n / 8
	_bit := uint32(n % 8)
	b.Bytes[_byte] |= (1 << _bit)
	return nil
}

// returns true if nth bit in b.Bytes is 1, false if 0
func (b *Bitset) GetBit(n int) (bool, error) {
	if n > b.Max {
		return false, fmt.Errorf("out of range")
	}
	_byte := n / 8
	_bit := uint32(n % 8)
	answer := ((b.Bytes[_byte] & (1 << _bit)) != 0)
	return answer, nil
}

// clears nth bit from bitset
func (b *Bitset) ClearBit(n int) error {
	if n > b.Max {
		return fmt.Errorf("out of range")
	}
	_byte := n / 8
	_bit := uint32(n % 8)
	b.Bytes[_byte] &= ((1 << _bit) ^ 255)
	return nil
}

func (b *Bitset) Clear() error {
	if b == nil {
		return fmt.Errorf("Can't clear nil bitset")
	}
	for _byte := range b.Bytes {
		b.Bytes[_byte] = 0
	}
	return nil
}
