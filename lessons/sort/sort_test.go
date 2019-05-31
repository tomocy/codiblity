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

func TestCanBuildTrianble(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			[]int{10, 2, 5, 1, 8, 20},
			true,
		},
		{
			[]int{10, 50, 5, 1},
			false,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := canBuildTrianble(test.input)
			if actual != test.expected {
				t.Errorf("got %t, but expected %t\n", actual, test.expected)
			}
		})
	}
}

func TestCountDiscIntersects(t *testing.T) {
	tests := []struct {
		input    []*disc
		expected int
	}{
		{
			[]*disc{
				&disc{point{-1}, point{1}},
				&disc{point{-4}, point{6}},
				&disc{point{0}, point{4}},
				&disc{point{2}, point{4}},
				&disc{point{0}, point{8}},
				&disc{point{5}, point{5}},
			},
			11,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countDiscIntersects(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
