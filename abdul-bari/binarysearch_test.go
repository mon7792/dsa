package abdulbari

import "testing"

func TestBinSearch_valid(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	key := 6
	size := 8

	res := BinSearch(a, size, key)
	t.Log(res)
}

func TestBinSearch_InValid(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	key := 26
	size := 8

	res := BinSearch(a, size, key)
	t.Log(res)
}
