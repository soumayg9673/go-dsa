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

/*
Time Complexity: O(n)
Auxilary Space: O(n)
*/
func OptimalSolution(arr []int) bool {
	n := len(arr)

	// Cummulative prefix sum
	prefixSum := make([]int, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = arr[i] + prefixSum[i-1]
	}

	// Cummulative suffix sum
	suffixSum := make([]int, n)
	suffixSum[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = arr[i] + suffixSum[i+1]
	}

	for i := range arr {
		switch i {
		case 0:
			left := 0
			right := suffixSum[i+1]
			if left == right {
				return true
			}
		case n - 1:
			left := prefixSum[n-2]
			right := 0
			if left == right {
				return true
			}
		default:
			left := prefixSum[i-1]
			right := suffixSum[i+1]
			if left == right {
				return true
			}
		}
	}

	return false
}
