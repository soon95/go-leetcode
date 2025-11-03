package no135

import (
	"fmt"
	"testing"
)

func Test_candy(t *testing.T) {
	//ratings := []int{1, 0, 2}
	//ratings := []int{1, 2, 2}
	ratings := []int{1, 3, 4, 5, 2}

	res := candy(ratings)

	fmt.Printf("%v", res)

}
