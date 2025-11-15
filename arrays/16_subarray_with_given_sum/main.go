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
}
