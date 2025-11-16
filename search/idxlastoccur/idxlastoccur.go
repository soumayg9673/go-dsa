package idxlastoccur

// Time Complexity: O(n)
func NaiveApproach(arr []int, s int) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == s {
			return i
		}
	}
	return -1
}

// Time Complexity: O(log(n))
// Auxilary Space: O(log(n))
func OptimalApproachRecursive(arr []int, s, low, high int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] > s {
		return OptimalApproachRecursive(arr, s, low, mid-1)
	} else if arr[mid] < s {
		return OptimalApproachRecursive(arr, s, mid+1, high)
	} else {
		if mid == len(arr)-1 || arr[mid+1] != arr[mid] {
			return mid
		} else {
			return OptimalApproachRecursive(arr, s, mid+1, high)
		}
	}
}

// Time Complexity: O(log(n))
// Auxilary Space: O(1)
func OptimalApproachIterative(arr []int, s int) int {
	low := 0
	high := len(arr)
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > s {
			high = mid - 1
		} else if arr[mid] < s {
			low = mid + 1
		} else {
			if mid == len(arr)-1 || arr[mid+1] != arr[mid] {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}
