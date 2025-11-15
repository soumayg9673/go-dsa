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
}
