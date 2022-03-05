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

// stackArr represent the array representation of the stack.
type stackArr struct {
	topPt int
	size  int
	arr   []string
}

// NewStackArr exposes the array based implementation for stack
func NewStackArr() StackFunction {
	return &stackArr{
		topPt: -1,
		size:  0,
		arr:   make([]string, 0),
	}
}

// Pop the element from the stack.
func (st *stackArr) Pop() {
	if st.size == 0 {
		return
	}

	st.arr = append(st.arr[:st.topPt], st.arr[st.topPt+1:]...)

	st.topPt = st.topPt - 1
	st.size = st.size - 1

}

// Push value in the array.
func (st *stackArr) Push(value string) {
	st.topPt = st.topPt + 1
	st.size = st.size + 1

	st.arr = append(st.arr, value)
}

// Peek value in the array.
func (st *stackArr) Peek() string {
	if st.size == 0 {
		return ""
	}
	return st.arr[st.topPt]
}

// Display elements in the stack.
func (st *stackArr) Display() {
	if st.size == 0 {
		return
	}
	fmt.Println("ELEMENT IN STACK")
	for i := st.topPt; i >= 0; i-- {
		fmt.Println(st.arr[i])
	}
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
