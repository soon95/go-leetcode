package no128

import (
	"fmt"
	"testing"
)

func Test_longestConsecutive(t *testing.T) {

	//nums := []int{100, 4, 200, 1, 3, 2}
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}

	res := longestConsecutive(nums)

	fmt.Printf("%v\n", res)

}
