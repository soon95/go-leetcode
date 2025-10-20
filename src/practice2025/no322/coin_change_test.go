package no322

import (
	"fmt"
	"testing"
)

func Test_coinChange(t *testing.T) {
	//coins := []int{1, 2, 5}
	//amount := 11

	coins := []int{1}
	amount := 0

	res := coinChange(coins, amount)

	fmt.Printf("%v", res)

}
