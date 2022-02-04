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
