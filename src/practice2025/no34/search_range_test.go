package no34

import (
	"fmt"
	"testing"
)

func Test_searchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6

	res := searchRange(nums, target)

	fmt.Printf("%v", res)
}
