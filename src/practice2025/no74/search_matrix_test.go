package no74

import (
	"fmt"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	target := 13

	res := searchMatrix(matrix, target)

	fmt.Printf("%v", res)

}
