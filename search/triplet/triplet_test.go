package triplet_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/search/triplet"
)

var (
	testData = []struct {
		list     []int
		sum      int
		expected bool
	}{
		{
			list:     []int{2, 3, 4, 8, 9, 20, 40},
			sum:      32,
			expected: true,
		},
		{
			list:     []int{1, 2, 5, 6},
			sum:      14,
			expected: false,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := triplet.NaiveApproach(data.list, data.sum)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
