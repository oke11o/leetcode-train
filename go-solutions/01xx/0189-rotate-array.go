package _1xx

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// Moriss traversal O(n)
//

// Given an array, rotate the array to the right by k steps, where k is non-negative.
// 0189. Rotate Array
// Medium
// Array, Math, Two Pointers
// https://leetcode.com/problems/rotate-array/
//
// Step 1.
// - how big array is possible?
// - how big k is possible? May be k == 0? Or may be k is negative?

// First solution.
// 1. Keep last k elements to tmp storage.
// 2. Save to slots removing el (using tmp val)
// 3. Set el to nums[i] from slots
func rotate_with_mem(nums []int, k int) {
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

//Tricky swap
//[1, 2, 3, 4, 5] -> [1, 2, 1, 4, 5] -> [1, 2, 1, 4, 3] -> [1, 5, 1, 4, 3] -> [1, 5, 1, 2, 3] -> [4, 5, 1, 2, 3]
//[1, 2, 3, 4] -> [1, 2, 1, 4] -> [3, 2, 1, 4]; [3, 2, 1, 2] -> [3, 4, 1, 2]
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	if k == 0 {
		return
	}
	// start - для того, чтобы сдвинуться не 1 вперед, если мы попали на туже позицию с которой начали.
	// Это может быть, если у нас длинна nums и k имеют общий делитель.
	counter, start, prevPos := 0, 0, 0
	swappingVal := nums[prevPos]
	for counter < n {
		if prevPos == start && counter > 0 {
			prevPos++
			start = prevPos
			swappingVal = nums[prevPos]
		}
		newPos := (prevPos + k) % n
		tmp := nums[newPos]
		nums[newPos] = swappingVal
		swappingVal = tmp
		prevPos = newPos
		counter++
	}
}
