package no79

func exist(board [][]byte, word string) bool {

	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
		for j := 0; j < len(board[i]); j++ {
			if visit(board, word, i, j, visited) {
				return true
			}
		}
	}

	return false
}

func visit(board [][]byte, word string, i, j int, visited [][]bool) bool {

	if len(word) == 0 {
		return true
	}

	if i < 0 || i >= len(board) {
		return false
	}

	if j < 0 || j >= len(board[i]) {
		return false
	}

	if visited[i][j] {
		return false
	}
	if board[i][j] != word[0] {
		return false
	}

	curVisited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		curVisited[i] = make([]bool, 0)
		curVisited[i] = append(curVisited[i], visited[i]...)
	}

	curVisited[i][j] = true

	return visit(board, word[1:], i+1, j, curVisited) ||
		visit(board, word[1:], i-1, j, curVisited) ||
		visit(board, word[1:], i, j+1, curVisited) ||
		visit(board, word[1:], i, j-1, curVisited)
}
