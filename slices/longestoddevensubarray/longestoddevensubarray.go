package longestoddevensubarray

/*
Algorith: Kadane's algorithm
=	Time Complexity: O(n)
-	Auxilary Space: O(1)
*/
func OptimalApproach(arr []int) int {
	res, temp := 1, 1

	for i := 1; i < len(arr); i++ {
		if arr[i-1]%2 != arr[i]%2 {
			temp++
			res = max(res, temp)
		} else {
			temp = 1
		}
	}

	return res
}

/*
func main() {
	data := [][]int{
		{10, 12, 14, 7, 8}, // 3
		{7, 10, 13, 14},    // 4
		{10, 12, 8, 4},     // 1
	}

	fmt.Println("Naive approach")
	for _, d := range data {
		fmt.Println(OptimalApproach(d))
	}

}
*/
