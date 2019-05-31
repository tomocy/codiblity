package stackqueue

import (
	"fmt"
	"testing"
)

func TestIsProperlyNested(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"{[()()]}", true},
		{"([)()]", false},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := isProperlyNested(test.input)
			if actual != test.expected {
				t.Errorf("got %t, but expected %t\n", actual, test.expected)
			}
		})
	}
}

func TestCountAlivingFishes(t *testing.T) {
	tests := []struct {
		input    []*fish
		expected int
	}{
		{
			[]*fish{
				&fish{0, 4, upstream},
				&fish{1, 3, downstream},
				&fish{2, 2, upstream},
				&fish{3, 1, upstream},
				&fish{4, 5, upstream},
			},
			2,
		},
		{
			[]*fish{
				&fish{0, 4, upstream},
				&fish{1, 3, downstream},
				&fish{2, 2, downstream},
				&fish{3, 1, upstream},
				&fish{4, 5, upstream},
			},
			2,
		},
		{
			[]*fish{
				&fish{0, 4, upstream},
				&fish{1, 3, upstream},
				&fish{2, 2, upstream},
			},
			3,
		},
		{
			[]*fish{
				&fish{0, 4, downstream},
				&fish{1, 3, downstream},
				&fish{2, 2, downstream},
			},
			3,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countFishesAlive(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}

func TestCountBlocksForStoneWall(t *testing.T) {
	tests := []struct {
		input    []*block
		expected int
	}{
		{
			[]*block{
				&block{8},
				&block{8},
				&block{5},
				&block{7},
				&block{9},
				&block{8},
				&block{7},
				&block{4},
				&block{8},
			},
			7,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countBlocksForStoneWall(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
