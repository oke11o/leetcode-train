package _1xx

// 1110. Delete Nodes And Return Forest
func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	result := []*TreeNode{}
	mapa := make(map[int]struct{}, len(toDelete))
	for _, v := range toDelete {
		mapa[v] = struct{}{}
	}
	queue := []*TreeNode{root}
	result = append(result, root)

	// Эти две мапы нужны для удаления нод из результата
	inResutValues := map[int]int{root.Val: 0}
	toSkipIdxs := map[int]struct{}{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if _, ok := mapa[node.Val]; ok {
			// так как этого ребенка добавили в результат чуть ниже, то тут делаем проверку, что может его все таки стоит удалить?
			if idx, ok2 := inResutValues[node.Val]; ok2 {
				toSkipIdxs[idx] = struct{}{}
			}
			if node.Left != nil {
				result = append(result, node.Left)
				inResutValues[node.Left.Val] = len(result) - 1
			}
			if node.Right != nil {
				result = append(result, node.Right)
				inResutValues[node.Right.Val] = len(result) - 1
			}
		}

		// тут мы добавляем детей в обход и удаляем из детей текущей
		// в результат детей добавляем выше
		if node.Left != nil {
			queue = append(queue, node.Left)
			if _, ok := mapa[node.Left.Val]; ok {
				node.Left = nil
			}
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			if _, ok := mapa[node.Right.Val]; ok {
				node.Right = nil
			}
		}
	}
	if len(toSkipIdxs) == 0 {
		return result
	}
	res := make([]*TreeNode, 0, len(result)-len(toSkipIdxs))
	for i, v := range result {
		if _, ok := toSkipIdxs[i]; !ok {
			res = append(res, v)
		}
	}

	return res
}
