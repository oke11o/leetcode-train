package go_solutions

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		idx := left + (right-left)/2
		if nums[idx] == target {
			return idx
		} else if nums[idx] < target {
			left = idx + 1
		} else {
			right = idx - 1
		}
	}
	return -1
}
