package maxslice

import (
	"fmt"
	"testing"
)

func TestFindMaximumProfit(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{23171, 21011, 21123, 21366, 21013, 21367},
			356,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findMaximumProfit(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
