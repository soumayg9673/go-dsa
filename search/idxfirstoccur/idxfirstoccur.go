package idxfirstoccur

// Time Complexity: O(n)
func NaiveApproach(arr []int, s int) int {
	for i := range arr {
		if arr[i] == s {
			return i
		}
	}
	return -1
}
