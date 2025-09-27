package no200

import (
	"fmt"
	"testing"
)

func Test_numIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	res := numIslands(grid)

	fmt.Printf("%v", res)

}
