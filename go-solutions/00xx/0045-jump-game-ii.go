package _0xx

// 0045. Jump Game II
// 2, 3, 1, 1, 2, 1, 1
// 0, 1, 2, 3, 4, 5, 6
func jump(nums []int) int {
	var max = func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	length := len(nums) - 1
	curr := -1
	next := 0
	ans := 0
	for i := 0; next < length; i++ {
		if i > curr {
			ans++
			curr = next
		}
		v := nums[i] + i
		next = max(next, v)
	}
	return ans
}
