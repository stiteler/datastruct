package datastruct

import (
	"bytes"
	"testing"
)

func newBitsetOrFatal(t *testing.T, size int) *Bitset {
	bs, err := NewBitset(size)
	if err != nil {
		t.Fatalf("new bitset: %v", err)
	}
	return bs
}

func TestBitsetSize(t *testing.T) {
	bs := newBitsetOrFatal(t, 8)
	if len(bs.Bytes) != 1 {
		t.Errorf("wrong size")
	}

	bs = newBitsetOrFatal(t, 9)
	if len(bs.Bytes) != 2 {
		t.Errorf("wrong size")
	}
}

func TestSetBit(t *testing.T) {
	bs := newBitsetOrFatal(t, 8)
	err := bs.SetBit(9)
	if err == nil {
		t.Errorf("can't set out of bitset range")
	}

	bs = newBitsetOrFatal(t, 8)
	// set bit 0, which is congruent to the integer
	err = bs.SetBit(0)
	if err != nil {
		t.Errorf("Failure to set bit")
	}
	expectedInt := 1
	expected := []byte{byte(expectedInt)}
	if !bytes.Equal(expected, bs.Bytes) {
		t.Errorf("bit not set correctly")
	}

	// todo: add tests for []byte > 1 byte
}

func TestGetBit(t *testing.T) {
	bs := newBitsetOrFatal(t, 8)
	bs.SetBit(3)

	if _, err := bs.GetBit(9); err == nil {
		t.Errorf("Expected GetBit out of range error")
	}

	if expFalse, _ := bs.GetBit(4); expFalse {
		t.Errorf("GetBit expected false, got: %v", expFalse)
	}

	if expTrue, _ := bs.GetBit(3); !expTrue {
		t.Errorf("GetBit expected true, got; %v", expTrue)
	}

}
