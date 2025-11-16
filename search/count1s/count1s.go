package count1s

import "github.com/soumayg9673/go-dsa/search/idxfirstoccur"

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

// Time Complexity: O(log(n))
func OptimalApproach(arr []int) int {
	idx := idxfirstoccur.OptimalApproachIterative(arr, 1)
	if idx != -1 {
		return len(arr) - idx
	}
	return 0
}
