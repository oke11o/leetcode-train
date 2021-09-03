package go_solutions

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	var (
		up     = 0
		bottom = len(matrix) - 1
		left   = 0
		right  = len(matrix[0]) - 1
	)

	direction := 0 // 0 - to left, 1 - to bottom; 2 - to right; 3 - to up

	for up <= bottom && left <= right {
		switch direction {
		case 0:
			for i := left; i <= right; i++ {
				result = append(result, matrix[up][i])
			}
			up++
			direction = 1
		case 1:
			for i := up; i <= bottom; i++ {
				result = append(result, matrix[i][right])
			}
			right--
			direction = 2
		case 2:
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
			direction = 3
		case 3:
			for i := bottom; i >= up; i-- {
				result = append(result, matrix[i][left])
			}
			left++
			direction = 0
		}
	}

	return result
}
