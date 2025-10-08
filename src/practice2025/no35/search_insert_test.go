package no35

import (
	"fmt"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 7

	res := searchInsert(nums, target)
	fmt.Printf("%v\n", res)

}
