package main

import (
	"fmt"
	"slices"
)

/*
To find largest element in slice using standard library

Time complexity:
	- Best/Worst case: Θ(n)
*/

func main() {
	data := [][]int{
		{1, 2, 3, 4, 5},
		{10, 9, 8, 7, 6},
	}

	for _, d := range data {
		l1 := slices.Max(d) // returns the maximum element in slice
		fmt.Println("largest element in array: ", l1)
	}
}
