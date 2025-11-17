package no289

import (
	"fmt"
	"testing"
)

func Test_gameOfLife(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}

	gameOfLife(board)

	fmt.Printf("%v", board)

}
