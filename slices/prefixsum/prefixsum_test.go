package prefixsum_test

import (
	"fmt"
	"slices"
	"testing"

	prefixsum "github.com/soumayg9673/go-dsa/slices/prefixsum"
)

var (
	testData = []struct {
		list     prefixsum.Data
		expected []int
	}{
		{
			list: prefixsum.Data{
				List: []int{2, 8, 3, 9, 6, 5, 4},
				Query: []struct {
					Start int
					End   int
				}{
					{Start: 0, End: 2},
					{Start: 1, End: 3},
					{Start: 2, End: 6},
				},
			},
			expected: []int{13, 20, 27},
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list.List)
		t.Run(name, func(t *testing.T) {
			curr := data.list.NaiveApproach()
			if !slices.Equal(curr, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list.List)
		t.Run(name, func(t *testing.T) {
			curr := data.list.OptimalApproach()
			if !slices.Equal(curr, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
