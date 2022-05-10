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
