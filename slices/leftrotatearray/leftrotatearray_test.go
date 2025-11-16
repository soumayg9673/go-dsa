package leftrotatearray_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/leftrotatearray"
)

func TestNaiveApproach(t *testing.T) {
	testData := []struct {
		list      []int
		rotateByD int
		expected  []int
	}{
		{
			list:      []int{1, 2, 3, 4, 5},
			rotateByD: 2,
			expected:  []int{3, 4, 5, 1, 2},
		},
		{
			list:      []int{10, 9, 8, 7, 6},
			rotateByD: 1,
			expected:  []int{9, 8, 7, 6, 10},
		},
	}

	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			leftrotatearray.NaiveApproach(data.list, data.rotateByD)
			if !slices.Equal(data.list, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, data.list)
			}
		})
	}
}

func TestOptimalApproach(t *testing.T) {
	testData := []struct {
		list      []int
		rotateByD int
		expected  []int
	}{
		{
			list:      []int{1, 2, 3, 4, 5},
			rotateByD: 2,
			expected:  []int{3, 4, 5, 1, 2},
		},
		{
			list:      []int{10, 9, 8, 7, 6},
			rotateByD: 1,
			expected:  []int{9, 8, 7, 6, 10},
		},
	}

	for _, data := range testData {
		name := fmt.Sprintf("Slice: %v", data.list)
		t.Run(name, func(t *testing.T) {
			leftrotatearray.OptimalApproach(data.list, data.rotateByD)
			if !slices.Equal(data.list, data.expected) {
				t.Errorf("expected %v but got %v", data.expected, data.list)
			}
		})
	}
}
