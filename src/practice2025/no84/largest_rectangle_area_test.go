package no84

import (
	"fmt"
	"testing"
)

func Test_largestRectangleArea(t *testing.T) {
	heights := []int{2, 1, 5, 6, 2, 3}
	//heights := []int{4, 2, 0, 3, 2, 5}
	//heights := []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}

	res := largestRectangleArea(heights)

	fmt.Printf("%v", res)

}
