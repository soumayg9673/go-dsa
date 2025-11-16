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

func Recursive(arr []int, s, low, high int) int {
	if low >= high {
		return -1
	} else {
		m := (low + high) / 2
		if arr[m] == s {
			return m
		} else if arr[m] > s {
			return Recursive(arr, s, low, m-1)
		} else {
			return Recursive(arr, s, m+1, high)
		}
	}
}
