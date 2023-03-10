package _7xx

/*
https://leetcode.com/problems/find-pivot-index/
724. Find Pivot Index
Easy
*/

// [1,7,3,6,5,6]
// [28,27,20,17,11,6]
// [1,8,11,17,]
func pivotIndex_2(nums []int) int {
	rightSums := make([]int, len(nums)+1)
	for i := len(nums) - 1; i >= 0; i-- {
		rightSums[i] = rightSums[i+1] + nums[i]
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		if sum == rightSums[i+1] {
			return i
		}
		sum += nums[i]
	}

	return -1
}

func pivotIndex(nums []int) int {
	rightSum := 0
	for i := 0; i < len(nums); i++ {
		rightSum += nums[i]
	}
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if leftSum == rightSum {
			return i
		}
		leftSum += nums[i]
		rightSum -= nums[i]
	}

	return -1
}
