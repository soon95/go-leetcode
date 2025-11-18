package no383

import (
	"fmt"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	ransomNote := "a"
	magazine := "ba"

	res := canConstruct(ransomNote, magazine)

	fmt.Printf("%v", res)

}
