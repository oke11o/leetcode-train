package go_solutions

func totalFruit(fruits []int) int {
	want := 2
	right := 0
	left := 0
	result := 0
	distinct := make(map[int]int)
	for ; right < len(fruits); right++ {
		rightVal := fruits[right]
		if _, ok := distinct[rightVal]; !ok {
			distinct[rightVal] = 0
		}
		distinct[rightVal]++
		newResult := right - left + 1
		for len(distinct) > want {
			leftVal := fruits[left]
			distinct[leftVal]--
			if distinct[leftVal] == 0 {
				delete(distinct, leftVal)
			}
			left++
			newResult = right - left
		}
		if result < newResult {
			result = newResult
		}
	}
	return result
}
