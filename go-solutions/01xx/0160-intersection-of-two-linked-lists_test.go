package _1xx

/*
https://leetcode.com/problems/intersection-of-two-linked-lists/
160. Intersection of Two Linked Lists
Easy
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var calcLen = func(node *ListNode) int {
		var ans int
		for node != nil {
			ans++
			node = node.Next
		}
		return ans
	}
	aLen := calcLen(headA)
	bLen := calcLen(headB)
	if aLen < bLen {
		headA, headB = headB, headA
		aLen, bLen = bLen, aLen
	}
	i := aLen - bLen
	for i > 0 {
		headA = headA.Next
		i--
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}
