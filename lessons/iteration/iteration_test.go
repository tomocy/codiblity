package iteration

import (
	"fmt"
	"testing"
)

func TestFindLongestBinaryGap(t *testing.T) {
	tests := []struct {
		input, expected int
	}{
		{9, 2},
		{1041, 5},
		{529, 4},
		{20, 1},
		{32, 0},
		{6, 0},
		{328, 2},
		{51712, 2},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%d", test.input), func(t *testing.T) {
			actual := findLongestBinaryGap(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
