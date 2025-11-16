package maxdiffwithorder_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/maxdiffwithorder"
)

var (
	testData = []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{2, 3, 10, 6, 4, 8, 1},
			expected: 8,
		},
		{
			list:     []int{7, 9, 5, 6, 3, 2},
			expected: 2,
		},
		{
			list:     []int{10, 20, 30},
			expected: 20,
		},
		{
			list:     []int{30, 10, 8, 2},
			expected: -2,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := maxdiffwithorder.NaiveApproach(data.list)
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
			curr := maxdiffwithorder.OptimalApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
