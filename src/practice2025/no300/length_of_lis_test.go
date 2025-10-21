package no300

import (
	"fmt"
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	//nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	//nums := []int{0, 1, 0, 3, 2, 3}
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}

	res := lengthOfLIS(nums)

	fmt.Printf("%v", res)

}
