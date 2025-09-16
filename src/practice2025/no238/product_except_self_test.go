package no238

import (
	"fmt"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	//nums := []int{1, 2, 3, 4}
	nums := []int{-1, 1, 0, -3, 3}
	res := productExceptSelf(nums)

	fmt.Printf("%v", res)
}
