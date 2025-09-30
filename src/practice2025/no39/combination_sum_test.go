package no39

import (
	"fmt"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7

	res := combinationSum(candidates, target)

	fmt.Printf("%v", res)
}
