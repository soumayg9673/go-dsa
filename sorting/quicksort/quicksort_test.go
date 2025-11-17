package quicksort_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/soumayg9673/go-dsa/sorting/quicksort"
)

var (
	testData = []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{1, 7, 3, 2, 6, 9, 8},
			expected: []int{1, 2, 3, 6, 7, 8, 9},
		},
		{
			list:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			list:     []int{11, 12, 13, 14, 15},
			expected: []int{11, 12, 13, 14, 15},
		},
	}
)

func TestLomuto(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			quicksort.Lomuto(data.list, 0, len(data.list)-1)
			if !slices.Equal(data.list, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, data.list)
			}
		})
	}
}
