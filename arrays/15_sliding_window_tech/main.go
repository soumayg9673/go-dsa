package main

import "fmt"

/*
Time Complexity: O(n^2)
Auxilary Space: O(1)
*/
func naiveApproach(arr []int, k int) int {
	res := arr[k-1]

	for i := 0; i < len(arr) && i+k-1 < len(arr); i++ {
		temp := 0
		for j := 0; j < k; j++ {
			temp += arr[i+j]
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
}
