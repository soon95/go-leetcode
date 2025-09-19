package no240

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	target := 2
	res := searchMatrix(matrix, target)

	fmt.Printf("%v", res)
}
