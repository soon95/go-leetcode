package no167

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	numbers := []int{2, 7, 9, 11, 15}
	target := 18

	res := twoSum(numbers, target)

	fmt.Printf("%v", res)

}
