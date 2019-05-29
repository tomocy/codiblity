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

func TestCountAsOperations(t *testing.T) {
	type input struct {
		x  int
		as []int
	}
	tests := []struct {
		input    input
		expected []int
	}{
		{
			input{
				5,
				[]int{3, 4, 4, 6, 1, 4, 4},
			},
			[]int{3, 2, 2, 4, 2},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countAsOperations(test.input.x, test.input.as)
			if len(actual) != len(test.expected) {
				t.Fatalf("len: got %d, but expected %d\n", len(actual), len(test.expected))
			}
			for i := range test.expected {
				if actual[i] != test.expected[i] {
					t.Errorf("at %d: got %d, but expected %d\n", i, actual[i], test.expected[i])
				}
			}
		})
	}
}

func TestFindSmallestPositiveInteger(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1, 3, 6, 4, 1, 2},
			5,
		},
		{
			[]int{1, 2, 3},
			4,
		},
		{
			[]int{-1, -3},
			1,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findSmallestPositiveInteger(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
