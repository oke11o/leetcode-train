package _6xx

// 0695. Max Area of Island (DFS)
// Medium
// Array, Matrix, BFS,DFS
// https://leetcode.com/problems/max-area-of-island/
func maxAreaOfIsland(grid [][]int) int {
	result := newMaxAreaOfIslandVisited(len(grid), len(grid[0]))
	return result.run(grid)
}

func newMaxAreaOfIslandVisited(rowLen, colLen int) *maxAreaOfIslandResult {
	result := &maxAreaOfIslandResult{
		visited: make([][]bool, rowLen),
	}
	for i := 0; i < rowLen; i++ {
		result.visited[i] = make([]bool, colLen)
	}
	return result
}

type maxAreaOfIslandResult struct {
	visited [][]bool
}

func (r *maxAreaOfIslandResult) run(grid [][]int) int {
	max := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !r.visited[i][j] {
				cnt := r.dfs(grid, i, j)
				if max < cnt {
					max = cnt
				}
			}
		}
	}
	return max
}

func (r *maxAreaOfIslandResult) dfs(grid [][]int, row int, col int) int {
	if r.visited[row][col] {
		return 0
	}
	r.visited[row][col] = true
	if grid[row][col] == 0 {
		return 0
	}
	cnt := 1
	//for _, n := range [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} { //neighbours
	//	row := i + n[0]
	//	if row < 0 || row >= len(grid) {
	//		continue
	//	}
	//	col := j + n[1]
	//	if col < 0 || col >= len(grid[0]) {
	//		continue
	//	}
	//	cnt += r.dfs(grid, row, col)
	//}
	if row >= 1 {
		cnt += r.dfs(grid, row-1, col)
	}
	if row < len(grid)-1 {
		cnt += r.dfs(grid, row+1, col)
	}
	if col >= 1 {
		cnt += r.dfs(grid, row, col-1)
	}
	if col < len(grid[0])-1 {
		cnt += r.dfs(grid, row, col+1)
	}
	return cnt
}
