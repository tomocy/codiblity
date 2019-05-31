package sort

import (
	"sort"
)

func countDistincts(as []int) int {
	dists := make(map[int]int)
	for _, a := range as {
		dists[a]++
	}

	return len(dists)
}

func findMaximalProduct(as []int) int {
	sorted := sortByAsc(as)
	smallest, smallest2 := sorted[0], sorted[1]
	largest, largest2, largest3 := sorted[len(sorted)-1], sorted[len(sorted)-2], sorted[len(sorted)-3]
	if 0 <= largest && largest2*largest3 < smallest*smallest2 {
		return smallest * smallest2 * largest
	}

	return largest3 * largest2 * largest
}

func canBuildTrianble(as []int) bool {
	if len(as) < 3 {
		return false
	}

	sorted := sortByAsc(as)
	for i := 0; i <= len(sorted)-3; i++ {
		if sorted[i+2] <= 0 {
			continue
		}
		if sorted[i+2] < sorted[i]+sorted[i+1] {
			return true
		}
	}

	return false
}

func sortByAsc(as []int) []int {
	sorted := make([]int, len(as))
	copy(sorted, as)
	sort.Ints(sorted)

	return sorted
}

func countDiscIntersects(discs []*disc) int {
	var cnt int
	sorted := sortByDiameter(discs)
	for i, disc := range sorted {
		if len(sorted)-1 <= i {
			break
		}
		for _, target := range sorted[i+1:] {
			if disc.doIntersect(target) {
				cnt++
			}
		}
	}

	return cnt
}

func sortByDiameter(src []*disc) []*disc {
	sorted := make(byDiameter, len(src))
	copy(sorted, src)
	sort.Sort(sorted)

	return sorted
}

type byDiameter []*disc

func (d byDiameter) Len() int {
	return len(d)
}

func (d byDiameter) Less(i, j int) bool {
	return d[i].diameter() < d[j].diameter()
}

func (d byDiameter) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

type disc struct {
	left, right point
}

type point struct {
	x int
}

func (d *disc) diameter() int {
	return d.right.x - d.left.x
}

func (d *disc) doIntersect(target *disc) bool {
	if target.right.x < d.left.x || d.right.x < target.left.x {
		return false
	}

	return true
}
