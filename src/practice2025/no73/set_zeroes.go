package no73

func setZeroes(matrix [][]int) {

	zeroI := make([]int, 0)
	zeroJ := make([]int, 0)

	for i, ints := range matrix {
		for j, num := range ints {
			if num == 0 {
				zeroI = append(zeroI, i)
				zeroJ = append(zeroJ, j)
			}
		}
	}
	for _, i := range zeroI {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = 0
		}
	}

	for _, j := range zeroJ {
		for i := 0; i < len(matrix); i++ {
			matrix[i][j] = 0
		}
	}

}
