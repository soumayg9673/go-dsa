package main

import "fmt"

type Data struct {
	list  []int
	query []struct {
		start int
		end   int
	}
}

/*
Time Complexity: O(nm)
Auxilary Space: O(1)
*/
func (d Data) naiveApproach() []int {
	list := make([]int, 0, len(d.query))
	for _, q := range d.query {
		sum := 0
		for i := q.start; i <= q.end; i++ {
			sum += d.list[i]
		}
		list = append(list, sum)
	}
	return list
}

func main() {
	data := []Data{
		{
			list: []int{2, 8, 3, 9, 6, 5, 4},
			query: []struct {
				start int
				end   int
			}{
				{start: 0, end: 2},
				{start: 1, end: 3},
				{start: 2, end: 6},
			},
		},
	}

	for _, d := range data {
		fmt.Println(d.naiveApproach())
	}
}
