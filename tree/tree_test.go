package tree

import "testing"

// func TestDisplay(t *testing.T) {
// 	//
// 	tr := CreateSimpleTree()
// 	Display(tr)
// }

func TestInOrder(t *testing.T) {
	//
	tr := create2CTree()
	displayInOrderIterative(tr.root)
}

func TestPreOrder(t *testing.T) {
	//
	tr := create2CTree()
	displayPreOrderIterative(tr.root)
}

func TestPostOrder(t *testing.T) {
	//
	tr := create2CTree()
	displayPostOrderIterative(tr.root)
}
