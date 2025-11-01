package no80

import (
	"fmt"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	//nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}

	res := removeDuplicates(nums)

	fmt.Printf("%v", res)

}
