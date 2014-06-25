// Code by Steaz, June 2014
// Testing Deq.go (Dequeue Data Structure)

package deq

import (
	"testing"
)

func TestNewDequeue(t *testing.T) {
	spawn := NewDequeue()
	if spawn == nil {
		t.Fatalf("spawned nil Dequeue on NewDequeue()")
	}
}

func TestSize(t *testing.T) {
	d := NewDequeue()
	if d.Size() != 0 {
		t.Errorf("Expected NewDequeue size == 0, got: %v", d.Size())
	}

	d.TailAdd(4)
	if d.Size() != 1 {
		t.Errorf("Expected Dequeue to have size of 1, got: %v", d.Size())
	}
}

func TestEmpty(t *testing.T) {
	d := NewDequeue()
	if !d.Empty() {
		t.Errorf("Expected NewDequeue Empty() to return true got: %v", d.Empty())
	}
	d.HeadAdd(4)
	if d.Empty() {
		t.Errorf("Expected Dequeue with 1 item for Empty() to return false, got: %v", d.Empty())
	}
}

func TestHeadAdd(t *testing.T) {
	d := NewDequeue()
	if err := d.HeadAdd(5); err != nil {
		t.Errorf("HeadAdd Error: %v", err)
	}

	if d.Size() != 1 {
		t.Errorf("Expected HeadAdd() to result in size of 1, got: ", d.Size())
	}
}

func TestHeadPeek(t *testing.T) {
	d := NewDequeue()

	under, err := d.HeadPeek()
	if err == nil {
		t.Errorf("HeadPeek underflow error expected, got: %v", err)
	}
	if under != nil {
		t.Errorf("Expected underflow *Object to be nil")
	}

	d.HeadAdd(5)

	obj, err := d.HeadPeek()
	if err != nil {
		t.Errorf("Unexpected HeadPeek() error: %v", err)
	}
	if *obj != 5 {
		t.Errorf("Expected HeadPeek() to return 5, got: %v", obj)
	}

	obj2, err := d.HeadPeek()
	if obj2 != obj {
		t.Errorf("Expected simultaneous calls to HeadPeek() to return equal objects, got %v, %v", obj, obj2)
	}
}

func TestTailAdd(t *testing.T) {
	d := NewDequeue()
	if err := d.TailAdd(5); err != nil {
		t.Errorf("TailAdd Error: %v", err)
	}

	if d.Size() != 1 {
		t.Errorf("Expected TailAdd() to result in size of 1, got: ", d.Size())
	}
	d.TailAdd(2)
	d.TailAdd(3)
	d.TailAdd(4)
	d.TailAdd(5)
	if d.Size() != 5 {
		t.Errorf("Expected TailAdd() after 4 additions to result in size of 4, got: %v", d.Size())
	}

}

func TestTailPeek(t *testing.T) {
	d := NewDequeue()

	under, err := d.TailPeek()
	if err == nil {
		t.Errorf("HeadPeek underflow error expected, got: %v", err)
	}
	if under != nil {
		t.Errorf("Expected underflow *Object to be nil")
	}

	d.TailAdd(5)

	obj, err := d.TailPeek()
	if err != nil {
		t.Errorf("Unexpected HeadPeek() error: %v", err)
	}
	if *obj != 5 {
		t.Errorf("Expected HeadPeek() to return 5, got: %v", obj)
	}

	obj2, err := d.TailPeek()
	if obj2 != obj {
		t.Errorf("Expected simultaneous calls to HeadPeek() to return equal objects, got %v, %v", obj, obj2)
	}
}

func TestHeadRemove(t *testing.T) {
	d := NewDequeue()

	u, err := d.HeadRemove()
	if err == nil {
		t.Errorf("Expected HeadRemove to return Underflow error if Dequeue Empty")
	}
	if u != nil {
		t.Errorf("Expected HeadRemove to return nil object if Dequeue Empty")
	}

	d.HeadAdd(5)
	d.HeadAdd(89)

	item, err := d.HeadRemove()
	if d.Size() != 1 {
		t.Errorf("Expected HeadRemove() to result in a size of 1, got: %v", d.Size())
	}
	if *item != Object(89) {
		t.Errorf("Expected HeadRemove() to return 89, got: ", item)
	}

	item, err = d.HeadRemove()
	if d.Size() != 0 {
		t.Errorf("Expected HeadRemove() to result in a size of 0, got: %v", d.Size())
	}
	if *item != Object(5) {
		t.Errorf("Expected HeadRemove() to return 5, got: ", item)
	}
}

func TestTailRemove(t *testing.T) {
	d := NewDequeue()

	u, err := d.TailRemove()
	if err == nil {
		t.Errorf("Expected TailRemove to return Underflow error if Dequeue Empty")
	}
	if u != nil {
		t.Errorf("Expected TailRemove to return nil object if Dequeue Empty")
	}

	d.TailAdd(13)
	d.TailAdd(45)

	item, err := d.TailRemove()
	if d.Size() != 1 {
		t.Errorf("Expected TailRemove() to result in a size of 1, got: %v", d.Size())
	}
	if *item != Object(13) {
		t.Errorf("Expected TailRemove() to return 13, got: ", item)
	}

	item, err = d.TailRemove()
	if d.Size() != 0 {
		t.Errorf("Expected TailRemove() to result in a size of 0, got: %v", d.Size())
	}
	if *item != Object(45) {
		t.Errorf("Expected TailRemove() to return 45, got: ", item)
	}
}
