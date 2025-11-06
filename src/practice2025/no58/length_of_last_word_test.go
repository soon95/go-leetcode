package no58

import (
	"fmt"
	"testing"
)

func Test_lengthOfLastWord(t *testing.T) {
	s := "   fly me   to   the moon  "

	res := lengthOfLastWord(s)

	fmt.Printf("%v", res)

}
