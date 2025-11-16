package prefixsum

type Data struct {
	List  []int
	Query []struct {
		Start int
		End   int
	}
}

/*
Time Complexity: O(nm)
Auxilary Space: O(1)
*/
func (d Data) NaiveApproach() []int {
	list := make([]int, 0, len(d.Query))
	for _, q := range d.Query {
		sum := 0
		for i := q.Start; i <= q.End; i++ {
			sum += d.List[i]
		}
		list = append(list, sum)
	}
	return list
}

/*
Time Complexity: O(n)
Auxilary Space: O(n)
*/
func (d Data) OptimalApproach() []int {

	// Calculate cummulative sum from start
	leftSum := make([]int, len(d.List))
	leftSum[0] = d.List[0]
	for i := 1; i < len(d.List); i++ {
		leftSum[i] = leftSum[i-1] + d.List[i]
	}

	list := make([]int, 0, len(d.Query))
	for _, q := range d.Query {
		if q.Start == 0 {
			list = append(list, leftSum[q.End])
		} else {
			list = append(list, leftSum[q.End]-leftSum[q.Start-1])
		}
	}

	return list
}
