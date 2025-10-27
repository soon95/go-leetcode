package no31

import (
	"fmt"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	//nums := []int{1, 2, 3}
	//nums := []int{1, 3, 2}
	//nums := []int{2, 1, 3}
	//nums := []int{2, 3, 1}
	//nums := []int{3, 1, 2}
	//nums := []int{3, 2, 1}
	//nums := []int{1, 1, 5}
	nums := []int{4, 2, 0, 2, 3, 2, 0}

	nextPermutation(nums)

	fmt.Printf("%v", nums)

}
