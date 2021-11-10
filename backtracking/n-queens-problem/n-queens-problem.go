package n_queens_problem

var result [][]int

/**
 * https://www.youtube.com/watch?v=xFv_Hl4B83A&list=PLDN4rrl48XKpZkf03iYFl-O29szjTrs_O&index=64
 */
func queensProblems(n int) [][]int {
	result = make([][]int, 0)

	backtrack(0, 0, make([]int, n))
	return result
}

// Состояние = какую королеву можно поставить на какую клетку?
func backtrack(row int, col int, current []int) {
	if row == len(current) {
		result = append(result, current)
		return
	}

	for i := 0; i < len(current); i++ {
		if boundFunc(row, i, current) {
			// tmp - for fix problem reset array under slice values
			tmp := make([]int, len(current))
			copy(tmp, current)
			tmp[row] = i
			backtrack(row+1, i, tmp)
		}
	}
	a := 1
	_ = a
}

// check attack
func boundFunc(row int, col int, current []int) bool {
	for i := row - 1; i >= 0; i-- {
		minusRow := row - i
		colCheck := col
		leftCheck := col - minusRow
		rightCheck := col + minusRow
		curPos := current[i]
		if curPos == colCheck || curPos == leftCheck || curPos == rightCheck {
			return false
		}
	}
	return true
}
