package no51

func solveNQueens(n int) [][]string {

	ans = make([][]string, 0)

	grid := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]bool, n)
	}

	dfs(grid, n, 0, n)

	return ans
}

var ans [][]string

func dfs(grid [][]bool, n, startIndex, restQ int) {

	if restQ == 0 {
		tmp := make([]string, 0)
		for i := 0; i < len(grid); i++ {
			row := ""
			for j := 0; j < len(grid[i]); j++ {
				if grid[i][j] {
					row += "Q"
				} else {
					row += "."
				}
			}
			tmp = append(tmp, row)
		}
		ans = append(ans, tmp)
	}

	for i := startIndex; i < n*n; i++ {

		x := i / n
		y := i % n

		if canPut(grid, x, y) {
			grid[x][y] = true
			dfs(grid, n, i+1, restQ-1)
			grid[x][y] = false
		}
	}
}

func canPut(grid [][]bool, i, j int) bool {

	for k := 0; k < i; k++ {
		if grid[k][j] {
			return false
		}
	}

	for k := 0; k < j; k++ {
		if grid[i][k] {
			return false
		}
	}

	x, y := i-1, j-1
	for x >= 0 && y >= 0 {
		if grid[x][y] {
			return false
		}
		x--
		y--
	}

	x, y = i-1, j+1
	for x >= 0 && y < len(grid) {
		if grid[x][y] {
			return false
		}
		x--
		y++
	}

	return true
}
