package tree

import (
	"fmt"

	"github.com/mon7792/dsa/tree/queue"
)

// tree is a collection of nodes and edges t = {N, E} map[nodes][]edges
// create a binary tree.

//  node stores the value
type node struct {
	data  int
	left  *node
	right *node
}

// store the element of array in tree balanced tree [1,2,3,4,5]
func dummyTree() *node {
	return &node{data: 1, left: &node{data: 2, left: &node{data: 4}, right: &node{data: 5}}, right: &node{data: 3}}
}

// Tree Traversal : Visiting every node in the tree.
// types: Inorder, Preorder, Postorder

// Inorder: Left - Root - Right
func InorderTraverse(tr *node) {
	if tr == nil {
		return
	}

	InorderTraverse(tr.left)
	fmt.Printf(" %d -> ", tr.data)
	InorderTraverse(tr.right)
}

// Preorder:  Root - Left - Right
func PreorderTraverse(tr *node) {
	if tr == nil {
		return
	}

	fmt.Printf(" %d -> ", tr.data)
	PreorderTraverse(tr.left)
	PreorderTraverse(tr.right)
}

// Postorder:  Left - Right - Root
func PostorderTraverse(tr *node) {
	if tr == nil {
		return
	}

	PostorderTraverse(tr.left)
	PostorderTraverse(tr.right)
	fmt.Printf(" %d -> ", tr.data)
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
