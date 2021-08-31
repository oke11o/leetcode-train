package go_solutions

func setZeroes(matrix [][]int) {
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
