package main

import "fmt"

/*
Time Complexity: O(n^2)
Auxilary Space: Θ(1)
*/
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

/*
Time Complexity: O(n)
Auxilary Space: Θ(1)
*/
func optimalSolution(arr []int) int {
	curr := arr[0]
	res := arr[0]

	for i := 1; i < len(arr); i++ {
		curr = max(arr[i], curr+arr[i])
		res = max(res, curr)
	}

	return res
}

func main() {
	data := [][]int{
		{2, 3, -8, 7, -1, 2, 3},   // 11
		{5, 8, 3},                 // 16
		{-6, -1, -7},              // -1
		{-5, 1, -2, 3, -1, 2, -2}, // 4
	}

	fmt.Println("Naive approach")
	for _, d := range data {
		fmt.Println(naiveApproach(d))
	}

	fmt.Println("Optimal approach")
	for _, d := range data {
		fmt.Println(optimalSolution(d))
	}
}
