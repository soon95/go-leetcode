package no53

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	//nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	nums := []int{-1}
	//nums := []int{5, 4, -1, 7, 8}

	res := maxSubArray(nums)

	fmt.Printf("%v", res)

}
