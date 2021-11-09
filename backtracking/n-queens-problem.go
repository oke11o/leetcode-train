package backtracking

var queenProblemsResult [][]int

/**
 * https://www.youtube.com/watch?v=xFv_Hl4B83A&list=PLDN4rrl48XKpZkf03iYFl-O29szjTrs_O&index=64
 */
func queensProblems(n int) [][]int {
	queenProblemsResult = make([][]int, 0)

	queensProblemsSolution()
	return queenProblemsResult
}

func queensProblemsSolution() {
	// Состояние = какую королеву можно поставить на какую клетку?

	for q1 := 0; q1 < 4; q1++ {

	}
	//
}
