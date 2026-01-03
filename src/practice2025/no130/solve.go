package no130

func solve(board [][]byte) {

	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				continue
			}

			if visited[i][j] {
				continue
			}

			if isCover(board, visited, i, j) {
				replace(board, i, j)
			}
		}
	}

}

func isCover(board [][]byte, visited [][]bool, i, j int) bool {

	cover := true
	var dfs func(board [][]byte, visited [][]bool, i, j int)
	dfs = func(board [][]byte, visited [][]bool, i, j int) {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
			return
		}

		if visited[i][j] {
			return
		}

		visited[i][j] = true

		if board[i][j] == 'O' {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
				cover = false
			}
			dfs(board, visited, i-1, j)
			dfs(board, visited, i+1, j)
			dfs(board, visited, i, j-1)
			dfs(board, visited, i, j+1)
		} else {
			return
		}
	}

	dfs(board, visited, i, j)

	return cover
}

func replace(board [][]byte, i, j int) {

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return
	}

	if board[i][j] == 'X' {
		return
	}

	board[i][j] = 'X'

	replace(board, i-1, j)
	replace(board, i+1, j)
	replace(board, i, j-1)
	replace(board, i, j+1)
}
