package _0xx

func setZeroes(matrix [][]int) {

	var row_zero, column_zero bool

	for r := 0; r < len(matrix); r++ {
		if matrix[r][0] == 0 {
			row_zero = true
			break
		}
	}
	for c := 0; c < len(matrix[0]); c++ {
		if matrix[0][c] == 0 {
			column_zero = true
			break
		}
	}

	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			if matrix[r][c] == 0 {
				matrix[r][0] = 0
				matrix[0][c] = 0
			}
		}
	}

	for r := 1; r < len(matrix); r++ {
		for c := 1; c < len(matrix[0]); c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if row_zero {
		for r := 0; r < len(matrix); r++ {
			matrix[r][0] = 0
		}
	}
	if column_zero {
		for c := 0; c < len(matrix[0]); c++ {
			matrix[0][c] = 0
		}
	}
}

// Use extra memory for coords
func setZeroes_2(matrix [][]int) {
	type pair struct {
		m, n int
	}
	var zeros []pair
	for m, row := range matrix {
		for n, val := range row {
			if val == 0 {
				zeros = append(zeros, pair{m: m, n: n})
			}
		}
	}
	for _, p := range zeros {
		for m := 0; m < len(matrix); m++ {
			matrix[m][p.n] = 0
		}
		for n := 0; n < len(matrix[0]); n++ {
			matrix[p.m][n] = 0
		}
	}
}
