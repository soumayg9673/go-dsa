package main

import "fmt"

func frequencySorted(a []int) {
	n := len(a)
	f, i := 1, 1
	for ; i < n; i++ {
		for ; i < n && a[i] == a[i-1]; i, f = i+1, f+1 {
		}
		fmt.Printf("%v occurs %v times in array %v\n", a[i-1], f, a)
		f = 1
	}
	if n == 1 || a[n-1] != a[n-2] {
		fmt.Printf("%v occurs %v times in array %v\n", a[n-1], 1, a)
	}
}

func frequencyUnsorted(a []int) {
	m := map[int]int{}
	for _, v := range a {
		m[v] += 1
	}
	fmt.Println(m)
}

func main() {
	data := [][]int{
		{10, 10, 10, 25, 30, 30},
		{10, 10, 10, 10},
		{10, 20, 30},
	}

	for _, d := range data {
		frequencySorted(d)
	}

	data2 := [][]int{
		{10, 25, 10, 30, 30, 10},
		{10, 10, 10, 10},
		{10, 30, 20},
	}

	for _, d := range data2 {
		frequencyUnsorted(d)
	}
}
