package no763

import (
	"fmt"
	"testing"
)

func Test_partitionLabels(t *testing.T) {

	//s := "ababcc"
	s := "ababcbacadefegdehijhklij"

	res := partitionLabels(s)

	fmt.Printf("%v", res)

}
