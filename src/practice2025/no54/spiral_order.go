package no54

func spiralOrder(matrix [][]int) []int {

	res := make([]int, 0)

	visited := make([][]bool, len(matrix))
	for i, ints := range matrix {
		visited[i] = make([]bool, len(ints))
	}

	i, j := 0, 0

	stepI, stepJ := 0, 1

	for i < len(matrix) && i >= 0 && j < len(matrix[0]) && j >= 0 && !visited[i][j] {

		res = append(res, matrix[i][j])

		visited[i][j] = true

		if stepJ == 1 {
			if j+1 == len(matrix[i]) || visited[i][j+1] {
				stepI = 1
				stepJ = 0
			}
		} else if stepI == 1 {
			if i+1 == len(matrix) || visited[i+1][j] {
				stepI = 0
				stepJ = -1
			}
		} else if stepJ == -1 {
			if j-1 < 0 || visited[i][j-1] {
				stepI = -1
				stepJ = 0
			}
		} else if stepI == -1 {
			if i-1 < 0 || visited[i-1][j] {
				stepI = 0
				stepJ = 1
			}
		}

		i = i + stepI
		j = j + stepJ
	}
	return res
}
