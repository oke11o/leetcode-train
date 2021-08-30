package missing_number

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}

	return len(nums)
}

func missingNumber2(nums []int) int {
	tmp := make([]*int, len(nums)+1)

	var in int
	for _, i := range nums {
		tmp[i] = &in
	}
	for idx, i := range tmp {
		if i == nil {
			return idx
		}
	}

	return len(nums)
}
