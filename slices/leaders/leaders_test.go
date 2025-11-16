package leaders_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/leaders"
)

var (
	testData = []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{7, 10, 4, 3, 6, 5, 2},
			expected: []int{10, 6, 5, 2},
		},
		{
			list:     []int{10, 20, 30},
			expected: []int{30},
		},
		{
			list:     []int{30, 20, 10},
			expected: []int{30, 20, 10},
		},
	}
)

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := leaders.OptimalApproach(data.list)
			if !slices.Equal(curr, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
