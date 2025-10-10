package no4

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	//nums1 := []int{1, 3}
	//nums2 := []int{2}

	//nums1 := []int{1, 2}
	//nums2 := []int{3, 4}

	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	res := findMedianSortedArrays(nums1, nums2)

	fmt.Printf("%v", res)
}
