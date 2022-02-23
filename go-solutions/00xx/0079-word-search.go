package _0xx

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && traverse(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

func traverse(board [][]byte, i, j int, word string) bool {
	if len(word) == 0 {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[0] {
		return false
	}

	orig_value := board[i][j]
	board[i][j] = '_'

	match := traverse(board, i-1, j, word[1:]) || traverse(board, i+1, j, word[1:]) || traverse(board, i, j-1, word[1:]) || traverse(board, i, j+1, word[1:])

	board[i][j] = orig_value

	return match
}
