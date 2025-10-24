package no72

import (
	"fmt"
	"testing"
)

func Test_minDistance(t *testing.T) {
	//word1 := "horse"
	//word2 := "ros"
	word1 := "a"
	word2 := "ab"

	res := minDistance(word1, word2)

	fmt.Printf("%v", res)

}
