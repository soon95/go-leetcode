package no240

func searchMatrix(matrix [][]int, target int) bool {

	for i := 0; i < len(matrix); i++ {

		if target < matrix[i][0] {
			return false
		}

		if target > matrix[i][len(matrix[0])-1] {
			continue
		}

		for j := 0; j < len(matrix[0]); j++ {
			if target == matrix[i][j] {
				return true
			}
		}

	}

	return false
}
