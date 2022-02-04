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
}

// NewSingleLinkList function to create a single linklist.
func NewSingleLinkList(rootNode string) SingleLinkListFunc {
	return &singleLinkList{node: &node{value: rootNode}}
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
