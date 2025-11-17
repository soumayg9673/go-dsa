package mergesort

// Merge sort is recursive sorting algorithm
// with divide and conquer algorithm.
// Also a stable sort.

// Time complexity: O(n log(n))
// Auxilary space: O(n)

func MergeSort(arr []int) []int {
	n := len(arr)

	if n < 2 {
		return arr
	}

	first := MergeSort(arr[:n/2])
	second := MergeSort(arr[n/2:])

	return merge(first, second)
}

func merge(first, second []int) (res []int) {
	i, j := 0, 0

	for i < len(first) && j < len(second) {
		if first[i] <= second[j] {
			res = append(res, first[i])
			i++
		} else {
			res = append(res, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		res = append(res, first[i])
	}
	for ; j < len(second); j++ {
		res = append(res, second[j])
	}

	return
}
