package leader

import "errors"

func countEquiLeader(as []int) int {
	leader, leaderCnt, err := findLeaderValueAndCount(as)
	if err != nil {
		return 0
	}

	var leftLen, leftCnt, cnt int
	for _, a := range as {
		leftLen++
		if a == as[leader] {
			leftCnt++
		}

		if leftLen/2 < leftCnt {
			rightLen, rightCnt := len(as)-leftLen, leaderCnt-leftCnt
			if rightLen/2 < rightCnt {
				cnt++
			}
		}
	}

	return cnt
}

func findLeader(as []int) (int, error) {
	leader, _, err := findLeaderAndCount(as)
	return leader, err
}

func findLeaderAndCount(as []int) (int, int, error) {
	val, cnt, err := findLeaderValueAndCount(as)
	if err != nil {
		return 0, 0, err
	}

	for i, a := range as {
		if a == val {
			return i, cnt, nil
		}
	}

	return 0, 0, errors.New("dev: no leader and count")
}

func findLeaderValueAndCount(as []int) (int, int, error) {
	cnts := mapCandidateCounts(as)
	for cand, cnt := range cnts {
		if len(as)/2 < cnt {
			return cand, cnt, nil
		}
	}

	return 0, 0, errors.New("no leader")
}

func mapCandidateCounts(as []int) map[int]int {
	cnts := make(map[int]int)
	for _, a := range as {
		cnts[a]++
	}

	return cnts
}
