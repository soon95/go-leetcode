package no153

import (
	"fmt"
	"testing"
)

func Test_findMin(t *testing.T) {
	//nums := []int{3, 4, 5, 1, 2}
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{3, 1, 2}

	res := findMin(nums)

	fmt.Printf("%v", res)

}
