package _7xx

func movesToChessboard(board [][]int) int {
	row, col, valid := chessboardValid(board)
	if !valid {
		return -1
	}
	_, _ = row, col

	return 0
}

func chessboardValid(board [][]int) (row int, col int, valid bool) {
	var rowSum, colSum int
	for i := 0; i < len(board); i++ {
		row <<= 1
		if board[i][0] == 1 {
			row++
		}
		col <<= 1
		if board[0][i] == 1 {
			col++
		}
		rowSum += board[i][0]
		colSum += board[0][i]
	}
	middle := len(board) / 2
	if !(middle == rowSum || middle == rowSum+1) {
		return
	}
	if !(middle == colSum || middle == colSum+1) {
		return
	}

	valid = true

	return
}

func bitSlice2int(bits []int) int {
	if len(bits) == 0 {
		return 0
	}
	if len(bits) == 1 {
		return bits[0]
	}
	result := 0
	for i := 0; i < len(bits); i++ {
		result <<= 1
		if bits[i] == 1 {
			result++
		}
	}
	return result
}
