package _9xx

func sortArrayByParityII(nums []int) []int {
	length := len(nums)
	var odd, even int
	odd = 1
	for length > odd && length > even {
		if nums[odd]%2 == 0 && nums[even]%2 == 1 {
			nums[even], nums[odd] = nums[odd], nums[even]
		}
		if even < length && nums[even]%2 == 0 {
			even += 2
		}
		if odd < length && nums[odd]%2 == 1 {
			odd += 2
		}
	}
	return nums
}
