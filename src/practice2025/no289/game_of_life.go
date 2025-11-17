package no289

func gameOfLife(board [][]int) {

	newBoard := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		newBoard[i] = make([]int, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			newBoard[i][j] = determineLife(board, i, j)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = newBoard[i][j]
		}
	}
}

func determineLife(board [][]int, m, n int) int {

	liveCellCnt := 0

	for i := max(0, m-1); i <= min(len(board)-1, m+1); i++ {

		for j := max(0, n-1); j <= min(len(board[i])-1, n+1); j++ {

			if i == m && j == n {
				continue
			}
			if board[i][j] == 1 {
				liveCellCnt++
			}
		}
	}

	if board[m][n] == 1 {
		if liveCellCnt >= 2 && liveCellCnt <= 3 {
			return 1
		} else {
			return 0
		}
	} else {

		if liveCellCnt == 3 {
			return 1
		} else {
			return 0
		}
	}
}
