package no42

import (
	"fmt"
	"testing"
)

func Test_trap(t *testing.T) {

	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height := []int{4, 2, 0, 3, 2, 5}

	res := trap(height)

	fmt.Printf("%v", res)

}
