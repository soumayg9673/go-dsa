package twopointerappr_test

import (
	"fmt"
	"testing"

	"github.com/soumayg9673/go-dsa/search/twopointerappr"
)

var (
	testData = []struct {
		list     []int
		sum      int
		expected bool
	}{
		{
			list:     []int{2, 5, 8, 12, 30},
			sum:      17,
			expected: true,
		},
		{
			list:     []int{3, 8, 13, 18},
			sum:      14,
			expected: false,
		},
	}
)

func TestNaiveApproach(t *testing.T) {
	for _, data := range testData {
		name := fmt.Sprintf("Slice %v", data.list)
		t.Run(name, func(t *testing.T) {
			curr := twopointerappr.NaiveApproach(data.list, data.sum)
			if curr != data.expected {
				t.Errorf("expected %v but got %v", data.expected, curr)
			}
		})
	}
}
