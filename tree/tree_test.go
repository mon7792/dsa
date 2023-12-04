package tree

import (
	"testing"
)

func TestInorderTraverse(t *testing.T) {
	t.Log("Recursion InOrder Traversal")
	t.Log("Exp: 4 ->  2 ->  5 ->  1 ->  3 -> ")
	InOrder(dummyTree())

}

func TestPreorderTraverse(t *testing.T) {
	t.Log("Recursion PreOrder Traversal")
	t.Log("Exp: 1 ->  2 ->  4 ->  5 ->  3 -> ")
	PreOrder(dummyTree())

}

func TestPostorderTraverse(t *testing.T) {
	t.Log("Recursion PostOrder Traversal")
	t.Log("Exp: 4 ->  5 ->  2 ->  3 ->  1 -> ")
	PostOrder(dummyTree())
}

func TestCheckIsFullBinaryTree(t *testing.T) {
	t.Log("check if the dummy tree is full binary tree")
	t.Log("Exp: 4 ->  5 ->  2 ->  3 ->  1 -> ")
	result := CheckIsFullBinaryTree(dummyTree())
	t.Log(result)
}

func TestBFS(t *testing.T) {
	t.Log("BFS")
	BFS(dummyTree())
	fmt.Println()
}
