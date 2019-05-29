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

func findOddOccurence(as []int) int {
	cnts := make(map[int]int)
	for _, a := range as {
		cnts[a]++
	}

	var unpair int
	for a, cnt := range cnts {
		if cnt%2 != 0 {
			unpair = a
			break
		}
	}

	return unpair
}
