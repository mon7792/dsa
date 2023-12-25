package tree

import (
	"fmt"

	"github.com/mon7792/dsa/tree/queue"
	"github.com/mon7792/dsa/tree/stack"
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

// InOrderStackTraverse traverse through the tree using stack: dfs
func InOrderStackTraverse(crNd *node) {
	st := stack.New()
	nd := crNd
	for {
		if nd != nil {
			st.Push(*nd)
			nd = nd.left
		} else {
			if st.IsEmpty() {
				break
			}
			curr := st.Pop()

			// type conversion and fix here
			crrNd, _ := curr.(node)
			fmt.Println(crrNd.data)
			nd = crrNd.right
		}
	}
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

// MaxDepthRecursive return the maximum depth of the tree using recurisve technique
func MaxDepthRecursive(root *node) {
	maxDepth := helperMaxDepthRecursive(root, 0)

	fmt.Println("Max Depth:", maxDepth)
}

func helperMaxDepthRecursive(root *node, level int) int {
	if root == nil {
		return level - 1
	}

	ldist := helperMaxDepthRecursive(root.left, level+1)
	rdist := helperMaxDepthRecursive(root.right, level+1)

	return helperMax(ldist, rdist)
}

// helperMax return the max number
func helperMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
