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

func findMinimalImpactFactors(s string, ps []int, qs []int) []int {
	factors := make([]int, len(ps))
	for i := range ps {
		begin, end := ps[i], qs[i]
		part := s[begin : end+1]
		categ := categorizeNucleotides(part)
		factors[i] = findMinimalImpactFactor(categ)
	}

	return factors
}

func categorizeNucleotides(s string) category {
	category := make(category)
	for _, r := range s {
		category[string(r)]++
	}

	return category
}

func findMinimalImpactFactor(categ category) int {
	var i, min int
	for nuc := range categ {
		prevI := i
		i++
		current := factors[nuc]
		if prevI == 0 {
			min = current
			continue
		}

		if current < min {
			min = current
		}
	}

	return min
}

var factors = map[string]int{
	"A": 1,
	"C": 2,
	"G": 3,
	"T": 4,
}

type category map[string]int

func detectBeginOfMinimalAverageOfSlice(as []int) int {
	var minIdx int
	var min float32
	for i := range as {
		maxIdx := len(as) - minLen
		if maxIdx < i {
			break
		}

		current := findMinimalAverageOfSlice(i, as)
		if i == 0 {
			minIdx, min = i, current
			continue
		}

		if current < min {
			minIdx, min = i, current
		}
	}

	return minIdx
}

func findMinimalAverageOfSlice(begin int, as []int) float32 {
	var min float32
	for i := range as[begin:] {
		end := begin + minLen + i
		if len(as) < end {
			break
		}

		current := calculateAverage(as[begin:end])
		if i == 0 {
			min = current
			continue
		}

		if current < min {
			min = current
		}
	}

	return min
}

func calculateAverage(as []int) float32 {
	var sum float32
	for _, a := range as {
		sum += float32(a)
	}

	return sum / float32(len(as))
}

var minLen = 2
