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

func TestHeadAdd(t *testing.T) {
	d := NewDequeue()
	if err := d.HeadAdd(5); err != nil {
		t.Errorf("HeadAdd Error: %v", err)
	}
}
