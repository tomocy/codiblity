package counting

import (
	"errors"
	"sort"
)

func isPermutation(as []int) bool {
	sorted := sortByAsc(as)
	var prev int
	for _, a := range sorted {
		if prev+1 != a {
			return false
		}

		prev = a
	}

	return true
}

func timeFrogJumpOverRiver(x int, as []int) (int, error) {
	felt := make(map[int]bool)
	for i := 1; i <= x; i++ {
		felt[i] = false
	}
	for i, a := range as {
		if x < a {
			continue
		}

		if !felt[a] {
			felt[a] = true
		}
		if isAbleToJumpOverRiver(felt) {
			return i, nil
		}
	}

	return 0, errors.New("could not jump over river")
}

func isAbleToJumpOverRiver(ready map[int]bool) bool {
	for _, v := range ready {
		if !v {
			return false
		}
	}

	return true
}

func countAsOperations(n int, as []int) []int {
	cnts := make([]int, n)
	for _, a := range as {
		if a == n+1 {
			adjustToMax(cnts)
			continue
		}

		cnts[a-1]++
	}

	return cnts
}

func adjustToMax(as []int) {
	max := max(as)
	for i := range as {
		as[i] = max
	}
}

func max(as []int) int {
	var max int
	for _, a := range as {
		if max < a {
			max = a
		}
	}

	return max
}

func findSmallestPositiveInteger(as []int) int {
	sorted := sortByAsc(as)
	var prev, smallest int
	for _, a := range sorted {
		if a <= prev {
			continue
		}

		if a != prev+1 {
			smallest = prev + 1
			break
		}

		prev = a
	}
	if smallest == 0 {
		smallest = sorted[len(sorted)-1] + 1
		if smallest <= 0 {
			smallest = 1
		}
	}

	return smallest
}

func sortByAsc(as []int) []int {
	sorted := make([]int, len(as))
	copy(sorted, as)
	sort.Sort(sort.IntSlice(sorted))

	return sorted
}
