package go_solutions

func findMinInRotatedSortedArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[0]
	}
	for {
		i := (left + right) / 2
		if i == left {
			return nums[i+1]
		}
		if nums[i] > nums[left] {
			left = i
		} else {
			right = i
		}
	}
}
