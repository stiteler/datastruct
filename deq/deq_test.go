package deq

import (
	"fmt"
	"testing"
)

func newDeqOrFatal(t *testing.T) *Dequeue {
	spawn := NewDequeue()
	if spawn == nil {
		t.Errorf("Received nil Dequeue from NewDequeue()")
	}
	fmt.Println("Spawn: ", spawn)
	return spawn
}
