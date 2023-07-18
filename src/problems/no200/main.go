package main

import "fmt"

var islands int
var visited [][]bool

/*
*200. 岛屿数量
https://leetcode.cn/problems/number-of-islands/
*/
func numIslands(grid [][]byte) int {

	islands = 0

	m := len(grid)
	n := len(grid[0])

	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if grid[i][j] == '1' && !visited[i][j] {
				islands++

				visit(grid, i, j)
			}

		}
	}

	return islands
}

func visit(grid [][]byte, x, y int) {

	m := len(grid)
	n := len(grid[0])

	if x < 0 || x >= m || y < 0 || y >= n {
		return
	}

	if visited[x][y] {
		return
	}

	visited[x][y] = true

	if grid[x][y] == '1' {
		// 如果是陆地

		visit(grid, x-1, y)
		visit(grid, x+1, y)
		visit(grid, x, y-1)
		visit(grid, x, y+1)
	}

}

func main() {

	//grid := [][]byte{
	//	{'1', '1', '1', '1', '0'},
	//	{'1', '1', '0', '1', '0'},
	//	{'1', '1', '0', '0', '0'},
	//	{'0', '0', '0', '0', '0'},
	//}

	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	fmt.Printf("%v", numIslands(grid))

}
