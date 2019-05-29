package complexity

func countFrogJumps(x, y, d int) int {
	diff := y - x
	cnt := diff / d
	if diff%d != 0 {
		cnt++
	}

	return cnt
}
