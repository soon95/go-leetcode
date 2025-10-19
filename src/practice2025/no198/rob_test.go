package no198

import (
	"fmt"
	"testing"
)

func Test_rob(t *testing.T) {
	//nums := []int{1, 2, 3, 1}
	nums := []int{2, 7, 9, 3, 1}

	res := rob(nums)

	fmt.Printf("%v", res)
}
