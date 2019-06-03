package maxslice

func findMaximumProfit(as []int) int {
	var min, max int
	for i, a := range as {
		if i == 0 || a < min {
			min = a
		}

		current := a - min
		if max < current {
			max = current
		}
	}

	return max
}
