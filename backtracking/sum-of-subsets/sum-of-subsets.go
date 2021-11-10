package sum_of_subsets

var result [][]int

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

func backtrack(in []int, from int, needed, curSum int, current []int) {
	if needed == curSum {
		result = append(result, current)
		return
	}

	for i := from; i < len(in); i++ {
		newVal := in[i]
		newSum := newVal + curSum
		if newSum > needed {
			continue
		}
		tmp := make([]int, len(current))
		copy(tmp, current)
		tmp = append(tmp, newVal)
		backtrack(in, i+1, needed, newSum, tmp)
	}
}
