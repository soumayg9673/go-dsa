package quicksort

// Quick sort
// Unstable algorithm

// Time complexity
// Best case - O(n log(n))
// Worst case - O(n^2) for sorted arrays
func Lomuto(arr []int, low, high int) {
	if low < high {
		mid := lomPartition(arr, low, high)
		Lomuto(arr, low, mid-1)
		Lomuto(arr, mid+1, high)
	}
}

func lomPartition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
