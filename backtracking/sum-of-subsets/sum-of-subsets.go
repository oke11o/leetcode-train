package sum_of_subsets

var result [][]int

/**
 * https://youtu.be/kyLxTdsT8ws
 */
func sumOfSubsets(in []int, sum int) [][]int {
	totalSum := 0
	for _, v := range in {
		totalSum += v
	}
	if totalSum < sum {
		return nil
	}

	result = make([][]int, 0)
	backtrack(in, 0, sum, 0, []int{})
	return result
}

func backtrack(in []int, from int, needSum, prevSum int, current []int) {
	if needSum == prevSum {
		result = append(result, current)
		return
	}

	for i := from; i < len(in); i++ {
		if boundFunc(in[i], needSum, prevSum) {
			tmp := make([]int, len(current))
			copy(tmp, current)
			tmp = append(tmp, in[i])
			backtrack(in, i+1, needSum, prevSum+in[i], tmp)
		}
	}
}

func boundFunc(val, needSum, prevSum int) bool {
	if val+prevSum > needSum {
		return false
	}
	return true
}
