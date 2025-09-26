package no114

import (
	"fmt"
	"testing"
)

func Test_doFlatten(t *testing.T) {

	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val: 3,
	//		},
	//		Right: &TreeNode{
	//			Val: 4,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val: 5,
	//		Right: &TreeNode{
	//			Val: 6,
	//		},
	//	},
	//}

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	flatten(root)

	fmt.Printf("%v", root)

}
