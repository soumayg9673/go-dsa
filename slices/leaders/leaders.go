package leaders

import (
	"slices"
)

func OptimalApproach[T []int](a T) (res T) {
	temp := len(a) - 1
	res = append(res, a[temp])

	for i := len(a) - 2; i >= 0; i-- {
		if a[temp] < a[i] {
			temp = i
			res = append(res, a[temp])
		}
	}

	slices.Reverse(res)
	return
}
