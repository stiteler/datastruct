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

// NewBitset makes a new bitset given a size (int)
func NewBitset(size int) (*Bitset, error) {
	if size < 0 {
		return nil, fmt.Errorf("invalid size")
	}

	num := (size + 7) / 8
	bytes := make([]byte, num)
	return &Bitset{size, bytes}, nil
}

// SetBit sets nth bit in b.Bytes to 1
func (b *Bitset) SetBit(n int) error {
	if n > b.Max {
		return fmt.Errorf("out of range")
	}
	_byte := n / 8
	_bit := uint32(n % 8)
	b.Bytes[_byte] |= (1 << _bit)
	return nil
}

// GetBit returns true if nth bit in b.Bytes is 1, false if 0
func (b *Bitset) GetBit(n int) (bool, error) {
	if n > b.Max {
		return false, fmt.Errorf("out of range")
	}
	_byte := n / 8
	_bit := uint32(n % 8)
	answer := ((b.Bytes[_byte] & (1 << _bit)) != 0)
	return answer, nil
}

// Clear clears the entire bitset
func (b *Bitset) Clear() error {
	if b == nil {
		return fmt.Errorf("Can't clear nil bitset")
	}
	for _byte := range b.Bytes {
		b.Bytes[_byte] = 0
	}
	return nil
}

// ClearBit clears nth bit from the bitset
func (b *Bitset) ClearBit(n int) error {
	if n > b.Max {
		return fmt.Errorf("out of range")
	}
	_byte := n / 8
	_bit := uint32(n % 8)
	b.Bytes[_byte] &= ((1 << _bit) ^ 255)
	return nil
}

// Cardinality returns count of '1' bits in bitset
func (b *Bitset) Cardinality() int {
	cardinality := 0
	for i := 0; i < b.Max; i++ {
		isSet, _ := b.GetBit(i)
		if isSet {
			cardinality++
		}
	}
	return cardinality
}

// IsSubset returns true if arg bitset is a subset of the given bitset
func (this *Bitset) IsSubset(that *Bitset) (bool, error) {
	// check first if impossible by cardinality:
	if this.Cardinality() > that.Cardinality() {
		return false, nil
	}

	// this error handling for getBit looks a little nasty
	// there's got to be a better way...
	for i := 0; i < this.Max; i++ {
		res, err := this.GetBit(i)
		if err != nil {
			return false, err
		}

		if res {
			res, err = that.GetBit(i)
			if err != nil {
				return false, err
			}
			if !res {
				return false, nil
			}
		}
	}
	return true, nil
}

// Contains returns true if given value is a member
func (b *Bitset) Contains(n int) bool {
	if n > b.Max {
		return false
	}

	contains, _ := b.GetBit(n)
	return contains
}

func (this *Bitset) Equals(that *Bitset) bool {

	return false
}

func (this *Bitset) Union(that *Bitset) *Bitset {

	return nil
}
