// Code by Steaz, June 2014
// Deq implements a double-ended queue data structure
// Main learning experience: empty interfaces to give 'generic' nature to data

package deq

import (
	"fmt"
)

// Dequeue is our main structure, Go inits all values to nil, nil, 0 automatically
type Dequeue struct {
	head  *Node
	tail  *Node
	count int
}

// Node is an entity of the Dequeue
type Node struct {
	item interface{}
	next *Node
}

// NewDequeue creates a new Dequeue, convenience wrapper
func NewDequeue() *Dequeue {
	return &Dequeue{}
}

// String allows Dequeue to be a Stringer
func (d *Dequeue) String() string {
	if d.Size() == 0 {
		return fmt.Sprintf("{}")
	}

	if d.Size() == 1 {
		return fmt.Sprintf("{%v}", d.head.item)
	}

	s := "{"
	runner := Node{nil, d.head}

	for runner.next != nil {
		// if last item, skip comma
		if runner.next.next == nil {
			s = fmt.Sprintf("%s%v", s, runner.next.item)
		} else {
			s = fmt.Sprintf("%s%v,", s, runner.next.item)
		}

		runner = *runner.next
	}
	s = s + "}"
	return s
}

// HeadAdd adds an item to the head of the Dequeue
func (d *Dequeue) HeadAdd(o interface{}) error {
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
	d.count++
	return nil
}

// TailAdd adds an item to the Tail of the Dequeue
func (d *Dequeue) TailAdd(o interface{}) error {
	n := &Node{o, nil}

	if d.count == 0 {
		d.head, d.tail = n, n
		d.count++
		return nil
	}

	d.count++
	d.tail.next, d.tail = n, n
	return nil
}

// HeadPeek returns the item at the Head without removal
func (d *Dequeue) HeadPeek() (interface{}, error) {
	if d.count == 0 {
		return nil, underflow()
	}

	return d.head.item, nil
}

// TailPeek returns the item at the Tail without removal
func (d *Dequeue) TailPeek() (interface{}, error) {
	if d.count == 0 {
		return nil, underflow()
	}

	return d.tail.item, nil
}

// HeadRemove removes the item at the head from the Dequeue
func (d *Dequeue) HeadRemove() (interface{}, error) {

	if d.count == 0 {
		return nil, underflow()
	}

	item := d.head.item

	if d.count == 1 {
		d.head, d.tail = nil, nil
		d.count--
		return item, nil
	}

	d.head = d.head.next
	d.count--
	return item, nil
}

// TailRemove removes the item at the tail from the Dequeue
func (d *Dequeue) TailRemove() (interface{}, error) {
	if d.count == 0 {
		return nil, underflow()
	}

	item := d.tail.item

	if d.count == 1 {
		d.head, d.tail = nil, nil
		d.count--
		return item, nil
	}

	// need to find the new tail
	newTail := &Node{nil, d.head}
	for newTail.next != d.tail {
		newTail = newTail.next
	}

	// nillify the old tail
	newTail.next = nil

	// set tail to newTail, decrement, and return
	d.tail = newTail
	d.count--
	return item, nil
}

// underflow wraps an error for removal from empty Dequeue
func underflow() error {
	return fmt.Errorf("Cannot peek/remove from empty Dequeue")
}

// Size returns number of objects in Dequeue, need to test
func (d *Dequeue) Size() int {
	return d.count
}

// Empty returns if Dequeue is empty in Go style (not isEmpty), need to test
func (d *Dequeue) Empty() bool {
	return d.count == 0
}
