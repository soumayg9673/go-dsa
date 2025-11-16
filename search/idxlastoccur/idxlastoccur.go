package idxlastoccur

// Time Complexity: O(n)
func NaiveApproach(arr []int, s int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == s {
			return i
		}
	}
	return -1
}
