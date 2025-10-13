package no215

import (
	"fmt"
	"testing"
)

func Test_findKthLargest(t *testing.T) {
	//nums := []int{3, 2, 1, 5, 6, 4}
	//k := 2
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4

	res := findKthLargest(nums, k)

	fmt.Printf("%v", res)

}
