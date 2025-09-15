package no189

import (
	"fmt"
	"testing"
)

func Test_rotate(t *testing.T) {
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//k := 3

	nums := []int{-1, -100, 3, 99}
	k := 2

	rotate(nums, k)

	fmt.Printf("%v", nums)

}
