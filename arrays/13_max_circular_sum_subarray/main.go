package main

import (
	"fmt"
	"slices"
)

func rotateArray(arr []int) {
	n := len(arr)
	slices.Reverse(arr[:1])
	slices.Reverse(arr[1:n])
	slices.Reverse(arr[:n])
}

func calculateMaxSubArraySum(arr []int) int {
	n := len(arr)
	res := arr[0]
	curr := arr[0]

	for i := 1; i < n; i++ {
		curr = max(curr+arr[i], arr[i])
		res = max(res, curr)
	}

	return res
}

/*
Time Complexity: O(n^2)
Auxilary Space: O(1)
*/
func naiveApproach1(arr []int) int {
	res := arr[0]

	for range arr {
		rotateArray(arr)
		m := calculateMaxSubArraySum(arr)
		res = max(m, res)
	}

	return res
}

/*
Time Complexity: O(n^2)
Auxilary Space: O(1)
*/
func naiveApproach2(arr []int) int {
	res := arr[0]

	for i := range arr {
		currMax := arr[i]
		currSum := arr[i]
		for j := 1; j < len(arr); j++ {
			idx := (i + j) % len(arr)
			currSum += arr[idx]
			currMax = max(currMax, currSum)
		}
		res = max(currMax, res)
	}

	return res
}

func main() {
	data := [][]int{
		{10, 5, -5},          // 15
		{5, -2, 3, 4},        // 12
		{2, 3, -4},           // 5
		{8, -4, 3, -4, 4},    // 12
		{-3, 4, 6, -2},       // 10
		{-8, 7, 6},           // 13
		{3, -4, 5, 6, -8, 7}, // 17
	}

	fmt.Println("Naive approach using left rotation and max subarray sum")
	for _, d := range data {
		fmt.Println(naiveApproach1(d))
	}

	fmt.Println("Naive approach")
	for _, d := range data {
		fmt.Println(naiveApproach2(d))
	}
}
