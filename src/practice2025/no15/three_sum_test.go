package no15

import (
	"fmt"
	"testing"
)

func Test_threeSum(t *testing.T) {

	//nums := []int{-1, 0, 1, 2, -1, -4}
	//nums := []int{0, 1, 1}
	//nums := []int{0, 0, 0, 0}
	nums := []int{-10, -7, -6, 11, 12, 13, 16}

	res := threeSum(nums)

	fmt.Printf("%+v", res)

}
