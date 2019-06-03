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

func TestFindMaximumSum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{1, 3},
			4,
		},
		{
			[]int{3, 2, -6, 4, 0},
			5,
		},
		{
			[]int{10},
			10,
		},
		{
			[]int{-10},
			-10,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findMaximumSum(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}

func TestFindMaximumSumOfDouble(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{3, 2, 6, -1, 4, 5, -1, 2},
			17,
		},
		{
			[]int{0, 10, -5, -2, 0},
			10,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findMaximumSumOfDouble(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
