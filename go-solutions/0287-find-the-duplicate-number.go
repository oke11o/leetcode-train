package go_solutions

func findDuplicate(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}

	var (
		slow = nums[0]
		fast = nums[0]
	)
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
