package main

import (
	"fmt"
	"slices"
)

func leaders[T []int](a T) (res T) {
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

func main() {
	data := [][]int{
		{7, 10, 4, 3, 6, 5, 2},
		{10, 20, 30},
		{30, 20, 10},
	}

	for _, d := range data {
		res := leaders(d)
		fmt.Println(res)
	}
}
