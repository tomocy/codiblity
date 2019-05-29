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

func TestTimeFrogJumpsOverRiver(t *testing.T) {
	type input struct {
		x  int
		as []int
	}
	type expected struct {
		cnt int
		err error
	}
	tests := []struct {
		input    input
		expected expected
	}{
		{
			input{
				5,
				[]int{1, 3, 1, 4, 2, 3, 5, 4},
			},
			expected{
				6,
				nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual, err := timeFrogJumpOverRiver(test.input.x, test.input.as)
			if err != test.expected.err {
				t.Errorf("err: got %v, but expected %v\n", err, test.expected.err)
			}
			if actual != test.expected.cnt {
				t.Errorf("cnt: got %d, but expected %d\n", actual, test.expected.cnt)
			}
		})
	}
}
