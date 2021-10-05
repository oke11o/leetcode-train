package _7xx

// 0733. Flood Fill (BFS)
func floodFill_bfs(image [][]int, sr int, sc int, newColor int) [][]int {
	rowLen := len(image)
	colLen := len(image[0])
	visited := make([][]bool, rowLen)
	for i := 0; i < rowLen; i++ {
		visited[i] = make([]bool, colLen)
	}

	oldColor := image[sr][sc]
	bfs := [][2]int{{sr, sc}}
	for len(bfs) > 0 {
		cur := bfs[0]
		bfs = bfs[1:]
		curRow := cur[0]
		if curRow < 0 || curRow >= rowLen {
			continue
		}
		curCol := cur[1]
		if curCol < 0 || curCol >= colLen {
			continue
		}
		if visited[curRow][curCol] {
			continue
		}
		if image[curRow][curCol] == oldColor {
			image[curRow][curCol] = newColor
			bfs = append(bfs, [2]int{curRow - 1, curCol})
			bfs = append(bfs, [2]int{curRow + 1, curCol})
			bfs = append(bfs, [2]int{curRow, curCol - 1})
			bfs = append(bfs, [2]int{curRow, curCol + 1})
		}
		visited[curRow][curCol] = true
	}
	return image
}

// 0733. Flood Fill (DFS)
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	f := newFloodFillDFS(image, sr, sc, newColor)
	return f.fill()
}

type floodFillDFS struct {
	image      [][]int
	sr         int
	sc         int
	newColor   int
	oldColor   int
	neighbours [4][2]int
	rowLen     int
	colLen     int
	visited    [][]bool
}

func newFloodFillDFS(image [][]int, sr int, sc int, newColor int) *floodFillDFS {
	f := &floodFillDFS{
		image:      image,
		sr:         sr,
		sc:         sc,
		newColor:   newColor,
		oldColor:   image[sr][sc],
		neighbours: [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
		rowLen:     len(image),
		colLen:     len(image[0]),
	}

	f.visited = make([][]bool, f.rowLen)
	for i := 0; i < f.rowLen; i++ {
		f.visited[i] = make([]bool, f.colLen)
	}
	return f
}

func (f *floodFillDFS) fill() [][]int {
	f.image[f.sr][f.sc] = f.newColor
	f.dfs(f.sr, f.sc)
	return f.image
}

func (f *floodFillDFS) dfs(i int, j int) {
	f.visited[i][j] = true
	for _, dv := range f.neighbours {
		newRow := i + dv[0]
		if newRow < 0 || newRow >= f.rowLen {
			continue
		}
		newCol := j + dv[1]
		if newCol < 0 || newCol >= f.colLen {
			continue
		}
		if !f.visited[newRow][newCol] && f.image[newRow][newCol] == f.oldColor {
			f.image[newRow][newCol] = f.newColor
			f.dfs(newRow, newCol)
		}
	}
}
