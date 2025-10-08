package no33

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3

	res := search(nums, target)

	fmt.Printf("%v", res)
}
