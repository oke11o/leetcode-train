package _9xx

// 0918. Maximum Sum Circular Subarray
func maxSubarraySumCircular(nums []int) int {
	length := len(nums)

	tmpResult := -9999999
	minIdx := 0
	min := 99999999999999
	dpMin := make([]int, length)
	dpMin[0] = nums[0]
	for i := 1; i < length; i++ {
		newSum := nums[i] + dpMin[i-1]
		if newSum > nums[i] {
			dpMin[i] = newSum
		} else {
			dpMin[i] = nums[i]
		}
		if min > dpMin[i] {
			min = dpMin[i]
			minIdx = i
		}
		if tmpResult < dpMin[i] {
			tmpResult = dpMin[i]
		}
	}

	result := -99999999
	dp := make([]int, length)
	for i := minIdx + 1; i < length+minIdx+1; i++ {
		curIdx := i % length
		prevIdx := (i - 1) % length
		num := nums[curIdx]

		newSum := dp[prevIdx] + num
		if newSum > num {
			dp[curIdx] = newSum
		} else {
			dp[curIdx] = num
		}

		if result < dp[curIdx] {
			result = dp[curIdx]
		}

	}

	if tmpResult > result {
		return tmpResult
	}

	return result
}
