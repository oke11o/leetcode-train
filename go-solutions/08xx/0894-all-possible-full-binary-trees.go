package _8xx

/**
 * https://leetcode.com/problems/all-possible-full-binary-trees/
 * 0894. All Possible Full Binary Trees
 * Medium
 * Tags: Dynamic Programming, Tree, Recursion, Memoization, Binary Tree
 *
 * Help video: https://www.youtube.com/watch?v=Ci-82MggDYI
 */
func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}
	result := make([]*TreeNode, 0)
	if n == 1 {
		result = append(result, &TreeNode{})
		return result
	}
	for leftCnt := 1; leftCnt < n; leftCnt += 2 {
		rightCnt := n - 1 - leftCnt
		lefts := allPossibleFBT(leftCnt)
		rights := allPossibleFBT(rightCnt)

		for _, l := range lefts {
			for _, r := range rights {
				node := &TreeNode{Left: l, Right: r}
				result = append(result, node)
			}
		}
	}

	return result
}

var allPossibleFBT_cache = map[int][]*TreeNode{}

func allPossibleFBT_withCache(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}
	if v, ok := allPossibleFBT_cache[n]; ok {
		return v
	}
	result := make([]*TreeNode, 0)
	if n == 1 {
		result = append(result, &TreeNode{})
		allPossibleFBT_cache[n] = result
		return result
	}

	for leftCnt := 1; leftCnt < n; leftCnt += 2 {
		rightCnt := n - 1 - leftCnt
		lefts := allPossibleFBT(leftCnt)
		rights := allPossibleFBT(rightCnt)

		for _, l := range lefts {
			for _, r := range rights {
				node := &TreeNode{Left: l, Right: r}
				result = append(result, node)
			}
		}
	}
	allPossibleFBT_cache[n] = result

	return result
}
