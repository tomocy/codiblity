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
