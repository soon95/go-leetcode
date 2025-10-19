package no118

import (
	"fmt"
	"testing"
)

func Test_generate(t *testing.T) {
	numRows := 5

	res := generate(numRows)

	fmt.Printf("%v", res)
}
