package no347

import (
	"fmt"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	res := topKFrequent(nums, k)

	fmt.Printf("%v", res)

}
