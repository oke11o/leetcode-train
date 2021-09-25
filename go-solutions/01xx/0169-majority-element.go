package _1xx

// 0169. Majority Element
func majorityElement(nums []int) int {
	result := 0
	vote := 0
	for _, v := range nums {
		if vote == 0 {
			result = v
			vote++
		} else if result == v {
			vote++
		} else {
			vote--
		}
	}
	return result
}
