package main

import (
	"fmt"
)

type Tree struct {
	Root  *Node
	Count int
}

type Node struct {
	Data   int
	Lchild *Node
	Rchild *Node
}

func main() {
	t := &Tree{}
	t.Add(5)
	t.Add(7)
	t.Add(1)
	t.Add(3)
	t.Print()
}

func (t *Tree) Add(data int) bool {
	// if empty
	if t.Root == nil {
		t.Root = &Node{Data: data}
		t.Count++
		return true
	}

	// place in order
	runner := t.Root
	for runner != nil {
		// if data belongs to left
		if data < runner.Data {
			// if belongs on this node.
			if runner.Lchild == nil {
				runner.Lchild = &Node{Data: data}
				t.Count++
				return true
			}
			runner = runner.Lchild

		} else {
			if runner.Rchild == nil {
				runner.Rchild = &Node{Data: data}
				t.Count++
				return true
			}
			runner = runner.Rchild
		}
	}
	return false
}

func (t Tree) Print() {
	t.Root.treePrint()
	fmt.Println()
}

func (n *Node) treePrint() {
	if n == nil {
		fmt.Print(".")
	} else {
		n.Lchild.treePrint()
		fmt.Print(n.Data)
		n.Rchild.treePrint()
	}
}
