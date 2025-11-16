package main

import (
	"fmt"
	"slices"
)

/*
To find largest element in slice
1. Find largest element -> using standard library
2. Find index of largest element -> using function

Time complexity:
	- Best/Worst case: Θ(n)
*/

func largest(a []int) (temp, idx int) {
	if len(a) == 0 {
		return 0, -1
	}

	temp = a[idx]
	for i := 1; i < len(a); i++ {
		if temp < a[i] {
			temp = a[i]
			idx = i
		}
	}

	return temp, idx
}

func main() {
	data := [][]int{
		{1, 2, 3, 4, 5},
		{10, 9, 8, 7, 6},
	}

	for _, d := range data {
		// Using standard library
		l1 := slices.Max(d) // returns the maximum element in slice
		fmt.Println("largest element in array: ", l1)

		l2, idx := largest(d) // returns the index
		if idx < 0 {
			fmt.Println("empty array")
		}
		fmt.Printf("largest element: %v is present at index: %v in array\n", l2, idx)
	}
}
