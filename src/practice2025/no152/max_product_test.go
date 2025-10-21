package no152

import (
	"fmt"
	"testing"
)

func Test_maxProduct(t *testing.T) {
	//nums := []int{2, 3, -2, 4}
	nums := []int{-2, 0, -1}

	res := maxProduct(nums)

	fmt.Printf("%v", res)
}
