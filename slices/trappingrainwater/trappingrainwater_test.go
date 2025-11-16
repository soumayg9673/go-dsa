package trappingrainwater_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/trappingrainwater"
)

var (
	testData = []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{2, 0, 2},
			expected: 2,
		},
		{
			list:     []int{3, 0, 1, 2, 5},
			expected: 6,
		},
		{
			list:     []int{1, 2, 3},
			expected: 0,
		},
		{
			list:     []int{3, 2, 1},
			expected: 0,
		},
		{
			list:     []int{3, 0, 2},
			expected: 2,
		},
		{
			list:     []int{4, 0, 5, 1},
			expected: 4,
		},
		{
			list:     []int{5, 1, 7, 0, 10},
			expected: 11,
		},
		{
			list:     []int{5, 1, 7, 0, 4},
			expected: 8,
		},
		{
			list:     []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
	}
)

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := trappingrainwater.OptimalApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
