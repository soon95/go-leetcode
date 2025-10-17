package no55

import (
	"fmt"
	"testing"
)

func Test_canJump(t *testing.T) {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{3, 2, 1, 0, 4}

	res := canJump(nums)

	fmt.Printf("%v", res)

}
