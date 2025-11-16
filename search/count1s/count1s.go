package count1s

// Time Complexity: O(n)
func NaiveApproach(arr []int) int {
	i := 0
	for ; i < len(arr); i++ {
		if arr[i] == 1 {
			break
		}
	}
	return len(arr) - i
}
