package no45

import (
	"fmt"
	"testing"
)

func Test_jump(t *testing.T) {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{0}

	res := jump(nums)

	fmt.Printf("%v", res)
}
