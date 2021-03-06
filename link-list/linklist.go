package linklist

import (
	"fmt"
	"strconv"
)

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
	AddNodeAtStart(value string)
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

// AddNodes add new node at the end of the list.
func (sl *singleLinkList) AddNodeAtStart(value string) {
	n := sl.node
	sl.node = &node{value: value, next: n}
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
	if pos < 1 {
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

func helperStrToInt(value string) int {
	intVal, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return intVal
}

// 1. WAP : write code to remove duplicates from an unsorted linked list.
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

// 2. WAP : write algorithm to remove the Kth to last element of the single linked list.
func removeNthNodeFromLast() {
	nodePosToDel := 2
	sll := NewSingleLinkList("A")
	sll.AddNode("B")
	sll.AddNode("C")
	sll.AddNode("D")
	sll.AddNode("E")
	sll.AddNode("F")
	sll.DisplayNodes()

	// count the length of the linklist
	linkLen := 0
	head := sll.Head()
	for head.next != nil {
		linkLen++
		head = head.next
	}
	linkLen++

	fmt.Println("THE LENGTH OF LINK LIST:", linkLen)

	sll.DeleteNodeAtPosition(linkLen - nodePosToDel + 1)
	sll.DisplayNodes()
}

func removeMiddleNode() {
	// Approach 1: Loop through the linklist get the total length.
	// get the ciel number of len/2
	// remove the node at that position

	// Approach 2: runner technique.
	//  two pointer:
	//  1. slow pointer : iterates one step at a time
	//  2. fast pointer : iterates two step at a time

	// slow pointer starts late
	// fast pointer check till node.next != nil || node.next.next != nil

	//  Q: what happens to the missing pointer.

	sll := NewSingleLinkList("A")
	sll.AddNode("B")
	sll.AddNode("C")
	sll.AddNode("D")
	sll.AddNode("E")
	sll.AddNode("F")
	sll.DisplayNodes()

	prevNd := sll.Head()
	slowPtr := sll.Head()
	fastPtr := sll.Head()

	for fastPtr.next.next != nil {

		fastPtr = fastPtr.next.next
		prevNd = slowPtr
		slowPtr = slowPtr.next
	}
	prevNd.next = prevNd.next.next
	sll.DisplayNodes()
}

func partitionLinkList() {

	partValue := 5
	//  APPROACH 1: any number less than the value should be added at the start of the list
	//  TODO: move the value towards the second half of the partition.
	sll := NewSingleLinkList("3")
	sll.AddNode("5")
	sll.AddNode("8")
	sll.AddNode("5")
	sll.AddNode("10")
	sll.AddNode("2")
	sll.AddNode("1")
	sll.DisplayNodes()

	iterPtr := sll.Head()
	// if headPtr.value >= val
	// remove capture the head element and put it into lastelem
	prevNd := iterPtr
	if iterPtr.next != nil {
		iterPtr = iterPtr.next
	}

	for iterPtr.next != nil {
		//  iterPtr.value < value
		//  add the element at the start of the list.
		intVal, err := strconv.Atoi(iterPtr.value)
		if err != nil {
			panic(err)
		}
		if intVal < partValue {
			//  remove the current node.
			prevNd.next = iterPtr.next
			//  add it to the start of the list
			sll.AddNodeAtStart(iterPtr.value)
		} else {
			prevNd = iterPtr
		}
		iterPtr = iterPtr.next
	}

	//  last node
	intVal, err := strconv.Atoi(iterPtr.value)
	if err != nil {
		panic(err)
	}
	if intVal < partValue {
		//  remove the current node.
		prevNd.next = iterPtr.next
		//  add it to the start of the list
		sll.AddNodeAtStart(iterPtr.value)
	}

	sll.DisplayNodes()
}

// WAP 1: Write a program to add numbers represented in the reverse linklist.
// INPUT: (7 -> 1 -> 6) + (5 -> 9 -> 2)  => (2 -> 1 -> 9)
func sumList() {
	// reverse order
	ftll := NewSingleLinkList("7")
	ftll.AddNode("1")
	ftll.AddNode("6")

	sdll := NewSingleLinkList("5")
	sdll.AddNode("9")
	sdll.AddNode("2")
	// helperStrToInt()

	fstValue := ""
	sndValue := ""
	fstPtr := ftll.Head()
	sndPtr := sdll.Head()
	for fstPtr.next != nil {

		fstValue = fstPtr.value + fstValue
		fstPtr = fstPtr.next
	}
	fstValue = fstPtr.value + fstValue

	for sndPtr.next != nil {

		sndValue = sndPtr.value + sndValue
		sndPtr = sndPtr.next
	}
	sndValue = sndPtr.value + sndValue

	totalVAl := fmt.Sprintf("%d", helperStrToInt(fstValue)+helperStrToInt(sndValue))

	resll := NewSingleLinkList(string(totalVAl[len(totalVAl)-1]))
	for i := len(totalVAl) - 2; i >= 0; i-- {
		resll.AddNode(string(totalVAl[i]))
	}

	resll.DisplayNodes()

}

// WAP function to check if the linklist is palindrome.
func palindromeCheck() bool {
	sll := NewSingleLinkList("m")
	sll.AddNode("a")
	// sll.AddNode("d")
	sll.AddNode("a")
	sll.AddNode("m")
	sll.DisplayNodes()

	fastPtr := sll.Head()
	slowPtr := sll.Head()

	linkLen := 1

	for fastPtr.next != nil && fastPtr.next.next != nil {

		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		linkLen = linkLen + 1
	}

	iter := 0
	if linkLen%2 != 0 {
		iter = 1
	}

	for slowPtr.next != nil {
		slowPtr = slowPtr.next
		nodePos := linkLen - iter
		cmpNode := sll.Head()

		for nodePos > 1 {
			cmpNode = cmpNode.next
			nodePos--
		}

		if cmpNode.value != slowPtr.value {
			return false
		}
		iter++
	}

	return true
}
