package no41

import (
	"fmt"
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	//nums := []int{1, 2, 0}
	//nums := []int{3, 4, -1, 1}
	//nums := []int{7, 8, 9, 11, 12}
	nums := []int{0, 2, 2, 1, 1}
	res := firstMissingPositive(nums)

	fmt.Printf("%v", res)
}
