package no106

import (
	"fmt"
	"testing"
)

func Test_buildTree(t *testing.T) {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	res := buildTree(inorder, postorder)

	fmt.Printf("%v", res)

}
