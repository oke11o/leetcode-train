package _4xx

// 0463. Island Perimeter
//
func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count += 4
				if j-1 >= 0 && grid[i][j-1] == 1 {
					count -= 2
				}
				if i-1 >= 0 && grid[i-1][j] == 1 {
					count -= 2
				}
			}
		}
	}
	return count
}

func islandPerimeter_simpleCounting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				if i == 0 {
					count++
				} else if grid[i-1][j] == 0 {
					count++
				}
				if i == len(grid)-1 {
					count++
				} else if grid[i+1][j] == 0 {
					count++
				}
				if j == 0 {
					count++
				} else if grid[i][j-1] == 0 {
					count++
				}
				if j == len(grid[0])-1 {
					count++
				} else if grid[i][j+1] == 0 {
					count++
				}
			}
		}
	}
	return count
}

func islandPerimeter_dfs(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	p := newIslandPerimeterDFSer(grid)
	return p.islandPerimeter()
}

type islandPerimeterDFSer struct {
	grid    [][]int
	visited [][]bool
	rowLen  int // that 's ok
	colLen  int // column
	res     int // Island perimeter
	dirx    []int
	diry    []int

	// global variable
}

func newIslandPerimeterDFSer(grid [][]int) *islandPerimeterDFSer {
	p := islandPerimeterDFSer{
		grid:   grid,
		rowLen: len(grid),
		colLen: len(grid[0]),
	}
	p.visited = make([][]bool, p.rowLen)
	for i := 0; i < p.rowLen; i++ {
		p.visited[i] = make([]bool, p.colLen)
	}
	p.dirx = []int{-1, 1, 0, 0}
	p.diry = []int{0, 0, -1, 1}
	return &p
}

func (p *islandPerimeterDFSer) islandPerimeter() int {
	for i := 0; i < p.rowLen; i++ {
		for j := 0; j < p.colLen; j++ {
			if p.grid[i][j] == 1 && !p.visited[i][j] {
				p.dfs(i, j)
			}
		}
	}
	return p.res
}

func (p *islandPerimeterDFSer) dfs(x, y int) {
	p.visited[x][y] = true // sign
	for i := 0; i < 4; i++ {
		xx := x + p.dirx[i]
		yy := y + p.diry[i]
		// If the current location is land, and the "new location" extended from the current location is a "puddle"
		// or "new location" beyond the boundary, an edge will be contributed to the perimeter
		if xx < 0 || xx >= p.rowLen || yy < 0 || yy >= p.colLen || p.grid[xx][yy] == 0 {
			p.res++
			continue
		}
		if p.visited[xx][yy] { // An edge is not contributed to the perimeter
			continue
		}
		p.dfs(xx, yy)
	}
}

func islandPerimeter_bfs(grid [][]int) int {
	bfs := [][2]int{}
	dir := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	found := false
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				bfs = append(bfs, [2]int{i, j})
				found = true
			}
		}
		if found {
			break
		}
	}

	count := 0
	for len(bfs) != 0 {
		q := bfs[0]
		bfs = bfs[1:]
		cx := q[0]
		cy := q[1]
		if grid[cx][cy] == 0 {
			continue
		}
		count += 4
		grid[cx][cy] = 0
		for i := 0; i < 4; i++ {
			nx := cx + dir[i][0]
			ny := cy + dir[i][1]
			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == 1 {
				bfs = append(bfs, [2]int{nx, ny})
				count -= 2
			}
		}
	}
	return count
}
