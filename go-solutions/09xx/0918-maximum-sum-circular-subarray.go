package _9xx

// 0918. Maximum Sum Circular Subarray
func maxSubarraySumCircular(nums []int) int {

	// Case 1: get the maximum sum using standard kadane' s algorithm
	maxKadane := kadane(nums)

	// Case 2: Now find the maximum sum that includes corner elements.
	maxWrap := 0
	for i := 0; i < len(nums); i++ {
		maxWrap += nums[i]
		nums[i] = 0 - nums[i]
	}

	// max sum with corner elements will be:
	// array-sum - (-max subarray sum of inverted array)
	maxWrap = maxWrap + kadane(nums)
	if maxWrap == 0 {
		return maxKadane
	}
	if maxWrap > maxKadane {
		return maxWrap
	}
	return maxKadane
}

func kadane(nums []int) int {
	result := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		newSum := curSum + nums[i]
		if newSum > nums[i] {
			curSum = newSum
		} else {
			curSum = nums[i]
		}
		if result < curSum {
			result = curSum
		}
	}
	return result
}
