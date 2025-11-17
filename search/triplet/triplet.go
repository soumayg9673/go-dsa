package triplet

import (
	"github.com/soumayg9673/go-dsa/search/twopointerappr"
)

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

// Time Complexity: O(n^2)
func OptimalApproach(arr []int, s int) bool {
	for i := 0; i < len(arr)-2; i++ {
		if ok := twopointerappr.OptimalApproach(arr[i+1:], s-arr[i]); ok {
			return true
		}
	}
	return false
}
