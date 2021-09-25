package _7xx

// 0740. Delete and Earn
// 1 <= nums.length <= 2 * 10^4
// 1 <= nums[i] <= 10^4
func deleteAndEarn(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	bins := make([]int, 10001)
	maxIdx := 0
	for _, num := range nums {
		bins[num]++
		if maxIdx < num {
			maxIdx = num
		}
	}
	sums := make([]int, 10001)
	sums[0] = bins[0]
	sums[1] = bins[1]

	for i := 2; i <= maxIdx; i++ {
		currentSum := bins[i] * i
		sums[i] = currentSum + sums[i-2]
		if sums[i] < sums[i-1] {
			sums[i] = sums[i-1]
		}
	}

	return sums[maxIdx]
}
