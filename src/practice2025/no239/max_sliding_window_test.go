package no239

import (
	"fmt"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 2

	res := maxSlidingWindow(nums, k)

	fmt.Printf("%v", res)

}
