package count1s_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/search/count1s"
)

var (
	testData = []struct {
		list     []int
		expected int
	}{
		{
			list:     []int{0, 0, 0, 1, 1, 1, 1},
			expected: 4,
		},
		{
			list:     []int{1, 1, 1, 1},
			expected: 4,
		},
		{
			list:     []int{0, 0, 0, 0},
			expected: 0,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := count1s.NaiveApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}

func TestOptimalApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := count1s.OptimalApproach(data.list)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
