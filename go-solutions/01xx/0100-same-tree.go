package _1xx

// 0100. Same Tree
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var popQueue = func(queue []*TreeNode) (*TreeNode, []*TreeNode) {
		if len(queue) == 0 {
			return nil, nil
		}
		res := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		return res, queue
	}
	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		p, queue = popQueue(queue)
		q, queue = popQueue(queue)
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left)
		queue = append(queue, q.Left)
		queue = append(queue, p.Right)
		queue = append(queue, q.Right)
	}
	return true
}

// 0100. Same Tree (Recursion)
func isSameTree_Rec(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree_Rec(p.Left, q.Left) && isSameTree_Rec(p.Right, q.Right)
}
