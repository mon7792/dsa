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

type platesStack struct {
	stacks        []*stackNodes // contains all the stack of plates.
	threshold     int           // max number of plates in each stack
	platesStackNo int           // number of plate in the current stack
	stackPos      int           // the stack currently filled
	top           *node         // the pointer to top node in the stack.
}

func NewStackOfPlates(threshold int) *platesStack {
	return &platesStack{
		stacks:        make([]*stackNodes, 0),
		threshold:     threshold,
		platesStackNo: 0,
		stackPos:      0,
		top:           nil,
	}
}

func (ps *platesStack) Push(nodeValue string) {

	// fmt.Println(ps.stacks)
	// return
	// check if the value is reached threshold
	if ps.platesStackNo == ps.threshold {
		ps.stackPos++
		ps.platesStackNo = 0
	}

	// 1st node in the stack
	if ps.platesStackNo == 0 {
		ps.stacks = append(ps.stacks, &stackNodes{top: &node{value: nodeValue}})

		// ps.stacks[ps.stackPos] = &stackNodes{top: &node{value: nodeValue}}
		ps.platesStackNo++
		ps.top = ps.stacks[ps.stackPos].top
		return
	}

	ps.stacks[ps.stackPos].top = &node{value: nodeValue, next: ps.stacks[ps.stackPos].top}
	ps.top = ps.stacks[ps.stackPos].top

	ps.platesStackNo++
}

func (ps *platesStack) Display() {
	for i := 0; i <= ps.stackPos; i++ {

		fmt.Println("STACK POS:", i+1)
		// display plates in the single stack
		nd := ps.stacks[i].top
		if nd == nil {
			continue
		}
		for nd.next != nil {
			fmt.Printf("%s ->", nd.value)
			nd = nd.next
		}
		fmt.Printf("%s", nd.value)
		fmt.Println()
	}
}

func (ps *platesStack) Pop() {
	//  no nodes to remove.
	if ps.platesStackNo == -1 && ps.stackPos == -1 {
		return
	}

	// check the element to remove from particular array
	if ps.platesStackNo > 0 {
		ps.stacks[ps.stackPos].top = ps.stacks[ps.stackPos].top.next
		ps.platesStackNo--
		return
	}

	if ps.platesStackNo == 0 {
		ps.stackPos--
		ps.platesStackNo = ps.threshold
		ps.stacks[ps.stackPos].top = ps.stacks[ps.stackPos].top.next
		ps.platesStackNo--
		return
	}
}

//  IMPLEMENT STACKOFPLATES IN DIFFERENT WAY.
