package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}
	root := &ListNode{Val: list[0]}
	prev := root
	for i := 1; i < len(list); i++ {
		node := &ListNode{Val: list[i]}
		prev.Next = node
		prev = node
	}
	return root
}
