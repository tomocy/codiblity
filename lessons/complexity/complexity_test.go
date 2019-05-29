package complexity

import (
	"fmt"
	"testing"
)

func TestCountFrogJumps(t *testing.T) {
	type input struct {
		x, y, d int
	}
	tests := []struct {
		input    input
		expected int
	}{
		{
			input{1, 2, 1},
			1,
		},
		{
			input{1, 2, 10},
			1,
		},
		{
			input{10, 85, 30},
			3,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countFrogJumps(test.input.x, test.input.y, test.input.d)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}

func TestFindMissingElement(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{2, 3, 1, 5},
			4,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findMissingElement(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}

func TestFindMinimalDifferenceOfSplitedTapes(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{3, 1},
			2,
		},
		{
			[]int{-1000, 1000},
			2000,
		},
		{
			[]int{3, 1, 2, 4, 3},
			1,
		},
		{
			[]int{3, -1, 2, -4, 3},
			1,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findMinimalDifferenceOfSplitedTapes(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
