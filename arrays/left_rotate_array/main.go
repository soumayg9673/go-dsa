package main

import (
	"fmt"
	"slices"
)

/*
Time Complexity: Θ(n)
Auxilary Space:
  - Worst case: O(n)
  - Best case: Θ(1)
*/
func leftRotate1(a []int, d int) {
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
Auxilary Space:
  - Best/Worst case: Θ(1)
*/
func leftRotate(a []int, d int) {
	n := len(a)
	slices.Reverse(a[:d])
	slices.Reverse(a[d:n])
	slices.Reverse(a[:n])
}

func main() {
	data1 := [][]int{
		{1, 2, 3, 4, 5},
		{10, 9, 8, 7, 6},
	}

	for _, d := range data1 {
		leftRotate1(d, 2)
		fmt.Println(d)
	}

	data2 := [][]int{
		{1, 2, 3, 4, 5},
		{10, 9, 8, 7, 6},
	}

	for _, d := range data2 {
		leftRotate(d, 2)
		fmt.Println(d)
	}

}
