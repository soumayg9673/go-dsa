package longestoddevensubarray_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/longestoddevensubarray"
)

var (
	testData = []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{10, 12, 14, 7, 8},
			expected: 3,
		},
		{
			list:     []int{7, 10, 13, 14},
			expected: 4,
		},
		{
			list:     []int{10, 12, 8, 4},
			expected: 1,
		},
	}
)

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := longestoddevensubarray.OptimalApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
