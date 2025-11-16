package stockbuysell_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/stockbuysell"
)

var (
	testData = []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{1, 5, 3, 8, 12},
			expected: 13,
		},
		{
			list:     []int{30, 20, 10},
			expected: 0,
		},
		{
			list:     []int{10, 20, 30},
			expected: 20,
		},
		{
			list:     []int{1, 5, 3, 1, 2, 8},
			expected: 11,
		},
		{
			list:     []int{6, 5, 3, 1, 2, 8},
			expected: 7,
		},
	}
)

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := stockbuysell.OptimalApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
