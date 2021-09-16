package go_solutions

func maxSubArray(nums []int) int {
	result := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]

		newSum := curSum + num
		if newSum > num {
			curSum = newSum
		} else {
			curSum = num
		}

		if result < curSum {
			result = curSum
		}
	}
	return result
}
