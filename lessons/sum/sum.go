package sum

func countPassingCars(as []int) int {
	var cnt int
	for i, a := range as {
		if i == len(as)-1 {
			break
		}
		if !isCarTravelingToEast(a) {
			continue
		}

		rests := as[i+1:]
		cnt += countCarsTravelingToWest(rests)
	}

	return cnt
}

func countCarsTravelingToWest(as []int) int {
	var cnt int
	for _, a := range as {
		if isCarTravelingToWest(a) {
			cnt++
		}
	}

	return cnt
}

func isCarTravelingToWest(n int) bool {
	return !isCarTravelingToEast(n)
}

func isCarTravelingToEast(n int) bool {
	return n == 0
}
