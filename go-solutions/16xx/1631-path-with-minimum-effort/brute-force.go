package _631_path_with_minimum_effort

/**
https://leetcode.com/problems/path-with-minimum-effort/
1631. Path With Minimum Effort
Medium
*/
func minimumEffortPath(heights [][]int) int {
	var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	const mathMax = 1<<31 - 1
	var maxSoFar = mathMax

	var isValidCell = func(x, y, rowLen, colLen int) bool {
		return x >= 0 && x <= rowLen-1 && y >= 0 && y <= colLen-1
	}

	var backtrack func(x, y int, heights [][]int, maxDifference int) int
	backtrack = func(x, y int, heights [][]int, maxDifference int) int {
		rowLen := len(heights)
		colLen := len(heights[0])
		if x == rowLen-1 && y == colLen-1 {
			if maxSoFar > maxDifference {
				maxSoFar = maxDifference
			}
			return maxDifference
		}

		currentHeight := heights[x][y]
		heights[x][y] = 0

		minEffort := mathMax
		for i := 0; i < 4; i++ {
			adjacentX := x + directions[i][0]
			adjacentY := y + directions[i][1]
			// heights[adjacentX][adjacentY] != 0 - is not Visited
			if isValidCell(adjacentX, adjacentY, rowLen, colLen) && heights[adjacentX][adjacentY] != 0 {
				currentDifference := heights[adjacentX][adjacentY] - currentHeight
				if currentDifference < 0 {
					currentDifference = -currentDifference
				}
				maxCurrentDifference := maxDifference
				if maxCurrentDifference < currentDifference {
					maxCurrentDifference = currentDifference
				}
				if maxCurrentDifference < maxSoFar {
					result := backtrack(adjacentX, adjacentY, heights, maxCurrentDifference)
					if minEffort > result {
						minEffort = result
					}
				}
			}
		}

		heights[x][y] = currentHeight
		return minEffort
	}

	return backtrack(0, 0, heights, 0)
}
