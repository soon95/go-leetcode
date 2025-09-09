package no1

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {

	//nums := []int{2, 7, 11, 15}
	nums := []int{3, 2, 4}
	target := 6

	res := twoSum(nums, target)

	fmt.Printf("%v", res)

}
