package array

func rotateCyclicly(as []int, k int) []int {
	if len(as) == 0 {
		return as
	}

	n := k % len(as)
	if n == 0 {
		return as
	}
	forward := as[:len(as)-n]
	back := as[len(as)-n:]

	return append(back, forward...)
}
