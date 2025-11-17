package twopointerappr

// This can be used in Increasing Sorted arrays

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

// Time Complexity: O(n)
// Auxilary Space: O(1)
func OptimalApproach(arr []int, s int) bool {
	for i, j := 0, len(arr)-1; i <= j; {
		if arr[i]+arr[j] == s {
			return true
		} else if arr[i]+arr[j] > s {
			j--
		} else {
			i++
		}
	}
	return false
}
