package queue

import "fmt"

//  node stores the value
type node struct {
	data interface{}
	next *node
}

// queue is FIFO: First In First Out
type linkStore struct {
	front *node
	back  *node
	size  int
}

type LinkStoreFunc interface {
	Enqueue(item interface{})
	IsEmpty() bool
	DeEnqueue() interface{}
	Display()
}

// New creates a new link queue structure.
func New() LinkStoreFunc {
	return &linkStore{}
}

// Enqueue add the item to the queue.
func (q *linkStore) Enqueue(item interface{}) {
	if q.IsEmpty() {
		q.front = &node{data: item}
		q.back = q.front
		q.size = 1
		return
	}

	// increase the size
	q.size++

	// add the element to the back
	q.back.next = &node{data: item}
	q.back = q.back.next
}

// DeEnqueue remove the element from the back.
func (q *linkStore) DeEnqueue() interface{} {
	if q.size == 0 {
		q.front = nil
		q.back = nil
		return nil
	}

	// decrease the size
	q.size--

	rmElem := q.front.data
	// remove the element from the front
	q.front = q.front.next
	return rmElem
}

// IsEmpty checks if any element is present in the queue.
func (q *linkStore) IsEmpty() bool {
	return q.size == 0
}

// Display the elements in the queue.
func (q *linkStore) Display() {
	if q.IsEmpty() {
		fmt.Println("EMPTY Queue")
		return
	}
	ql := q.front
	for ql.next != nil {
		fmt.Printf(" %d -> ", ql.data)
		ql = ql.next
	}
	fmt.Printf(" %d ", ql.data)
	fmt.Println()
}
