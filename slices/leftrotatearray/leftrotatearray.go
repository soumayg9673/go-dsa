package leftrotatearray

import (
	"slices"
)

/*
Time Complexity: Θ(n)
Auxilary Space:
  - Worst case: O(n)
  - Best case: Θ(1)
*/
func NaiveApproach(a []int, d int) {
	n := len(a)

	temp := []int{}
	for i := range d {
		temp = append(temp, a[i])
	}

	for i := d; i < n; i++ {
		a[i-d] = a[i]
	}

	for i := range temp {
		a[n-d+i] = temp[i]
	}
}

/*
Time Complexity: Θ(n)
Auxilary Space: Θ(1)
*/
func OptimalApproach(a []int, d int) {
	n := len(a)
	slices.Reverse(a[:d])
	slices.Reverse(a[d:n])
	slices.Reverse(a[:n])
}
