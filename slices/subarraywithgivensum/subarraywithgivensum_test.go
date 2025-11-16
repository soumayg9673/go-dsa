package subarraywithgivensum_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/subarraywithgivensum"
)

var (
	testData = []struct {
		list     []int
		sum      int
		expected bool
	}{
		{
			list:     []int{1, 4, 20, 3, 10, 5},
			sum:      33,
			expected: true,
		},
		{
			list:     []int{1, 4, 0, 0, 3, 10, 5},
			sum:      7,
			expected: true,
		},
		{
			list:     []int{2, 4},
			sum:      3,
			expected: false,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := subarraywithgivensum.NaiveApproach(data.list, data.sum)
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
			curr := subarraywithgivensum.OptimalApproach(data.list, data.sum)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
