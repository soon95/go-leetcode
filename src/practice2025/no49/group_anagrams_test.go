package no49

import (
	"fmt"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//strs := []string{""}

	res := groupAnagrams(strs)

	fmt.Printf("%v", res)

}
