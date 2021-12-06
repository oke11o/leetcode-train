package _1xx

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 0116. Populating Next Right Pointers in Each Node
// Medium
// Tree, DFS, BFS, Binary Tree
// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/
// https://www.youtube.com/watch?v=t2XlsZ1PEy4&ab_channel=CleanCoder
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	queue = append(queue, nil)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			if len(queue) > 0 {
				node.Next = queue[0]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		} else {
			if len(queue) == 0 {
				break
			}
			queue = append(queue, nil)
		}
	}
	return root
}
