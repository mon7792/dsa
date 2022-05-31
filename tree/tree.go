package tree

import "fmt"

type node struct {
	name     string
	children []*node
}

type tree struct {
	root *node
}

func New() *tree {
	return &tree{}
}

func CreateSimpleTree() *tree {
	t := New()

	// root node
	t.root = &node{}
	t.root.name = "A"
	t.root.children = make([]*node, 0)
	// 1st level
	t.root.children = append(t.root.children, &node{name: "B"}, &node{name: "C"}, &node{name: "D"})

	// 2nd Level
	t.root.children[0].children = make([]*node, 0)
	t.root.children[1].children = make([]*node, 0)
	t.root.children[2].children = make([]*node, 0)

	t.root.children[0].children = append(t.root.children[0].children, &node{name: "E"}, &node{name: "F"})
	t.root.children[1].children = append(t.root.children[1].children, &node{name: "G"}, &node{name: "H"})
	t.root.children[2].children = append(t.root.children[2].children, &node{name: "I"}, &node{name: "J"})

	return t
}

func Display(t *tree) {
	// display the 1st level
	// fmt.Println(t.root.name)
	// children := t.root.children
	// i := 0
	// for i < len(children) {
	// 	fmt.Print(children[i].name, " ")
	// 	i++
	// }
	// fmt.Println()

	queu := NewQue()

	// 1st iteration
	nd := t.root
	fmt.Println(nd.name)
	if nd.children != nil {
		queu.Push(nd.children...)
	}
	nd = queu.Pop()

	for len(queu.arr) != 0 {
		fmt.Println(nd.name)
		if nd.children != nil {
			queu.Push(nd.children...)
		}
		nd = queu.Pop()
	}

}

type que struct {
	arr []*node
}

// NewQue docs
func NewQue() *que {
	return &que{arr: make([]*node, 0)}
}

func (q *que) Push(value ...*node) {
	q.arr = append(q.arr, value...)
}

func (q *que) Pop() *node {
	val := q.arr[0]
	q.arr = q.arr[1:]

	return val
}

// TREE traversal : 2 child

type tree2C struct {
	root *node2C
}

type node2C struct {
	data  int
	left  *node2C
	right *node2C
}

func create2CTree() tree2C {
	tr := tree2C{root: &node2C{1, nil, nil}}

	tr.root.left = &node2C{12, nil, nil}
	tr.root.left.left = &node2C{5, nil, nil}
	tr.root.left.right = &node2C{6, nil, nil}

	tr.root.right = &node2C{9, nil, nil}
	tr.root.right.left = &node2C{10, nil, nil}
	tr.root.right.right = &node2C{11, nil, nil}

	return tr
}

func displayInOrder(nd *node2C) {
	if nd != nil {
		displayInOrder(nd.left)
		fmt.Println(nd.data)
		displayInOrder(nd.right)
	}
}

func displayInOrderIterative(nd *node2C) {
	st := NewStack()
	current := nd

	for current != nil || !st.IsEmpty() {
		if current != nil {
			st.Push(current)
			current = current.left
		} else {
			current = st.Pop()
			fmt.Println(current.data)
			current = current.right
		}
	}

}

// displayPreOrderIterative : root -> left -> right
func displayPreOrderIterative(nd *node2C) {
	st := NewStack()
	current := nd

	for current != nil || !st.IsEmpty() {
		if current != nil {
			st.Push(current)
			fmt.Println(current.data)
			current = current.left
		} else {
			current = st.Pop()
			current = current.right
		}
	}

}

// displayPostOrderIterative : left -> right -> root
func displayPostOrderIterative(nd *node2C) {
	st := NewStack()
	current := nd
	var pre *node2C
	height := 0

	for current != nil || !st.IsEmpty() {
		if current != nil {
			st.Push(current)
			current = current.left
			height++
		} else {
			current = st.Peek()

			if current.right == nil || current.right == pre {
				fmt.Println(current.data)
				st.Pop()
				pre = current
				current = nil
				height--
			} else {
				current = current.right
				height++
			}

		}
	}

}

//  node represent individual element of the stack
type stNode struct {
	value *node2C
	next  *stNode
}

// stackNodes contains the list of nodes present in the stack.
type stackNodes struct {
	top *stNode
}

// StackFunction contains stack functionality.
type StackFunction interface {
	Pop() *node2C
	Push(value *node2C)
	IsEmpty() bool
	Peek() *node2C
}

// NewStack creates a new stack.
func NewStack() StackFunction {
	return &stackNodes{
		top: nil,
	}
}

// Pop remove the top item from the stack.
func (s *stackNodes) Pop() *node2C {
	if s.top == nil {
		return nil
	}
	val := s.top.value
	s.top = s.top.next

	return val
}

// Push add the node with value at the top of the stack.
func (s *stackNodes) Push(value *node2C) {
	if s.top == nil {
		s.top = &stNode{value: value}
		return
	}

	s.top = &stNode{value: value, next: s.top}
}

func (s *stackNodes) IsEmpty() bool {

	return s.top == nil
}

func (s *stackNodes) Peek() *node2C {

	return s.top.value
}
