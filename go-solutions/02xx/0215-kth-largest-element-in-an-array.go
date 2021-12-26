package _2xx

/**
 * 0215. Kth Largest Element in an Array
 * Medium
 * #array, #devide_and_conquer, #sorting, #head_priority_queue, #quickselect
 * https://leetcode.com/problems/kth-largest-element-in-an-array/
 *
 * Reference: - 0414. Third Maximum Number
 */
func findKthLargest(nums []int, k int) int {
	result := make([]int, k)
	maxMin := -1 << (64 - 1)
	for i := 0; i < k; i++ {
		result[i] = maxMin
	}
	for _, n := range nums {
		for i := 0; i < k; i++ {
			// Тут проверяем, есть ли такой элемент в нашем максимуме. Но для данного условия это не нежно
			//if result[i] == n {
			//	break
			//}
			if result[i] < n {
				for j := k - 1; j > i; j-- {
					result[j] = result[j-1]
				}
				result[i] = n
				break
			}
		}
	}
	return result[k-1]
}
