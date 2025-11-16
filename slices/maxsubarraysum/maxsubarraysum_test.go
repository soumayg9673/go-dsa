package maxsubarraysum_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/maxsubarraysum"
)

var (
	testData = []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{2, 3, -8, 7, -1, 2, 3},
			expected: 11,
		},
		{
			list:     []int{5, 8, 3},
			expected: 16,
		},
		{
			list:     []int{-6, -1, -7},
			expected: -1,
		},
		{
			list:     []int{-5, 1, -2, 3, -1, 2, -2},
			expected: 4,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := maxsubarraysum.NaiveApproach(data.list)
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
			curr := maxsubarraysum.OptimalApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
