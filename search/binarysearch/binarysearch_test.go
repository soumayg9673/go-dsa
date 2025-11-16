package binarysearch_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/search/binarysearch"
)

var (
	testData = []struct {
		list        []int
		search      int
		expectedIdx int
	}{
		{
			list:        []int{10, 20, 30, 40, 50, 60},
			search:      20,
			expectedIdx: 1,
		},
		{
			list:        []int{5, 15, 25},
			search:      25,
			expectedIdx: 2,
		},
		{
			list:        []int{5, 10, 15, 25, 35},
			search:      20,
			expectedIdx: -1,
		},
		{
			list:        []int{10, 15},
			search:      20,
			expectedIdx: -1,
		},
		{
			list:        []int{10, 15},
			search:      5,
			expectedIdx: -1,
		},
		{
			list:        []int{10, 10},
			search:      10,
			expectedIdx: 1, // 0 or 1
		},
	}
)

func TestIterative(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := binarysearch.Iterative(data.list, data.search)
			if curr != data.expectedIdx {
				t.Errorf("expected %v but got %v", data.expectedIdx, curr)
			}
		})
	}
}

func TestRecursive(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := binarysearch.Recursive(data.list, data.search, 0, len(data.list))
			if curr != data.expectedIdx {
				t.Errorf("expected %v but got %v", data.expectedIdx, curr)
			}
		})
	}
}
