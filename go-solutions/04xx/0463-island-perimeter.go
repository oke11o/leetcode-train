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
