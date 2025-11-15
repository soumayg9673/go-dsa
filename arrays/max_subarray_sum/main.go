package main

import "fmt"

func naiveApproach(arr []int) int {
	res := arr[0]

	for i := range arr {
		temp := 0
		for j := i; j < len(arr); j++ {
			temp += arr[j]
			res = max(res, temp, arr[i])
		}
	}
	return res
}

func main() {
	data := [][]int{
		{2, 3, -8, 7, -1, 2, 3},
		{5, 8, 3},
		{-6, -1, -7},
	}

	fmt.Println("Naive approach")
	for _, d := range data {
		fmt.Println(naiveApproach(d))
	}
}
