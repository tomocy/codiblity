package maxslice

import "math"

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

func findMaximumSum(as []int) int {
	var maxAs, max int
	for i, a := range as {
		maxAs += a
		if maxAs < a {
			maxAs = a
		}

		if i == 0 || max < maxAs {
			max = maxAs
		}
	}

	return max
}

func findMaximumSumOfDouble(as []int) int {
	if len(as) <= 3 {
		return 0
	}

	n := len(as) - 2
	lefts, rights := make([]int, n), make([]int, n)
	for i := 0; i < n-1; i++ {
		left, right := as[i+1], as[n-i]
		front, back := i+1, n-2-i
		lefts[front] = int(math.Max(0, float64(left+lefts[front-1])))
		rights[back] = int(math.Max(0, float64(right+rights[back+1])))
	}

	max := lefts[0] + rights[0]
	for i := 1; i < len(lefts); i++ {
		max = int(math.Max(float64(max), float64(lefts[i]+rights[i])))
	}

	return max
}
