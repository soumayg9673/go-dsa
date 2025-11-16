package trappingrainwater

/*
Time Complexity: Θ(n)
Auxilary Space: Θ(n)
*/
func OptimalApproach(w []int) (vol int) {
	n := len(w)

	lMax := make([]int, n)
	lMax[0] = w[0]
	for i := 1; i < n; i++ {
		lMax[i] = max(w[i], lMax[i-1])
	}

	rMax := make([]int, n)
	rMax[n-1] = w[n-1]
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(w[i], rMax[i+1])
	}

	for i := 1; i < n-1; i++ {
		vol += min(lMax[i], rMax[i]) - w[i]
	}
	return
}
