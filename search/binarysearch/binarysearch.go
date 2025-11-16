package binarysearch

func Iterative(arr []int, s int) int {
	low := 0
	high := len(arr)

	for low < high {
		m := (low + high) / 2
		if arr[m] == s {
			return m
		} else if arr[m] > s {
			high = m - 1
		} else {
			low = m + 1
		}
	}

	return -1
}
