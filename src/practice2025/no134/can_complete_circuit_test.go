package no134

import (
	"fmt"
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	//gas := []int{1, 2, 3, 4, 5}
	//cost := []int{3, 4, 5, 1, 2}

	gas := []int{4, 5, 1, 2, 3}
	cost := []int{1, 2, 3, 4, 5}

	res := canCompleteCircuit(gas, cost)

	fmt.Printf("%v", res)
}
