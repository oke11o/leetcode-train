package _0xx

// 0055. Jump Game
func canJump(nums []int) bool {
	lastOKPosition := len(nums) - 1
	for i := lastOKPosition - 1; i >= 0; i-- {
		v := i + nums[i]
		if v >= lastOKPosition {
			lastOKPosition = i
		}
	}
	return lastOKPosition == 0
}
