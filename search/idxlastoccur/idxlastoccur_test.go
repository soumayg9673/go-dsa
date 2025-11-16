package idxlastoccur_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/search/idxlastoccur"
)

var (
	testData = []struct {
		list        []int
		search      int
		expectedIdx int
	}{
		{
			list:        []int{1, 10, 10, 10, 20, 20, 40},
			search:      20,
			expectedIdx: 5,
		},
		{
			list:        []int{10, 20, 30},
			search:      15,
			expectedIdx: -1,
		},
		{
			list:        []int{15, 15, 15},
			search:      15,
			expectedIdx: 2,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := idxlastoccur.NaiveApproach(data.list, data.search)
			if curr != data.expectedIdx {
				t.Errorf("expected %v but got %v", data.expectedIdx, curr)
			}
		})
	}
}

func TestOptimalApproachRecursive(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := idxlastoccur.OptimalApproachRecursive(data.list, data.search, 0, len(data.list))
			if curr != data.expectedIdx {
				t.Errorf("expected %v but got %v", data.expectedIdx, curr)
			}
		})
	}
}

func TestOptimalApproachIterative(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := idxlastoccur.OptimalApproachIterative(data.list, data.search)
			if curr != data.expectedIdx {
				t.Errorf("expected %v but got %v", data.expectedIdx, curr)
			}
		})
	}
}
