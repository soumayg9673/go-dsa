package maxsubarraysum

/*
Time Complexity: O(n^2)
Auxilary Space: Θ(1)
*/
func NaiveApproach(arr []int) int {
	res := arr[0]
	for i := range arr {
		temp := 0
		for j := i; j < len(arr); j++ {
			temp += arr[j]
			res = max(res, temp, arr[i])
		}
	}
	return res
}

/*
Time Complexity: O(n)
Auxilary Space: Θ(1)
*/
func OptimalApproach(arr []int) int {
	curr := arr[0]
	res := arr[0]

	for i := 1; i < len(arr); i++ {
		curr = max(arr[i], curr+arr[i])
		res = max(res, curr)
	}

	return res
}
