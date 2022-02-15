package linklist

import "testing"

// TestNewLinkList display 1 -> 2 -> 3 -> 4 -> 5 link list.
func TestNewLinkList(t *testing.T) {
	sll := NewSingleLinkList("1")
	sll.AddNode("2")
	sll.AddNode("3")
	sll.AddNode("4")
	sll.AddNode("5")
	sll.DisplayNodes()
	sll.DeleteNode()
	sll.DeleteNode()
	sll.DeleteNode()
	sll.DeleteNode()
	sll.DeleteNode()
	sll.DisplayNodes()
}

func TestDelLinkList(t *testing.T) {
	sll := NewSingleLinkList("1")
	sll.AddNode("2")
	sll.AddNode("3")
	sll.AddNode("4")
	sll.AddNode("5")
	sll.DisplayNodes()
	sll.DeleteNodeAtPosition(1)
	sll.DeleteNodeAtPosition(4)
	sll.DeleteNodeAtPosition(2)
	sll.DisplayNodes()
	sll.DeleteNodeAtPosition(100)
}

func TestLinkList(t *testing.T) {
	removeDup()
}

func TestRemoveNthNodeFromLast(t *testing.T) {
	removeNthNodeFromLast()
}

func TestMiddleNode(t *testing.T) {
	removeMiddleNode()
}

func TestPartitionLinkList(t *testing.T) {
	partitionLinkList()
}

func TestSumList(t *testing.T) {
	sumList()
}

func TestPalindromeCheck(t *testing.T) {
	res := palindromeCheck()
	t.Log(res)
}
