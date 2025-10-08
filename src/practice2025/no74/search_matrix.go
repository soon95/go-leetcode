package no74

func searchMatrix(matrix [][]int, target int) bool {

	n := len(matrix) * len(matrix[0])

	start, end := 0, n-1

	for start <= end {

		mid := (start + end) / 2

		x := mid / len(matrix[0])
		y := mid % len(matrix[0])

		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}

	}

	return false
}
