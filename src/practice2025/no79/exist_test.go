package no79

import (
	"fmt"
	"testing"
)

func Test_exist(t *testing.T) {
	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//word := "ABCCED"
	//word := "SEE"
	//word := "ABCB"

	board := [][]byte{{'A', 'B'}, {'C', 'D'}}

	word := "ACDB"

	res := exist(board, word)

	fmt.Printf("%v", res)

}
