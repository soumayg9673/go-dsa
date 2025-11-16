package main

import (
	"fmt"
	"slices"
)

/*
To reverse the slice
1. Use standard library

Time Complexity: Θ(n)
Auxilary Space: Θ(1)

*/

func main() {
	data := [][]int{
		{1, 2, 3, 4, 5},
		{10, 9, 8, 7, 6},
	}

	for _, d := range data {
		slices.Reverse(d)
		fmt.Println(d)
	}
}
