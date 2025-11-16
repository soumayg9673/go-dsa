package main

import (
	"fmt"
	"math"
)

/*
Time Complexity: O(nk)
Auxilary Space: O(1)
*/
func naiveApproach(arr []int, k int) int {
	res := math.MinInt

	for i := 0; i < len(arr) && i+k-1 < len(arr); i++ {
		temp := 0
		for j := range k {
			temp += arr[i+j]
		}
		res = max(res, temp)
	}

	return res
}

/*
Time Complexity: O(n)
Auxilary Space: O(1)
*/
func optimalSolution(arr []int, k int) int {
	res := math.MinInt
	temp := 0

	for i := 0; i+k-1 < len(arr); i++ {
		if i == 0 {
			for j := range k {
				temp += arr[j]
			}
		} else {
			temp = temp - arr[i-1] + arr[i+k-1]
		}
		res = max(res, temp)
	}

	return res
}

func main() {
	data := []struct {
		l []int
		k int
	}{
		{
			l: []int{1, 8, 30, -5, 20, 7},
			k: 3,
		}, // 45
		{
			l: []int{5, -10, 6, 90, 3},
			k: 2,
		}, // 96
	}

	fmt.Println("Naive approach")
	for _, d := range data {
		fmt.Println(naiveApproach(d.l, d.k))
	}

	fmt.Println("Optimal approach")
	for _, d := range data {
		fmt.Println(optimalSolution(d.l, d.k))
	}
}
