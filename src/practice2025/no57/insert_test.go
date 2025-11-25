package no57

import (
	"fmt"
	"testing"
)

func Test_insert(t *testing.T) {
	//intervals := [][]int{{1, 3}, {6, 9}}
	//newInterval := []int{2, 5}
	//intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	intervals := [][]int{{3, 5}, {12, 15}}
	newInterval := []int{6, 6}

	res := insertV2(intervals, newInterval)

	fmt.Printf("%v", res)

}
