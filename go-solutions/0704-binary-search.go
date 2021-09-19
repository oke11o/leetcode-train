package go_solutions

func search(nums []int, target int) int {
	left := -1
	right := len(nums)
	for (right - left) > 1 {
		idx := (right + left) / 2
		if nums[idx] > target {
			right = idx
		} else {
			left = idx
		}
	}
	return left
}
