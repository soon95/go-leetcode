package no71

import (
	"fmt"
	"testing"
)

func Test_simplifyPath(t *testing.T) {
	//path := "/home/"
	path := "/../"

	res := simplifyPath(path)

	fmt.Printf("%v", res)

}
