package _2xx

// 0283. Move Zeroes
// Medium
// https://leetcode.com/problems/move-zeroes/
func moveZeroes(nums []int) {
	shift := 0
	for i := 0; i < len(nums); i++ {
		nums[i], nums[i-shift] = nums[i-shift], nums[i]
		if nums[i-shift] == 0 {
			shift++
		}
	}
}
