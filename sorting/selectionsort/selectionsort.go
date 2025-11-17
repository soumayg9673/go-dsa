package selectionsort

// Selection sort
// Unstable algorith

// Time complexity: O(n^2)
func SelectionSort(arr []int) {
	for i := range arr {
		smallIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallIdx] {
				smallIdx = j
			}
		}
		arr[i], arr[smallIdx] = arr[smallIdx], arr[i]
	}
}
