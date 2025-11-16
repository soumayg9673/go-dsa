package frequency

import (
	"fmt"
)

func OptimalApproach(a []int) map[int]int {
	m := map[int]int{}
	n := len(a)
	f, i := 1, 1
	for ; i < n; i++ {
		for ; i < n && a[i] == a[i-1]; i, f = i+1, f+1 {
		}
		fmt.Printf("%v occurs %v times in array %v\n", a[i-1], f, a)
		m[a[i-1]] = f
		f = 1
	}
	if n == 1 || a[n-1] != a[n-2] {
		fmt.Printf("%v occurs %v times in array %v\n", a[n-1], 1, a)
		m[a[i-1]] = f
	}

	return m
}
