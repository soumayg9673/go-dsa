package main

import "fmt"

/*
Sorting: Bubble Sort
Time Complexity:
	- Best case: O(n)
	- Worst case: O(n^2)
*/

func bubbleSort(a []int) {
	swapping := true
	end := len(a)
	for swapping {
		swapping = false
		for i := 1; i < end; i++ {
			if a[i-1] > a[i] {
				a[i-1], a[i] = a[i], a[i-1]
				swapping = true
			}
		}
		end--
	}
}

func main() {
	data := [][]int{
		{1, 2, 3, 4, 5},
		{10, 9, 8, 7, 6, 5},
		{13, 12, 15, 11, 14},
	}

	for _, d := range data {
		bubbleSort(d)
		fmt.Println(d)
	}
}
