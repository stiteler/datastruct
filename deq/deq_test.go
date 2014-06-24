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