package no28

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	haystack := "sadbutdad"
	needle := "dad"

	res := strStr(haystack, needle)

	fmt.Printf("%v", res)

}
