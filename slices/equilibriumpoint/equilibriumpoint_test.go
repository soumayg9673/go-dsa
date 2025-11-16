package equilibriumpoint_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/slices/equilibriumpoint"
)

var (
	testData = []struct {
		list     []int
		expected bool
	}{
		{list: []int{3, 4, 8, -9, 20, 6}, expected: true},
		{list: []int{4, 2, -2}, expected: true},
		{list: []int{4, 2, 2}, expected: false},
		{list: []int{0, 0, 2}, expected: true},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Array: %v", data.list)
		t.Run(name, func(t *testing.T) {
			val := equilibriumpoint.NaiveApproach(data.list)
			if val != data.expected {
				t.Errorf("Expected %v but got %v", data.expected, val)
			}
		})
	}
}

func TestOptimalSolution(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Array: %v", data.list)
		t.Run(name, func(t *testing.T) {
			val := equilibriumpoint.OptimalSolution(data.list)
			if val != data.expected {
				t.Errorf("Expected %v but got %v", data.expected, val)
			}
		})
	}
}
