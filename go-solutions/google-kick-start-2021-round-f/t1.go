package google_kick_start_2021_round_f

func minDistanceToTrashBin(mask string) int {
	sums := make([]int, len(mask))
	curSum := -99999999
	for i := 0; i < len(mask); i++ {
		if mask[i] == '1' {
			curSum = 0
			sums[i] = curSum
		} else {
			curSum++
			sums[i] = curSum
		}
	}
	result := 0
	curSum = 99999999
	for i := len(mask) - 1; i >= 0; i-- {
		if mask[i] == '1' {
			curSum = 0
		} else {
			curSum++
		}

		prevSum := sums[i]
		if prevSum < curSum {
			result += prevSum
		} else {
			result += curSum
		}
	}

	return result
}
