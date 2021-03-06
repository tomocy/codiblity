package sum

import (
	"fmt"
	"testing"
)

func TestCountPassingCars(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{0},
			0,
		},
		{
			[]int{1},
			0,
		},
		{
			[]int{0, 1, 0, 1, 1},
			5,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countPassingCars(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}

func TestFindMinimalImpactFactors(t *testing.T) {
	type input struct {
		s  string
		ps []int
		qs []int
	}
	tests := []struct {
		input    input
		expected []int
	}{
		{
			input{
				"CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6},
			},
			[]int{2, 4, 1},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := findMinimalImpactFactors(test.input.s, test.input.ps, test.input.qs)
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

func TestDetectBeginOfMinimalAverageOfSlice(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{4, 2, 2, 5, 1, 5, 8},
			1,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := detectBeginOfMinimalAverageOfSlice(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}

func TestCountDivisibles(t *testing.T) {
	type input struct {
		a, b, k int
	}
	tests := []struct {
		input    input
		expected int
	}{
		{
			input{
				6, 11, 2,
			},
			3,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countDivisibles(test.input.a, test.input.b, test.input.k)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
