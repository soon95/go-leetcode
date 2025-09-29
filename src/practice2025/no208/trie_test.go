package no208

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {

	trie := Constructor()

	trie.Insert("apple")
	fmt.Printf("%v\n", trie.Search("apple"))
	fmt.Printf("%v\n", trie.Search("app"))
	fmt.Printf("%v\n", trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Printf("%v\n", trie.Search("app"))
}
