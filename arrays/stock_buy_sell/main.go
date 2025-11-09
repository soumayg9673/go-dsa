package main

import "fmt"

/*
Time Complexity: Θ(n)
*/
func maxProfit(m []int) (p int) {
	for i := 1; i < len(m); i++ {
		if m[i] > m[i-1] {
			p += m[i] - m[i-1]
		}
	}
	return
}

func main() {
	data := [][]int{
		{1, 5, 3, 8, 12},   // 13
		{30, 20, 10},       // 0
		{10, 20, 30},       // 20
		{1, 5, 3, 1, 2, 8}, // 11
		{6, 5, 3, 1, 2, 8}, // 7
	}

	for _, d := range data {
		res := maxProfit(d)
		fmt.Println(res)
	}
}
