package no228

import (
	"fmt"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}

	res := summaryRanges(nums)

	fmt.Printf("%v", res)

}
