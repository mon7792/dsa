package abdulbari

func BinSearch(arr []int, size int, key int) int {
	l := 0
	r := size - 1

	for l <= r {

		mid := int((l + r) / 2)

		if arr[mid] == key {
			return mid
		} else if arr[mid] > key {
			r = arr[mid] - 1
		} else {
			l = arr[mid] + 1
		}
	}

	return -1

}
