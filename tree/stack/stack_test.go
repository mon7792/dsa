package stack

import "testing"

func TestStack(t *testing.T) {
	st := New()

	st.Push(1)
	st.Push(2)
	st.Display()
	st.Push(3)
	st.Display()
	st.Pop()
	t.Log(st.Peek())
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	st.Pop()
	st.Display()
	st.Push(4)
	st.Push(5)
	st.Display()
}
