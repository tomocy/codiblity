package counting

import (
	"fmt"
	"testing"
)

func TestIsPermutation(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			[]int{4, 1, 3, 2},
			true,
		},
		{
			[]int{4, 1, 3},
			false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := isPermutation(test.input)
			if actual != test.expected {
				t.Errorf("got %t, but expected %t\n", actual, test.expected)
			}
		})
	}
}
