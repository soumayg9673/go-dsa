package main

import (
	"fmt"
)

func maxCons(arr []int) (res int) {
	t := 0
	for i := range arr {
		if arr[i] == 1 {
			t++
			res = max(res, t)
		} else {
			t = 0
		}
	}
	return max(res, t)
}

func main() {
	data := [][]int{
		{0, 1, 1, 0, 1, 0},
		{1, 1, 1, 1},
		{0, 0, 0},
		{1, 0, 1, 1, 1, 1, 0, 1, 1},
	}

	for _, d := range data {
		fmt.Println(maxCons(d))
	}
}
