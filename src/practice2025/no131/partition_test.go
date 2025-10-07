package no131

import (
	"fmt"
	"testing"
)

func Test_partition(t *testing.T) {
	s := "aab"

	res := partition(s)

	fmt.Printf("%v", res)
}
