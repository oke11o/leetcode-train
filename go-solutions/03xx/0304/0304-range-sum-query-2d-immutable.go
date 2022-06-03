package _304

/**
https://leetcode.com/problems/range-sum-query-2d-immutable/
304. Range Sum Query 2D - Immutable
Medium
*/
type NumMatrix struct {
	matrix [][]int
	cache  [][]int
}

/**
Вообщем, если делать кэш размерностью mn, то приходится дело оооочнь много проверок на предыдущие значения.
Как на построении, так и на поиске.
Но если сделать первый row/col пустыми, то все становится сильно проще.
Только надо не запутаться, что значение в cache сдвинуты на +1 вправо и вниз
*/
func Constructor(matrix [][]int) NumMatrix {
	cache := make([][]int, len(matrix)+1)
	cache[0] = make([]int, len(matrix[0])+1)

	for row := 0; row < len(matrix); row++ {
		cache[row+1] = make([]int, len(matrix[0])+1)
		for col := 0; col < len(matrix[0]); col++ {
			// почему надо вычитать верхнее левое можно легко понять, когда строишь первый квадрат 2х2
			cache[row+1][col+1] = matrix[row][col] + cache[row][col+1] + cache[row+1][col] - cache[row][col]
		}
	}

	return NumMatrix{
		matrix: matrix,
		cache:  cache,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	return this.cache[row2+1][col2+1] - this.cache[row2+1][col1] - this.cache[row1][col2+1] + this.cache[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
