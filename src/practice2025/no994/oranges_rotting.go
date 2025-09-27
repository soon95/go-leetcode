package no994

func orangesRotting(grid [][]int) int {

	times := 0

	freshSum := 0

	for {

		curFreshSum := 0
		snapshot := make([][]int, len(grid))
		for i := range snapshot {
			snapshot[i] = make([]int, len(grid[i]))
		}

		for i := range grid {
			for j := range grid[i] {

				if grid[i][j] == 1 {
					curFreshSum++
				}

				if grid[i][j] == 2 || canRotting(grid, i, j) {
					snapshot[i][j] = 2
				} else {
					snapshot[i][j] = grid[i][j]
				}
			}
		}

		if curFreshSum == 0 || curFreshSum == freshSum {
			freshSum = curFreshSum

			break
		}

		grid = snapshot

		freshSum = curFreshSum

		times++
	}

	if freshSum != 0 {
		return -1
	} else {
		return times
	}
}

func canRotting(grid [][]int, i, j int) bool {

	if grid[i][j] == 0 {
		return false
	}

	if i-1 >= 0 && grid[i-1][j] == 2 {
		return true
	}

	if i+1 < len(grid) && grid[i+1][j] == 2 {
		return true
	}

	if j-1 >= 0 && grid[i][j-1] == 2 {
		return true
	}

	if j+1 < len(grid[i]) && grid[i][j+1] == 2 {
		return true
	}

	return false
}
