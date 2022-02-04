package linklist

import "fmt"

//  node struct represent the node in the linklist.
//  it has two fields value and pointer to next node.
type node struct {
	value string
	next  *node
}

type singleLinkList struct {
	node *node
}

// SingleLinkListFunc exposes the function realted to single link list.
type SingleLinkListFunc interface {
	AddNode(value string)
	DisplayNodes()
	DeleteNode()
	Head() *node
	DeleteNodeAtPosition(x int)
}

// NewSingleLinkList function to create a single linklist.
func NewSingleLinkList(rootNode string) SingleLinkListFunc {
	return &singleLinkList{node: &node{value: rootNode}}
}

// Head returns the root node of the list.
func (sl *singleLinkList) Head() *node {
	return sl.node
}

// AddNodes add new node at the end of the list.
func (sl *singleLinkList) AddNode(value string) {
	n := sl.node
	for n.next != nil {
		n = n.next
	}
	n.next = &node{value: value}
}

// DisplayNodes display all the nodes from the start to end.
func (sl *singleLinkList) DisplayNodes() {
	if sl.node == nil {
		return
	}
	n := sl.node
	for n.next != nil {
		fmt.Print(n.value, " -> ")
		n = n.next
	}
	fmt.Print(n.value)
	fmt.Println()
}

// DeleteNode delete node at the end of the list.
func (sl *singleLinkList) DeleteNode() {
	if sl.node.next == nil {
		sl.node = nil
		return
	}
	n := sl.node
	for n.next != nil && n.next.next != nil {
		n = n.next
	}
	n.next = nil
}

// DeleteNodeAtPosition delete node at the x position from of the list.
func (sl *singleLinkList) DeleteNodeAtPosition(pos int) {
	//  delete single node
	if sl.node.next == nil {
		sl.node = nil
		return
	}
	//  remove 1st element in the list
	if pos == 1 {
		head := sl.node
		sl.node = sl.node.next
		head.next = nil
		return
	}

	// remove other element in the list
	nodePos := 1
	n := sl.node
	for n.next != nil && pos-1 != nodePos {
		n = n.next
		nodePos++
	}

	if n.next == nil && pos-1 != nodePos {
		fmt.Println("no node at this position ", pos)
		return
	}

	lfNode := n
	delNode := n.next
	rgNode := n.next.next
	lfNode.next = rgNode
	delNode.next = nil
}

// // WAP : write code to remove duplicates from an unsorted linked list.
func removeDup() {
	sll := NewSingleLinkList("1")
	sll.AddNode("2")
	sll.AddNode("2")
	sll.AddNode("2")
	sll.AddNode("4")
	sll.AddNode("2")
	sll.DisplayNodes()

	head := sll.Head()
	for head.next != nil {
		cmpVal := head.value
		nd := head
		for nd.next != nil {
			fmt.Print(nd.value, "->")
			if nd.next.value == cmpVal {
				lfNode := nd
				delNode := nd.next
				rgNode := nd.next.next
				lfNode.next = rgNode
				delNode.next = nil
			} else {
				nd = nd.next
			}
		}
		fmt.Print(nd.value)
		fmt.Println()
		head = head.next
	}

	sll.DisplayNodes()
}
