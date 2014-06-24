package deq

import (
	"fmt"
)

// Dequeue is our main structure, Go inits all values to nil, nil, 0 automatically
type Dequeue struct {
	head  *Node
	tail  *Node
	count int32
}

// Payload is a type I'm defining for ease of flexiblity later
type Object int32

type Node struct {
	// what should the item be? an empty interface? for now, int
	item Object
	next *Node
}

// NewDequeue creates a new Dequeue, I even need this?
func NewDequeue() *Dequeue {
	return &Dequeue{}
}

// Stubs:

// String allows Dequeue to be a Stringer
func (d *Dequeue) String() string {
	return fmt.Sprintf("Test String")
}

// HeadAdd ...
func (d *Dequeue) HeadAdd(o Object) error {
	// what do I need to clean on HeadAdd?

	// create new node, n
	n := &Node{o, nil}

	// if this is a new dequeue, just pop the new node on
	if d.count == 0 {
		d.head, d.tail = n, n
		d.count++
		return nil
	}

	// othewise point next head, then point head at new
	n.next, d.head = d.head, n
	return nil
}

// TailAdd ...
// func (d *Dequeue) TailAdd() {}

// HeadPeek ...
// func (d *Dequeue) HeadPeek() Object {}

// TailPeek ...
// func (d *Dequeue) TailPeek() Object {}

// HeadRemove ...
// func (d *Dequeue) HeadRemove() Object {}

// TailRemove ...
// func (d *Dequeue) TailRemove() Object {}

// Size returns number of objects in Dequeue, need to test
func (d *Dequeue) Size() int32 {
	return d.count
}

// Empty returns if Dequeue is empty in Go style (not isEmpty), need to test
func (d *Dequeue) Empty() bool {
	return d.count == 0
}
