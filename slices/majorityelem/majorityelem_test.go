package majorityelem_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/majorityelem"
)

var (
	testData = []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{8, 8, 8, 3, 4},
			expected: []int{0, 1, 2},
		},
		{
			list:     []int{3, 7, 4, 7, 7, 5},
			expected: nil,
		},
		{
			list:     []int{20, 30, 40, 50, 50, 50, 50},
			expected: []int{3, 4, 5, 6},
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := majorityelem.NaiveApproach(data.list)
			if !slices.Equal(curr, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := majorityelem.OptimalApproach(data.list)
			if !slices.Equal(curr, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
