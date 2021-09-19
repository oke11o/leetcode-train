package go_solutions

func search(nums []int, target int) int {
	left := -1
	right := len(nums)
	for (right - left) > 1 {
		idx := (right + left) / 2
		if nums[idx] <= target {
			left = idx
		} else {
			right = idx
		}
	}
	if left != -1 && nums[left] == target {
		return left
	}
	return -1
}
