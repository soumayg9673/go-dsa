package frequency_test

import (
	"fmt"
	"maps"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/frequency"
)

var (
	testData = []struct {
		list     []int
		expected map[int]int
	}{
		{
			list: []int{10, 10, 10, 25, 30, 30},
			expected: map[int]int{
				10: 3,
				25: 1,
				30: 2,
			},
		},
		{
			list: []int{10, 10, 10, 10},
			expected: map[int]int{
				10: 4,
			},
		},
		{
			list: []int{10, 20, 30},
			expected: map[int]int{
				10: 1,
				20: 1,
				30: 1,
			},
		},
	}
)

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := frequency.OptimalApproach(data.list)
			if !maps.Equal(curr, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
