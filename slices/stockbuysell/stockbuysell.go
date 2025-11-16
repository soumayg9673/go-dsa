package stockbuysell

/*
Time Complexity: Θ(n)
*/
func OptimalApproach(m []int) (p int) {
	for i := 1; i < len(m); i++ {
		if m[i] > m[i-1] {
			p += m[i] - m[i-1]
		}
	}
	return
}
