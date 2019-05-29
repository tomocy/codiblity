package complexity

import (
	"fmt"
	"testing"
)

func TestFrogJump(t *testing.T) {
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
