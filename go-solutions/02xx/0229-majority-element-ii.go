package _2xx

// 0229. Majority Element II
// Boyer Moore
func majorityElement(nums []int) []int {
	res1 := 0
	res2 := 1
	var vote1, vote2 int
	for _, num := range nums {
		if num == res1 {
			vote1++
		} else if num == res2 {
			vote2++
		} else if vote1 == 0 {
			res1 = num
			vote1++
		} else if vote2 == 0 {
			res2 = num
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	vote1 = 0
	vote2 = 0
	for _, num := range nums {
		if num == res1 {
			vote1 += 2
		} else {
			vote1--
		}
		if num == res2 {
			vote2 += 2
		} else {
			vote2--
		}
	}

	result := make([]int, 0)
	if vote1 > 0 {
		result = append(result, res1)
	}
	if vote2 > 0 {
		result = append(result, res2)
	}
	return result
}
