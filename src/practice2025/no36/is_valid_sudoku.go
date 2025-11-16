package no36

func isValidSudoku(board [][]byte) bool {

	for i := 0; i < len(board); i++ {
		tmp := make(map[string]bool)
		for j := 0; j < len(board); j++ {
			if board[i][j] == '.' {
				continue
			}

			if _, exist := tmp[string(board[i][j])]; exist {
				return false
			} else {
				tmp[string(board[i][j])] = true
			}
		}
	}

	for j := 0; j < len(board); j++ {
		tmp := make(map[string]bool)
		for i := 0; i < len(board); i++ {
			if board[i][j] == '.' {
				continue
			}
			if _, exist := tmp[string(board[i][j])]; exist {
				return false
			} else {
				tmp[string(board[i][j])] = true
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := make(map[string]bool)
			for m := 3 * i; m < 3*(i+1); m++ {
				for n := 3 * j; n < 3*(j+1); n++ {
					if board[m][n] == '.' {
						continue
					}

					if _, exist := tmp[string(board[m][n])]; exist {
						return false
					} else {
						tmp[string(board[m][n])] = true
					}
				}
			}
		}

	}

	return true
}
