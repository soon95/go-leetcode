package no437

import (
	"fmt"
	"testing"
)

func Test_pathSum(t *testing.T) {

	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Right: &TreeNode{
				Val: 11,
			},
		},
	}

	res := pathSum(root, 8)

	fmt.Printf("%v", res)

}
