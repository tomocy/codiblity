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
