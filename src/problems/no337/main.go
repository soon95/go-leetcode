package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type TreeNode = common.TreeNode

/*
*337. 打家劫舍 III
https://leetcode.cn/problems/house-robber-iii/
*/
func rob(root *TreeNode) int {

	moneyWithCur, moneyWithoutCur := doRob(root)

	return max(moneyWithCur, moneyWithoutCur)
}

func doRob(root *TreeNode) (int, int) {

	if root == nil {
		return 0, 0
	}

	moneyWithLeft, moneyWithoutLeft := doRob(root.Left)
	moneyWithRight, moneyWithoutRight := doRob(root.Right)

	moneyWithCur := moneyWithoutLeft + moneyWithoutRight + root.Val
	moneyWithoutCur := max(moneyWithLeft, moneyWithoutLeft) + max(moneyWithRight, moneyWithoutRight)

	return moneyWithCur, moneyWithoutCur
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func main() {

	//node1 := &TreeNode{
	//	Val: 3,
	//}
	//node2 := &TreeNode{
	//	Val: 2,
	//}
	//node3 := &TreeNode{
	//	Val: 3,
	//}
	//node4 := &TreeNode{
	//	Val: 3,
	//}
	//node5 := &TreeNode{
	//	Val: 1,
	//}
	//node1.Left = node2
	//node1.Right = node3
	//node2.Right = node4
	//node3.Right = node5

	node1 := &TreeNode{
		Val: 3,
	}
	node2 := &TreeNode{
		Val: 4,
	}
	node3 := &TreeNode{
		Val: 5,
	}
	node4 := &TreeNode{
		Val: 1,
	}
	node5 := &TreeNode{
		Val: 3,
	}
	node6 := &TreeNode{
		Val: 1,
	}
	node1.Left = node2
	node1.Right = node3
	node2.Left = node4
	node2.Right = node5
	node3.Right = node6

	fmt.Printf("%v", rob(node1))

}
