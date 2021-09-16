package go_solutions

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0, len(matrix)*len(matrix[0]))
	direction := "left"
	topEdge := 0
	leftEdge := 0
	bottomEdge := len(matrix) - 1
	rightEdge := len(matrix[0]) - 1

	for topEdge <= bottomEdge && leftEdge <= rightEdge {
		switch direction {
		case "left":
			for i := leftEdge; i <= rightEdge; i++ {
				result = append(result, matrix[topEdge][i])
			}
			topEdge++
			direction = "down"
		case "down":
			for i := topEdge; i <= bottomEdge; i++ {
				result = append(result, matrix[i][rightEdge])
			}
			rightEdge--
			direction = "right"
		case "right":
			for i := rightEdge; i >= leftEdge; i-- {
				result = append(result, matrix[bottomEdge][i])
			}
			bottomEdge--
			direction = "up"
		case "up":
			for i := bottomEdge; i >= topEdge; i-- {
				result = append(result, matrix[i][leftEdge])
			}
			leftEdge++
			direction = "left"
		}
	}
	return result
}
