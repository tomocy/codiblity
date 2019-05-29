package iteration

import (
	"strconv"
	"strings"
)

func findLongestBinaryGap(n int) int {
	gaps := findBinaryGaps(n)
	var max int
	for _, gap := range gaps {
		current := len(gap)
		if max < current {
			max = current
		}
	}

	return max
}

func findBinaryGaps(n int) []string {
	bin := strconv.FormatInt(int64(n), 2)
	trimed := trimTrailingZeros(bin)
	ss := strings.Split(trimed, "1")
	var gaps []string
	for _, s := range ss {
		if s != "" {
			gaps = append(gaps, s)
		}
	}

	return gaps
}

func trimTrailingZeros(s string) string {
	return strings.TrimRight(s, "0")
}
