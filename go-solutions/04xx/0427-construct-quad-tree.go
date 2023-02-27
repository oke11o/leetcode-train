package _4xx

/*
427. Construct Quad Tree
Medium
#array, #divide_and_conquer, #tree, #matrix
*/
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	return solve(grid, 0, 0, len(grid))
}

func sameValue(grid [][]int, x1, y1 int, length int) bool {
	for i := x1; i < x1+length; i++ {
		for j := y1; j < y1+length; j++ {
			if grid[i][j] != grid[x1][y1] {
				return false
			}
		}
	}
	return true
}

func solve(grid [][]int, x1, y1 int, length int) *Node {
	if sameValue(grid, x1, y1, length) {
		return &Node{Val: grid[x1][y1] == 1, IsLeaf: true}
	}
	root := &Node{}
	l := length / 2
	root.TopLeft = solve(grid, x1, y1, l)
	root.TopRight = solve(grid, x1, y1+l, l)
	root.BottomLeft = solve(grid, x1+l, y1, l)
	root.BottomRight = solve(grid, x1+l, y1+l, l)
	return root
}
