some data structures in go!
==========

#### bitset.go: 

[![GoDoc](https://godoc.org/github.com/steaz/datastruct/bitset?status.png)](https://godoc.org/github.com/steaz/datastruct/bitset)

implementation of bitset in go (clearly), but what's more important is this is my first true foray into the realm of TDD (test driven development).  I'm writing a test for each method, ensuring it fails, then implementing the method to make it pass. A lot easier than I thought at first, and loving the testing package that comes stock with go. 

#### deq.go

[![GoDoc](https://godoc.org/github.com/steaz/datastruct/deq?status.png)](https://godoc.org/github.com/steaz/datastruct/deq)

implementation of a Dequeue in Go, best known as a double ended queue, pronounced "Deck" (right?). Further practicing some TDD and refreshing my LL traversal algs. 

#### linked.go: 

this is my implementation of a doubly linked list in go when I was first learning the language.  I did borrow a lot of the code from the standard library's linked list implementation, and treated it as an exercise.  I will, at some point in the future, come back and test this as I did for the bitset, may practice some more TDD and redo it entirely. 
