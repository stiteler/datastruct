package deq

// Dequeue is our main structure, Go inits all values to nil, nil, 0 automatically
type Dequeue struct {
	Head  *Node
	Tail  *Node
	Count int32
}

// Payload is a type I'm defining for ease of flexiblity later
type Object int32

type Node struct {
	// what should the item be? an empty interface? for now, int
	Item Object
	Next *Node
}

// NewDequeue creates a new Dequeue, I even need this?
func NewDequeue() *Dequeue {
	return &Dequeue{}
}

// Stubs:

// String allows Dequeue to be a Stringer
// func (d *Dequeue) String() string {}

// HeadAdd ...
// func (d *Dequeue) HeadAdd() {}

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
	return d.Count
}

// Empty returns if Dequeue is empty in Go style (not isEmpty), need to test
func (d *Dequeue) Empty() bool {
	return d.Count == 0
}
