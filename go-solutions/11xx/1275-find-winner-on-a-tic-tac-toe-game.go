package _1xx

func tictactoe(moves [][]int) string {
	dashboard := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	for stepIdx, step := range moves {
		val := int(0)
		if stepIdx%2 == 0 {
			val = 1 // A
		} else {
			val = 10 // B
		}
		dashboard[step[0]][step[1]] = val
	}
	diag := []int{dashboard[0][0] + dashboard[1][1] + dashboard[2][2], dashboard[2][0] + dashboard[1][1] + dashboard[0][2]}
	if diag[0] == 3 || diag[1] == 3 {
		return "A"
	}
	if diag[0] == 30 || diag[1] == 30 {
		return "B"
	}
	raws := []int{0, 0, 0}
	cols := []int{0, 0, 0}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			raws[i] += dashboard[i][j]
			cols[j] += dashboard[i][j]
		}
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
