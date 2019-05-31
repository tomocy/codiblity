package sort

import (
	"fmt"
	"testing"
)

func TestCountDistincts(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{2, 1, 1, 2, 3, 1},
			3,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countDistincts(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}

func TestFindMaximalProduct(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{-3, 1, 2, -2, 5, 6},
			60,
		},
		{
			[]int{-5, -6, -4, -7, -10},
			-120,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findMaximalProduct(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %v\n", actual, test.expected)
			}
		})
	}
}
