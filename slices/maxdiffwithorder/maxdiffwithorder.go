package maxdiffwithorder

/*
Maximum value of arr[j] - arr[i], where j > i
*/

/*
Time Complexity: O(n^2)
Auxilary Space: Θ(1)
*/
func NaiveApproach(a []int) (res int) {
	if len(a) < 2 {
		panic("bad input")
	}
	if len(a) == 2 {
		return a[1] - a[0]
	}

	res = a[1] - a[0]
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			res = max(res, a[j]-a[i])
		}
	}
	return
}

/*
Time Complexity: Θ(n)
Auxilary Space: Θ(1)
*/
func OptimalApproach(a []int) (res int) {
	res = a[1] - a[0]
	m := a[0]
	for i := 1; i < len(a); i++ {
		res = max(res, a[i]-m)
		m = min(m, a[i])
	}
	return
}
