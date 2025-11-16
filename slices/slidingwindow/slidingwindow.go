package slidingwindow

import (
	"math"
)

/*
Time Complexity: O(nk)
Auxilary Space: O(1)
*/
func NaiveApproach(arr []int, k int) int {
	res := math.MinInt

	for i := 0; i < len(arr) && i+k-1 < len(arr); i++ {
		temp := 0
		for j := range k {
			temp += arr[i+j]
		}
		res = max(res, temp)
	}

	return res
}

/*
Time Complexity: O(n)
Auxilary Space: O(1)
*/
func OptimalApproach(arr []int, k int) int {
	res := math.MinInt
	temp := 0

	for i := 0; i+k-1 < len(arr); i++ {
		if i == 0 {
			for j := range k {
				temp += arr[j]
			}
		} else {
			temp = temp - arr[i-1] + arr[i+k-1]
		}
		res = max(res, temp)
	}

	return res
}
