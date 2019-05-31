package leader

import (
	"fmt"
	"testing"
)

func TestFindAnyLeaderIndex(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{3, 4, 3, 2, 3, -1, 3, 3},
			[]int{0, 2, 4, 6, 7},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual, err := findLeader(test.input)
			if len(test.expected) == 0 && err == nil {
				t.Errorf("got %d, but expected to be not found\n", actual)
				return
			}
			for _, expected := range test.expected {
				if actual == expected {
					return
				}
			}
			t.Errorf("got %d, but expected any of %v\n", actual, test.expected)
		})
	}
}

func TestCountEquiLeader(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{4, 3, 4, 4, 4, 2},
			2,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.input), func(t *testing.T) {
			actual := countEquiLeader(test.input)
			if actual != test.expected {
				t.Errorf("got %d, but expected %d\n", actual, test.expected)
			}
		})
	}
}
