package no46

import (
	"fmt"
	"testing"
)

func Test_permute(t *testing.T) {
	nums := []int{1, 2, 3}
	//nums := []int{1}

	res := permute(nums)

	fmt.Printf("%v", res)
}
