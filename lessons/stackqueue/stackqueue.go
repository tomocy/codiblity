package stackqueue

func isProperlyNested(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var expecteds []string
	for _, r := range s {
		var right string
		if right, ok := parens[string(r)]; ok {
			expecteds = append(expecteds, right)
			continue
		}

		if len(expecteds) == 0 {
			return false
		}
		right = string(r)
		expected := expecteds[len(expecteds)-1]
		expecteds = expecteds[:len(expecteds)-1]
		if right != expected {
			return false
		}
	}

	return len(expecteds) == 0
}

var parens = map[string]string{
	"[": "]",
	"{": "}",
	"(": ")",
}
