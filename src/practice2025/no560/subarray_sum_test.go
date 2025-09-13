package no560

import (
	"fmt"
	"testing"
)

func Test_subarraySum(t *testing.T) {
	//nums := []int{1, 1, 1}
	//k := 2
	//nums := []int{1, 2, 3}
	//k := 2
	nums := []int{-1, -1, 1}
	k := 0

	res := subarraySum(nums, k)

	fmt.Printf("%v", res)

}
