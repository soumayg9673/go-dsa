package maxcons1_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/maxcons1"
)

var (
	testData = []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{0, 1, 1, 0, 1, 0},
			expected: 2,
		},
		{
			list:     []int{1, 1, 1, 1},
			expected: 4,
		},
		{
			list:     []int{0, 0, 0},
			expected: 0,
		},
		{
			list:     []int{1, 0, 1, 1, 1, 1, 0, 1, 1},
			expected: 4,
		},
	}
)

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := maxcons1.OptimalApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
