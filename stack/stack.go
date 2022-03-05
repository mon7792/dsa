// package stack contains sample programs for practising the stack concepts.
package stack

import "fmt"

//  node represent individual element of the stack
type node struct {
	value string
	next  *node
}

// stackNodes contains the list of nodes present in the stack.
type stackNodes struct {
	top *node
}

// StackFunction contains stack functionality.
type StackFunction interface {
	Pop()
	Push(value string)
	Peek() string
	Display()
}

// NewStack creates a new stack.
func NewStack(value string) StackFunction {
	return &stackNodes{
		top: &node{value: value},
	}
}

// Pop remove the top item from the stack.
func (s *stackNodes) Pop() {
	if s.top == nil {
		return
	}
	s.top = s.top.next
}

// Push add the node with value at the top of the stack.
func (s *stackNodes) Push(value string) {
	if s.top == nil {
		s.top = &node{value: value}
		return
	}

	s.top = &node{value: value, next: s.top}
}

// Peek returns the value at the top of the stack.
func (s *stackNodes) Peek() string {
	if s.top == nil {
		return ""
	}
	return s.top.value
}

func (s *stackNodes) Display() {
	if s.top == nil {
		return
	}
	ndList := s.top
	fmt.Println("----")
	for ndList.next != nil {
		fmt.Println(ndList.value)
		fmt.Println("----")
		ndList = ndList.next
	}
	fmt.Println(ndList.value)
	fmt.Println("----")

}

type stackArr struct {
	topPt int
	size  int
	arr   []string
}

func stackInArray(stackElem ...string) {

	// load size at once
	stArr := stackArr{
		arr: make([]string, 0),
	}

	// load stack elem in the array.
	for i := range stackElem {
		stArr.topPt = i
		stArr.size = i + 1
		stArr.arr = append(stArr.arr, stackElem[i])

	}

	fmt.Println(stArr.topPt, stArr.size, stArr.arr)

}

// func threeStackOneArray() {

// 	arr := make([]string, 10, 20)
// 	fmt.
// }
