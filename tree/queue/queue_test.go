package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {

	q := New()

	fmt.Println(q.IsEmpty())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Display()

	q.DeEnqueue()
	q.DeEnqueue()
	q.DeEnqueue()
	q.DeEnqueue()
	q.Display()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Display()
}
