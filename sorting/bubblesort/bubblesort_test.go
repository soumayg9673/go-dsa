package bubblesort_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/soumayg9673/go-dsa/sorting/bubblesort"
)

var (
	testData = []struct {
		list     []int
		expected []int
	}{
		{
			list:     []int{1, 7, 3, 2, 6, 9, 8},
			expected: []int{1, 2, 3, 6, 7, 8, 9},
		},
		{
			list:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			list:     []int{11, 12, 13, 14, 15},
			expected: []int{11, 12, 13, 14, 15},
		},
	}
)

func TestBubbleSort(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := bubblesort.BubbleSort(data.list)
			if !slices.Equal(curr, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
