package maxcircularsumsubarray_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/maxcircularsumsubarray"
)

var (
	testData = []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{10, 5, -5},
			expected: 15,
		},
		{
			list:     []int{5, -2, 3, 4},
			expected: 12,
		},
		{
			list:     []int{2, 3, -4},
			expected: 5,
		},
		{
			list:     []int{8, -4, 3, -4, 4},
			expected: 12,
		},
		{
			list:     []int{-3, 4, 6, -2},
			expected: 10,
		},
		{
			list:     []int{-8, 7, 6},
			expected: 13,
		},
		{
			list:     []int{3, -4, 5, 6, -8, 7},
			expected: 17,
		},
		{
			list:     []int{-5, -3},
			expected: -3,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := maxcircularsumsubarray.NaiveApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := maxcircularsumsubarray.OptimalApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
