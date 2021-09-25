package _1xx

// 0108. Convert Sorted Array to Binary Search Tree
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	node := &TreeNode{
		Val:   nums[middle],
		Left:  sortedArrayToBST(nums[:middle]),
		Right: sortedArrayToBST(nums[middle+1:]),
	}
	return node
}
