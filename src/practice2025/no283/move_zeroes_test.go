package no283

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	//nums := []int{0, 1, 0, 3, 12}
	nums := []int{0}

	moveZeroes(nums)

	fmt.Printf("%v", nums)

}
