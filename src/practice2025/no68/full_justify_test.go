package no68

import (
	"fmt"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	//words := []string{"This", "is", "an", "example", "of", "text", "justification."}
	words := []string{"shall", "be"}
	maxWidth := 16

	res := fullJustify(words, maxWidth)

	fmt.Printf("%v", res)
}
