package equilibriumpoint

/*
Time Complexity: O(n^2)
Auxilary Space: O(1)
*/
func NaiveApproach(arr []int) bool {
	left, right := 0, 0
	for i := range arr {
		for j := range i {
			left += arr[j]
		}
		for j := i + 1; j < len(arr); j++ {
			right += arr[j]
		}

		if left == right {
			return true
		}

		left = 0
		right = 0
	}
	return false
}
