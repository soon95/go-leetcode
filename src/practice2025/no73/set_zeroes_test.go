package no73

import (
	"fmt"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}

	setZeroes(matrix)

	fmt.Printf("%+v", matrix)

}
