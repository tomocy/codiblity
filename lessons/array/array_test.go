package array

import (
	"fmt"
	"testing"
)

func TestRotateCyclicly(t *testing.T) {
	type input struct {
		as []int
		n  int
	}
	tests := []struct {
		input    input
		expected []int
	}{
		{
			input{
				[]int{3, 8},
				0,
			},
			[]int{3, 8},
		},
		{
			input{
				[]int{3, 8},
				1,
			},
			[]int{8, 3},
		},
		{
			input{
				[]int{3, 8},
				2,
			},
			[]int{3, 8},
		},
		{
			input{
				[]int{3, 8},
				5,
			},
			[]int{8, 3},
		},
		{
			input{
				[]int{3, 8, 9, 7, 6},
				3,
			},
			[]int{9, 7, 6, 3, 8},
		},
		{
			input{
				[]int{0, 0, 0},
				3,
			},
			[]int{0, 0, 0},
		},
		{
			input{
				[]int{1, 2, 3, 4},
				4,
			},
			[]int{1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := rotateCyclicly(test.input.as, test.input.n)
			if len(actual) != len(test.expected) {
				t.Fatalf("length: got %d, but expected %d\n", len(actual), len(test.expected))
			}
			for i := range test.expected {
				if actual[i] != test.expected[i] {
					t.Errorf("at %d: got %d, but expected %d\n", i, actual[i], test.expected[i])
				}
			}
		})
	}
}

func TestFindOddOccurence(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1},
			1,
		},
		{
			[]int{9, 3, 9, 3, 9, 7, 9},
			7,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1},
			1,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findOddOccurence(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
