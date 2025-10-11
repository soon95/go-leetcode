package no20

import (
	"fmt"
	"testing"
)

func Test_isValid(t *testing.T) {

	//str := "()[]{}"
	str := "(]"

	res := isValid(str)

	fmt.Printf("%v", res)

}
