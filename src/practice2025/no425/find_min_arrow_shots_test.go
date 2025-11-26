package no425

import (
	"fmt"
	"testing"
)

func Test_findMinArrowShots(t *testing.T) {
	//points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

	res := findMinArrowShots(points)

	fmt.Printf("%v", res)
}
