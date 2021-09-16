package go_solutions

func productExceptSelf(nums []int) []int {
	length := len(nums)
	left := make([]int, length)
	left[0] = 1
	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right := make([]int, length)
	right[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = left[i] * right[i]
	}

	return result
}

func productExceptSelf2(nums []int) []int {
	length := len(nums)

	result := make([]int, length)
	result[0] = 1
	for i := 1; i < length; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	right := make([]int, length)
	right[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
		result[i] *= right[i]
	}

	return result
}
