package majorityelem

/*
Time Complexity: O(n^2)
Auxilary Space: O(1)
*/
func NaiveApproach(arr []int) []int {
	for i := range arr {
		l := []int{}
		temp := 1
		l = append(l, i)
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				temp++
				l = append(l, j)
			}
		}
		if temp > len(arr)/2 {
			return l
		}
	}
	return nil
}

/*
Time Complexity: O(n)
Auxilary Space: O(1)
*/
func OptimalApproach(arr []int) []int {
	res := 0
	count := 1

	for i := 1; i < len(arr); i++ {
		if arr[res] == arr[i] {
			count++
		} else {
			count--
		}

		if count == 0 {
			res = i
			count = 1
		}
	}

	l := []int{}
	count = 0

	for i := range arr {
		if arr[i] == arr[res] {
			count++
			l = append(l, i)
		}
	}

	if count > len(arr)/2 {
		return l
	}

	return nil
}
