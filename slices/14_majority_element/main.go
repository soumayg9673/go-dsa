package main

import (
	"fmt"
)

/*
Time Complexity: O(n^2)
Auxilary Space: O(1)
*/
func naiveApproach(arr []int) []int {
	l := []int{}
	for i := range arr {
		temp := 1
		l = append(l, i)
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				temp++
				l = append(l, j)
			}
		}
		if temp > len(arr)/2 {
			return l
		}
		l = []int{}
	}
	return l
}

/*
Time Complexity: O(n)
Auxilary Space: O(1)
*/
func optimalSolution(arr []int) []int {
	res := 0
	count := 1

	for i := 1; i < len(arr); i++ {
		if arr[res] == arr[i] {
			count++
		} else {
			count = 0
		}

		if count == 0 {
			res = i
			count = 1
		}
	}

	l := []int{}
	count = 0

	for i := range arr {
		if arr[i] == arr[res] {
			count++
			l = append(l, i)
		}
	}

	if count > len(arr)/2 {
		return l
	} else {
		return []int{}
	}
}

func main() {
	data := [][]int{
		{8, 3, 4, 8, 8},              // 0 or 3 or 4
		{3, 7, 4, 7, 7, 5},           // -1
		{20, 30, 40, 50, 50, 50, 50}, // 3 or 4 or 5 or 6
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
