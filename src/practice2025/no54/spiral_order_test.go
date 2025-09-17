package no54

import (
	"fmt"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	//matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{{1}, {2}}

	res := spiralOrder(matrix)

	fmt.Printf("%v", res)
}
