package _0xx

import "sort"

// 0078. Subsets
// Find All Subsets of a given set of integers
// Medium
// #amazon
// Backtracking, #backtracking, #with_leetcode_solution
// https://leetcode.com/problems/subsets/
// https://www.educative.io/m/find-all-subsets
// https://www.educative.io/blog/crack-amazon-coding-interview-questions
func subsets(nums []int) [][]int {
	var result [][]int
	result = append(result, []int{})
	for _, n := range nums {
		length := len(result)
		for i := 0; i < length; i++ {
			// По условию не важна сортировка. Но leetcode ее проверяет.
			// А если сортировать slice, то меняется и исходный result[i].
			// Поэтому у item устанавливаем capacity
			item := result[i][:len(result[i]):len(result[i])]
			item = append(item, n)
			sort.Ints(item)
			result = append(result, item)
		}
	}
	return result
}
