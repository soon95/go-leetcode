package no11

import (
	"fmt"
	"testing"
)

func Test_maxArea(t *testing.T) {

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	res := maxArea(height)

	fmt.Printf("%v", res)
}
