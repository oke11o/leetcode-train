package _1xx

func singleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result = result ^ n
	}
	return result
}
