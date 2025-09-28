package no207

import (
	"fmt"
	"testing"
)

func Test_canFinish(t *testing.T) {
	numCourses := 2
	//prerequisites := [][]int{{1, 0}}
	prerequisites := [][]int{{1, 0}, {0, 1}}

	res := canFinish(numCourses, prerequisites)

	fmt.Printf("%v", res)

}
