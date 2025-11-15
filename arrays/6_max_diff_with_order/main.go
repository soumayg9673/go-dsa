package main

import "fmt"

/*
Maximum value of arr[j] - arr[i], where j > i
*/

/*
Time Complexity: O(n^2)
Auxilary Space: Θ(1)
*/
func maxDiff1(a []int) (res int) {
	if len(a) < 2 {
		panic("bad input")
	}
	if len(a) == 2 {
		return a[1] - a[0]
	}

	res = a[1] - a[0]
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			res = max(res, a[j]-a[i])
		}
	}
	return
}

/*
Time Complexity: Θ(n)
Auxilary Space: Θ(1)
*/
func maxDiff(a []int) (res int) {
	res = a[1] - a[0]
	m := a[0]
	for i := 1; i < len(a); i++ {
		res = max(res, a[i]-m)
		m = min(m, a[i])
	}
	return
}

func main() {
	data1 := [][]int{
		{2, 3, 10, 6, 4, 8, 1}, // 8
		{7, 9, 5, 6, 3, 2},     // 2
		{10, 20, 30},           // 20
		{30, 10, 8, 2},         // -2
	}

	for _, d := range data1 {
		res := maxDiff1(d)
		fmt.Println(res)
	}

	data2 := [][]int{
		{2, 3, 10, 6, 4, 8, 1}, // 8
		{7, 9, 5, 6, 3, 2},     // 2
		{10, 20, 30},           // 20
		{30, 10, 8, 2},         // -2
	}

	for _, d := range data2 {
		res := maxDiff(d)
		fmt.Println(res)
	}
}
