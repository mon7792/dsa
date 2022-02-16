package stack

import "testing"

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
