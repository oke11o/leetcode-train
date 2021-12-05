package _6xx

func maxAreaOfIsland(grid [][]int) int {
	result := newMaxAreaOfIslandResult(grid)
	result.run()
	return result.max
}

func newMaxAreaOfIslandResult(grid [][]int) *maxAreaOfIslandResult {
	result := &maxAreaOfIslandResult{
		grid:       grid,
		visited:    make([][]bool, len(grid)),
		neighbours: [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		rowLen:     len(grid),
		colLen:     len(grid[0]),
	}
	for i := 0; i < result.rowLen; i++ {
		result.visited[i] = make([]bool, result.colLen)
	}
	return result
}

type maxAreaOfIslandResult struct {
	grid   [][]int
	max    int
	tmpCnt int

	neighbours [4][2]int
	rowLen     int
	colLen     int
	visited    [][]bool
}

func (r *maxAreaOfIslandResult) run() {
	for i := 0; i < r.rowLen; i++ {
		for j := 0; j < r.colLen; j++ {
			if !r.visited[i][j] {
				r.tmpCnt = 0
				r.dfs(i, j)
				if r.max < r.tmpCnt {
					r.max = r.tmpCnt
				}
			}
		}
	}
}

func (r *maxAreaOfIslandResult) dfs(i int, j int) {
	if r.visited[i][j] {
		return
	}
	r.visited[i][j] = true
	if r.grid[i][j] == 0 {
		return
	}
	r.tmpCnt++
	for _, n := range r.neighbours {
		row := i + n[0]
		if row < 0 || row >= r.rowLen {
			continue
		}
		col := j + n[1]
		if col < 0 || col >= r.colLen {
			continue
		}
		r.dfs(row, col)
	}
}
