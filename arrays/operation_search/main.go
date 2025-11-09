package main

import "fmt"

/*
Array/Slice operation: Search
	- Best case: O(log(n)) using Binary search (only if the array is sorted)
	- Worst case: O(n)
*/

// Find the element by index. If no element found return -1
func search[T int](a [5]T, n T) int {
	for i, v := range a {
		if v == n {
			return i
		}
	}
	return -1
}

func main() {
	data := [2][5]int{
		{20, 5, 7, 25, 12},
		{1, 2, 3, 4, 5},
	}

	for _, d := range data {
		fmt.Println(search(d, 5))
	}
}
