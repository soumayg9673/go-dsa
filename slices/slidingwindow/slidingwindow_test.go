package slidingwindow_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/slidingwindow"
)

var (
	testData = []struct {
		list     []int
		k        int
		expected int
	}{
		{
			list:     []int{1, 8, 30, -5, 20, 7},
			k:        3,
			expected: 45,
		},
		{
			list:     []int{5, -10, 6, 90, 3},
			k:        2,
			expected: 96,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := slidingwindow.NaiveApproach(data.list, data.k)
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
			curr := slidingwindow.OptimalApproach(data.list, data.k)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
