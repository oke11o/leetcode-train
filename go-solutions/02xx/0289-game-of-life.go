package _2xx

/**
https://leetcode.com/problems/game-of-life/
289. Game of Life
Medium
*/
func gameOfLife(board [][]int) {
	rowLen := len(board)
	if rowLen == 0 {
		return
	}
	neighbors := []int{0, 1, -1}
	colLen := len(board[0])
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			// calc neighbors
			liveNeighbors := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if i == 0 && j == 0 {
						continue
					}
					r := row + neighbors[i]
					c := col + neighbors[j]
					if r < 0 || r >= rowLen || c < 0 || c >= colLen {
						continue
					}
					if board[r][c] == 1 || board[r][c] == 2 || board[r][c] == 3 {
						liveNeighbors++
					}
					// 1. это сосед снизу - 0,1
					// 2. сосед сверху. Он уже изменени.
					// Из 1 мы можем или жить или умереть. То есть перейти в 2 или в 3.
				}
			}
			if board[row][col] == 1 {
				if liveNeighbors == 2 || liveNeighbors == 3 {
					board[row][col] = 3 // останется жить
				} else {
					board[row][col] = 2 // умрет
				}
			}
			if board[row][col] == 0 && liveNeighbors == 3 {
				board[row][col] = 4 // оживить
			}

		}
	}

	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if board[row][col] == 1 || board[row][col] == 3 || board[row][col] == 4 {
				board[row][col] = 1
			} else {
				board[row][col] = 0
			}
		}
	}
}
