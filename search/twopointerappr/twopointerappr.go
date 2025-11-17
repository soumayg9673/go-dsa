package twopointerappr

// Time Complexity: O(n^2)
// Auxilary Space: O(1)
func NaiveApproach(arr []int, s int) bool {
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if s == arr[i]+arr[j] {
				return true
			}
		}
	}
	return false
}
