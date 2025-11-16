package subarraywithgivensum

/*
Time Complexity: O(n^2)
Auxilary Space: O(1)
*/
func NaiveApproach(arr []int, s int) bool {
	for i := range arr {
		temp := s - arr[i]
		for j := i + 1; j < len(arr); j++ {
			temp -= arr[j]
			if temp == 0 {
				return true
			}
		}
	}
	return false
}

/*
Time Complexity: O(n)
Auxilary Space: O(1)
*/
func OptimalApproach(arr []int, sum int) bool {
	temp := 0

	for s, e := 0, 0; s < len(arr) && e < len(arr); {
		if temp == sum {
			return true
		} else if temp > sum {
			temp -= arr[s]
			s++
		} else {
			temp += arr[e]
			e++
		}
	}

	return false
}
