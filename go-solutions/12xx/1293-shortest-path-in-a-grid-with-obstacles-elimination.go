package _2xx

type Position struct {
	r, c, k int
}

type Queue struct {
	positions []Position
}

func (l *Queue) Poll() Position {
	if l.IsEmpty() {
		return Position{}
	}
	v := l.positions[0]
	l.positions = l.positions[1:]
	return v
}

func (l *Queue) Size() int {
	return len(l.positions)
}

func (l *Queue) Offer(v Position) {
	l.positions = append(l.positions, v)
}

func (l *Queue) IsEmpty() bool {
	return len(l.positions) == 0
}

// 1293. Shortest Path in a Grid with Obstacles Elimination
// 1. BFS
// Time: O(MNK) && Space: O(MNK)
func shortestPath(grid [][]int, k int) int {
	var initVisited = func(n, m, k int) [][][]bool {
		visited := make([][][]bool, n)
		for i := 0; i < n; i++ {
			visited[i] = make([][]bool, m)
			for j := 0; j < m; j++ {
				visited[i][j] = make([]bool, k)
			}
		}
		return visited
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	n := len(grid)
	m := len(grid[0])
	q := Queue{}
	visited := initVisited(n, m, k+1)
	visited[0][0][0] = true
	q.Offer(Position{0, 0, 0})

	result := 0
	for !q.IsEmpty() {
		size := q.Size()
		for i := 0; i < size; i++ {
			info := q.Poll()
			r := info.r
			c := info.c
			curK := info.k

			if r == n-1 && c == m-1 {
				return result
			}

			for _, dir := range dirs {
				nextR := dir[0] + r
				nextC := dir[1] + c
				nextK := curK
				if nextR >= 0 && nextR < n && nextC >= 0 && nextC < m {
					if grid[nextR][nextC] == 1 {
						nextK++
					}
					if nextK <= k && !visited[nextR][nextC][nextK] {
						visited[nextR][nextC][nextK] = true
						q.Offer(Position{r: nextR, c: nextC, k: nextK})
					}
				}
			}
		}
		result++
	}

	return -1
}

// TODO: BFS
