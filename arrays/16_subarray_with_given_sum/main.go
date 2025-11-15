package main

import "fmt"

/*
Time Complexity: O(n^2)
Auxilary Space: O(1)
*/
func naiveApproach(arr []int, s int) bool {
	for i := range arr {
		temp := s - arr[i]
		for j := i + 1; j < len(arr); j++ {
			temp -= arr[j]
			if temp == 0 {
				return true
			}
		}
	}
	return false
}

/*
Time Complexity: O(n)
Auxilary Space: O(1)
*/
func optimalSolution(arr []int, sum int) bool {
	temp := 0

	for s, e := 0, 0; s < len(arr) && e < len(arr); {
		if temp == sum {
			return true
		} else if temp > sum {
			temp -= arr[s]
			s++
		} else {
			temp += arr[e]
			e++
		}
	}

	return false
}

func main() {
	data := []struct {
		l []int
		s int
	}{
		{l: []int{1, 4, 20, 3, 10, 5}, s: 33},  // true
		{l: []int{1, 4, 0, 0, 3, 10, 5}, s: 7}, // true
		{l: []int{2, 4}, s: 3},                 // false
	}

	fmt.Println("Naive approach")
	for _, d := range data {
		fmt.Println(naiveApproach(d.l, d.s))
	}

	fmt.Println("Optimal approach")
	for _, d := range data {
		fmt.Println(optimalSolution(d.l, d.s))
	}
}
