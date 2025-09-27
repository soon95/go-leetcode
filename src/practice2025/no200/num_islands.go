package no200

func numIslands(grid [][]byte) int {

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[i]))
	}

	res := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}
			if visited[i][j] {
				continue
			}

			res++

			spread(grid, visited, i, j)
		}
	}
	return res
}

func spread(grid [][]byte, visited [][]bool, i, j int) {

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}

	if grid[i][j] == '0' {
		return
	}

	if visited[i][j] {
		return
	}

	visited[i][j] = true

	spread(grid, visited, i+1, j)
	spread(grid, visited, i-1, j)
	spread(grid, visited, i, j+1)
	spread(grid, visited, i, j-1)
}
