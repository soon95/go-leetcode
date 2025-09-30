package no78

import (
	"fmt"
	"testing"
)

func Test_subsets(t *testing.T) {
	nums := []int{1, 2, 3}
	//nums := []int{1}

	res := subsets(nums)

	fmt.Printf("%v", res)
}
