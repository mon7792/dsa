package tree

import (
	"fmt"
	"testing"
)

func TestInorderTraverse(t *testing.T) {
	t.Log("Recursion InOrder Traversal")
	t.Log("Exp: 4 ->  2 ->  5 ->  1 ->  3 -> ")
	InorderTraverse(dummyTree())
	fmt.Println()
}

func TestPreorderTraverse(t *testing.T) {
	t.Log("Recursion PreOrder Traversal")
	t.Log("Exp: 1 ->  2 ->  4 ->  5 ->  3 -> ")
	PreorderTraverse(dummyTree())
	fmt.Println()
}

func TestPostorderTraverse(t *testing.T) {
	t.Log("Recursion PostOrder Traversal")
	t.Log("Exp: 4 ->  5 ->  2 ->  3 ->  1 -> ")
	PostorderTraverse(dummyTree())
	fmt.Println()
}

func TestBFS(t *testing.T) {
	t.Log("BFS")
	BFS(dummyTree())
	fmt.Println()
}
