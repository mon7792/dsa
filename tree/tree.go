package tree

import (
	"fmt"

	"github.com/mon7792/dsa/tree/queue"
)

// tree is a collection of nodes and edges t = {N, E} map[nodes][]edges
// create a binary tree.

// node stores the value
type node struct {
	data  int
	left  *node
	right *node
}

func dummyTree() *node {
	return &node{data: 1, left: &node{data: 12, left: &node{data: 5}, right: &node{data: 6}}, right: &node{data: 9}}
}

// Inorder Recusrive Traversal
func InOrder(nd *node) {
	if nd == nil {
		return
	}
	InOrder(nd.left)
	fmt.Println(nd.data)
	InOrder(nd.right)
}

// Preorder Recusrive Traversal
func PreOrder(nd *node) {
	if nd == nil {
		return
	}
	fmt.Println(nd.data)
	PreOrder(nd.left)
	PreOrder(nd.right)
}

// Postorder Recusrive Traversal
func PostOrder(nd *node) {
	if nd == nil {
		return
	}
	PostOrder(nd.left)
	PostOrder(nd.right)
	fmt.Println(nd.data)
}

// CheckIsFullBinaryTree
func CheckIsFullBinaryTree(nd *node) bool {

	// traverse through evey node and check if the children is 0 or 2
	if nd == nil {
		return true
	}

	if nd.left == nil && nd.right == nil {
		return true
	}

	if nd.left != nil && nd.right != nil {
		return CheckIsFullBinaryTree(nd.left) && CheckIsFullBinaryTree(nd.right)
	}

	return false

}

// BFS using queue
func BFS(tr *node) {
	q := queue.New()

	q.Enqueue(tr)

	for !q.IsEmpty() {
		nd := q.DeEnqueue().(*node)
		fmt.Printf("%d ->", nd.data)
		if nd.left != nil {
			q.Enqueue(nd.left)
		}

		if nd.right != nil {
			q.Enqueue(nd.right)
		}
	}

}
