package insertionsort

// Insertion sort
// Best for small lists and partially sorted arrays.
// Also a stable sort.

// Time complexity
// best case - O(n) for pre-sorted array
// worst case - O(n^2)
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}
