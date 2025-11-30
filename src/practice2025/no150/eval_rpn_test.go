package no150

import (
	"fmt"
	"testing"
)

func Test_evalRPN(t *testing.T) {
	//tokens := []string{"2", "1", "+", "3", "*"}
	//tokens := []string{"4", "13", "5", "/", "+"}
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}

	res := evalRPN(tokens)

	fmt.Printf("%v", res)

}
