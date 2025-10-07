package no51

import (
	"fmt"
	"testing"
)

func Test_solveNQueens(t *testing.T) {
	n := 4

	res := solveNQueens(n)

	fmt.Printf("%v", res)
}
