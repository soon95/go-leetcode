package no739

import (
	"fmt"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}

	res := dailyTemperatures(temperatures)

	fmt.Printf("%v", res)

}
