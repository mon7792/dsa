package stack

import "fmt"

type node struct {
	data interface{}
	next *node
}

type linkStack struct {
	head *node
	size int
}

// LinkStackFunc exposes all the stack functionality.
type LinkStackFunc interface {
	Push(item interface{})
	Pop()
	IsEmpty() bool
	Peek() interface{}
	Display()
}

// New Stack
func New() LinkStackFunc {
	return &linkStack{}
}

// Push add the element to the stack.
func (s *linkStack) Push(item interface{}) {
	if s.IsEmpty() {
		s.head = &node{data: item}
		s.size = 1
		return
	}
	newNd := &node{data: item, next: s.head}
	s.head = newNd
	s.size++
}

// Pop remove the top element from the stack
func (s *linkStack) Pop() {
	if s.IsEmpty() {
		s.head = nil
		s.size = 0
		return
	}
	s.head = s.head.next
	s.size--
}

// IsEmpty() check if the stack is empty
func (s *linkStack) IsEmpty() bool {
	return s.size == 0
}

// Peek return the element at the top of the stack.
func (s *linkStack) Peek() interface{} {
	return s.head.data
}

// Display the elements in the stack
func (s *linkStack) Display() {
	if s.IsEmpty() {
		return
	}
	st := s.head
	for st.next != nil {
		fmt.Printf("|  %d  |\n", st.data)
		st = st.next
	}
	fmt.Printf("|  %d  |\n", st.data)
	fmt.Print("|_____|\n")
}
