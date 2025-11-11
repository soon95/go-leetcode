package no209

import (
	"fmt"
	"testing"
)

func Test_minSubArrayLen(t *testing.T) {
	target := 7
	nums := []int{2, 3, 1, 2, 7, 3}

	res := minSubArrayLen(target, nums)

	fmt.Printf("%v", res)
}
