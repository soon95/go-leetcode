package no48

func rotate(matrix [][]int) {

	length := len(matrix)

	for i := 0; i < length/2; i++ {

		for j := 0; j < (length+1)/2; j++ {

			tmp := matrix[i][j]

			matrix[i][j] = matrix[length-1-j][i]
			matrix[length-1-j][i] = matrix[length-1-i][length-1-j]
			matrix[length-1-i][length-1-j] = matrix[j][length-1-i]
			matrix[j][length-1-i] = tmp
		}
	}

	return
}
