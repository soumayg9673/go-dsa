package main

import (
	"fmt"
)

/*
Time Complexity: Θ(n)
Auxilary Space: Θ(n)
*/
func calculateWaterCollectN(w []int) (vol int) {
	n := len(w)

	lMax := make([]int, n)
	lMax[0] = w[0]
	for i := 1; i < n; i++ {
		lMax[i] = max(w[i], lMax[i-1])
	}

	rMax := make([]int, n)
	rMax[n-1] = w[n-1]
	for i := n - 2; i >= 0; i-- {
		rMax[i] = max(w[i], rMax[i+1])
	}

	for i := 1; i < n-1; i++ {
		vol += min(lMax[i], rMax[i]) - w[i]
	}
	return
}

func main() {
	data := [][]int{
		{2, 0, 2},                            // 2
		{3, 0, 1, 2, 5},                      // 6
		{1, 2, 3},                            // 0
		{3, 2, 1},                            // 0
		{3, 0, 2},                            // 2
		{4, 0, 5, 1},                         // 4
		{5, 1, 7, 0, 10},                     // 11
		{5, 1, 7, 0, 4},                      // 8
		{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, // 6
	}

	for _, d := range data {
		vol := calculateWaterCollectN(d)
		fmt.Println(vol)
	}

}
