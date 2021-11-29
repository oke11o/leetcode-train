package _1xx

// 0189. Rotate Array
// Medium
// Array, Math, Two Pointers
// https://leetcode.com/problems/rotate-array/
func rotate(nums []int, k int) {
	N := len(nums)
	k = k % N
	if k == 0 {
		return
	}
	slots := make([]int, k)
	copy(slots, nums[N-k:])
	kPointer := 0
	for i := 0; i < N; i++ {
		tmp := nums[i]
		nums[i] = slots[kPointer%k]
		slots[kPointer%k] = tmp
		kPointer++
	}
}
