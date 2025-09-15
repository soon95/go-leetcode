package no56

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//intervals := [][]int{{1, 4}, {4, 5}}
	intervals := [][]int{{4, 7}, {1, 4}}

	res := merge(intervals)

	fmt.Printf("%+v", res)
}
