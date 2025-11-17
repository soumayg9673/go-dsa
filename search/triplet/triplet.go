package triplet

// This can be used in increasing sorted arrays.
// Sort array before using on unsorted array or decreasing sorted arrays.

// Time Complexity: O(n^3)
func NaiveApproach(arr []int, s int) bool {
	n := len(arr)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if arr[i]+arr[j]+arr[k] == s {
					return true
				}
			}
		}
	}
	return false
}
