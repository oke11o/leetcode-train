package amazon

// 2. Determine if the sum of two integers is equal to the given value
// https://www.educative.io/blog/crack-amazon-coding-interview-questions
// Amazon
func findSumOfTwo(list []int, sum int) bool {
	m := make(map[int]struct{}, len(list))
	for _, v := range list {
		if _, ok := m[sum-v]; ok {
			return true
		}
		m[v] = struct{}{}
	}

	return false
}
