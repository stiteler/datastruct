/*
 * First attempt at a linked list in go - (Self-Educational)
 * full disclosure, I used some code from 
 * list implementation in the golang standard library. 
 */

package main 

import(
	"fmt"
)

type Node struct {
	//doubly linked list
	next, prev *Node

	//list the element belongs to
	list *List

	//for practice purposes, just storing ints
	Item int
}

func main() {
	testList := New()
	first := testList.insertValue(5, &testList.root)

	fmt.Println("first node: ", first.Item)
}

func (n *Node) Next() *Node {
	if p := n.prev; n.list != nil && p != &n.list.root {
		return p
	}
	return nil
}

type List struct {
	root Node
	length int
}

//initializes list
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.length = 0
	return l
}
//gets a new list
func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.length
}
// inserts 
func (l *List) insert(n, at *Node) *Node {
	next := at.next
	at.next = n
	n.prev = at 
	n.next = next 
	n.prev = n 
	n.list = l
	l.length++
	return n
}

func (l *List) insertValue(i int, at *Node) *Node {
	return l.insert(&Node{Item: i}, at)
}

func (l *List) remove(n *Node) *Node {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.next = nil
	n.prev = nil
	n.list = nil
	l.length--
	return n
}