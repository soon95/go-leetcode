package no274

import (
	"fmt"
	"testing"
)

func Test_hIndex(t *testing.T) {
	//citations := []int{3, 0, 6, 1, 5}
	citations := []int{1, 3, 1}

	res := hIndex(citations)

	fmt.Printf("%v", res)
}
