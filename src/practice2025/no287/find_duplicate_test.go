package no287

import (
	"fmt"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}

	res := findDuplicate(nums)

	fmt.Printf("%v", res)

}
