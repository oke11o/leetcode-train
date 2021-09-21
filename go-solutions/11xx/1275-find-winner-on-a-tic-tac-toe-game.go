package _1xx

func tictactoe(moves [][]int) string {
	dashboard := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	raws := []int{0, 0, 0}
	cols := []int{0, 0, 0}
	mainDiagonal := 0
	backDiagonal := 0
	for stepIdx, step := range moves {
		var val int
		if stepIdx%2 == 0 {
			val = 1 // A
		} else {
			val = 10 // B
		}
		i, j := step[0], step[1]
		dashboard[i][j] = val
		raws[i] += dashboard[i][j]
		cols[j] += dashboard[i][j]
		if i == j {
			mainDiagonal += dashboard[i][j]
		}
		if j+i == 2 {
			backDiagonal += dashboard[i][j]
		}
	}
	if mainDiagonal == 3 || backDiagonal == 3 {
		return "A"
	}
	if mainDiagonal == 30 || backDiagonal == 30 {
		return "B"
	}
	for _, di := range [][]int{raws, cols} {
		for _, v := range di {
			if v == 3 {
				return "A"
			} else if v == 30 {
				return "B"
			}
		}
	}
	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}
