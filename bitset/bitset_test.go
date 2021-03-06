package bitset

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
	//todo: can replace call to bytes with call to Equals in package
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

func TestClear(t *testing.T) {
	bs := newBitsetOrFatal(t, 8)
	bs = nil
	if err := bs.Clear(); err == nil {
		t.Errorf("Expected clearing nil bitset error")
	}

	bs = newBitsetOrFatal(t, 8)
	bs.SetBit(1)
	bs.SetBit(3)
	bs.SetBit(7)
	bs.Clear()
	for _byte := range bs.Bytes {
		if _byte != 0 {
			t.Errorf("Expected bytes == 0 after b.Clear()")
		}
	}
}

func TestClearBit(t *testing.T) {
	bs := newBitsetOrFatal(t, 8)
	if err := bs.SetBit(9); err == nil {
		t.Errorf("Expected ClearBit out of range error")
	}
	bs.SetBit(4)
	bs.ClearBit(4)
	if ok, _ := bs.GetBit(4); ok {
		t.Errorf("Expected GetBit after clear to be false")
	}
}

func TestCardinality(t *testing.T) {
	bs := newBitsetOrFatal(t, 8)
	bs.SetBit(1)
	bs.SetBit(2)
	bs.SetBit(3)
	if bs.Cardinality() != 3 {
		t.Errorf("Expected count of 3, got: %v", bs.Cardinality())
	}
}

func TestCopy(t *testing.T) {
	bs1 := newBitsetOrFatal(t, 16)
	bs1.SetBit(14)
	bs1.SetBit(4)

	bs2 := bs1.Copy()
	equal := bs2.Equals(bs1)
	if !equal {
		t.Errorf("Copy failure, bs2: %v", bs2)
	}
}

func TestIsSubset(t *testing.T) {
	bs1 := newBitsetOrFatal(t, 8)
	bs2 := newBitsetOrFatal(t, 8)
	bs1.SetBit(1)
	bs1.SetBit(2)
	bs1.SetBit(5)
	// bs2 should be a subset of bs1
	bs2.SetBit(1)
	bs2.SetBit(5)

	answer, err := bs2.IsSubset(bs1)
	if err != nil {
		t.Fatalf("unexpected error testing isSubset: %v", err)
	}
	if answer != true {
		t.Errorf("IsSubset test failure")
	}
}

func TestContains(t *testing.T) {
	bs := newBitsetOrFatal(t, 8)
	bs.SetBit(5)
	test := bs.Contains(9)
	if test != false {
		t.Errorf("calling Contains with n > Bitset.Max must return false")
	}

	test = bs.Contains(6)
	if test != false {
		t.Errorf("Expected Contains to return false, got: %v", test)
	}

	test = bs.Contains(5)
	if test != true {
		t.Errorf("Expected Contains to return true, got: %v", test)
	}
}

func TestEquals(t *testing.T) {
	bs1 := newBitsetOrFatal(t, 8)
	bs2 := newBitsetOrFatal(t, 16)

	bs1.SetBit(1)
	bs1.SetBit(5)
	bs1.SetBit(7)

	bs2.SetBit(1)
	bs2.SetBit(5)
	bs2.SetBit(7)

	test := bs1.Equals(bs2)
	if test != true {
		t.Errorf("Expected Equals to return true, got: %v", test)
	}

	bs2.SetBit(11)
	test = bs1.Equals(bs2)
	if test != false {
		t.Errorf("Expected Unequal Equals to return false, got: %v", test)
	}
}

func TestUnion(t *testing.T) {
	bs1 := newBitsetOrFatal(t, 16)
	bs2 := newBitsetOrFatal(t, 8)

	// set up bitsets
	bs1.SetBit(2)
	bs1.SetBit(13)
	bs1.SetBit(11)

	bs2.SetBit(5)
	bs2.SetBit(3)
	bs2.SetBit(7)

	bs3 := bs1.Union(bs2)
	if bs3 == nil {
		t.Fatalf("union bitset returned nil")
	}

	bs4 := newBitsetOrFatal(t, 16)
	bs4.SetBit(2)
	bs4.SetBit(3)
	bs4.SetBit(5)
	bs4.SetBit(7)
	bs4.SetBit(11)
	bs4.SetBit(13)

	equal := bs3.Equals(bs4)
	if !equal {
		t.Errorf("Expected bs3 to equal bs4 after union, got: %v\n bs3: %v\n bs4: %v\n", equal, bs3, bs4)
	}
}

func TestDifference(t *testing.T) {
	bs1 := newBitsetOrFatal(t, 16)
	bs2 := newBitsetOrFatal(t, 8)

	// set up bitsets
	bs1.SetBit(1)
	bs1.SetBit(2)
	bs1.SetBit(3)

	bs2.SetBit(1)
	bs2.SetBit(2)
	bs2.SetBit(4)

	bs3 := bs1.Difference(bs2)
	if bs3 == nil {
		t.Fatalf("difference bitset returned nil")
	}

	bs4 := newBitsetOrFatal(t, 16)
	bs4.SetBit(3)

	equal := bs4.Equals(bs3)
	if !equal {
		t.Errorf("Expected bs3 to equal bs4 after difference, got %v\n bs3: %v\n bs4: %v\n", equal, bs3, bs4)
	}
}

func TestIntersect(t *testing.T) {
	bs1 := newBitsetOrFatal(t, 16)
	bs2 := newBitsetOrFatal(t, 8)

	// set up bitsets
	bs1.SetBit(1)
	bs1.SetBit(2)
	bs1.SetBit(3)

	bs2.SetBit(1)
	bs2.SetBit(5)
	bs2.SetBit(4)

	bs3 := bs1.Intersect(bs2)
	if bs3 == nil {
		t.Fatalf("intersect returned nil")
	}

	bs4 := newBitsetOrFatal(t, 16)
	bs4.SetBit(1)

	equal := bs4.Equals(bs3)
	if !equal {
		t.Errorf("Expected bs3 to equal bs4 after intersect, got %v\n bs3: %v\n bs4: %v\n", equal, bs3, bs4)
	}
}
