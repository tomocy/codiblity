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

func sortByAsc(as []int) []int {
	sorted := make([]int, len(as))
	copy(sorted, as)
	sort.Sort(sort.IntSlice(sorted))

	return sorted
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
