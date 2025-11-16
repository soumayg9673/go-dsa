package maxcons1

func OptimalApproach(arr []int) (res int) {
	t := 0
	for i := range arr {
		if arr[i] == 1 {
			t++
			res = max(res, t)
		} else {
			t = 0
		}
	}
	return max(res, t)
}

// func main() {
// 	data := [][]int{
// 		{0, 1, 1, 0, 1, 0},          // 2
// 		{1, 1, 1, 1},                // 4
// 		{0, 0, 0},                   // 0
// 		{1, 0, 1, 1, 1, 1, 0, 1, 1}, // 4
// 	}

// 	for _, d := range data {
// 		fmt.Println(maxCons(d))
// 	}
// }
