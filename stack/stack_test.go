package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	// 1,2,3,4
	//  pop () 2,3,4
	//  push(8) 8,2,3,4
	//  push(5) 5,8,2,3,4
	//  peek() 5

	st := NewStack("4")
	st.Push("3")
	st.Push("2")
	st.Push("1")
	t.Log("INITAL STACK")
	st.Display()

	t.Log("POP STACK")
	st.Pop()
	st.Display()
	t.Log("PUSH STACK")
	st.Push("8")
	st.Display()
	t.Log("PUSH STACK")
	st.Push("5")
	st.Display()
	t.Log("PEEK STACK")
	t.Log(st.Peek())

}

func TestStackArr(t *testing.T) {
	// stackInArray("1", "2", "3", "4")
	st := NewStackArr()

	st.Display()
	st.Push("1")
	st.Push("2")
	st.Push("3")
	st.Push("4")

	st.Display()
	st.Pop()
	st.Pop()

	st.Display()
}

func TestStackOfPlates(t *testing.T) {

	st := NewStackOfPlates(3)

	st.Push("1")
	st.Push("2")
	st.Push("3")

	st.Push("4")
	st.Push("5")
	st.Push("6")
	st.Push("7")
	st.Display()

	fmt.Println("==========")
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()

	st.Display()
}
