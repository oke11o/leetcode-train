package go_solutions

type NumArray struct {
	sum []int
}

func Constructor(nums []int) NumArray {
	result := NumArray{
		sum: make([]int, len(nums)),
	}
	result.sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		result.sum[i] = result.sum[i-1] + nums[i]
	}
	return result
}

func (a *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return a.sum[right]
	}
	return a.sum[right] - a.sum[left-1]
}
