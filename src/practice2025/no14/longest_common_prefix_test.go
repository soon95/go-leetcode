package no14

import (
	"fmt"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	//strs := []string{"flower", "flow", "flight"}
	//strs := []string{"dog", "racecar", "car"}
	strs := []string{""}

	res := longestCommonPrefix(strs)

	fmt.Printf("%v", res)

}
