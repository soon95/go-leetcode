package no994

import (
	"fmt"
	"testing"
)

func Test_orangesRotting(t *testing.T) {
	//grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	grid := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}

	res := orangesRotting(grid)

	fmt.Printf("%v", res)

}
