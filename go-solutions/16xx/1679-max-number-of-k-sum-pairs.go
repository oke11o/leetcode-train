package _6xx

/**
https://leetcode.com/problems/max-number-of-k-sum-pairs/
1679. Max Number of K-Sum Pairs
Medium
*/
func maxOperations(nums []int, k int) int {
	result := 0
	hm := make(map[int]int)
	for _, v := range nums {
		_, ok := hm[k-v]
		if ok {
			result++
			hm[k-v]--
			if hm[k-v] == 0 {
				delete(hm, k-v)
			}
		} else {
			hm[v]++
		}
	}
	return result
}
