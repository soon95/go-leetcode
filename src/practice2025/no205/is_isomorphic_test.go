package no205

import (
	"fmt"
	"testing"
)

func Test_isIsomorphic(test *testing.T) {
	//s := "egg"
	//t := "add"
	s := "foo"
	t := "bar"

	res := isIsomorphic(s, t)

	fmt.Printf("%v", res)

}
