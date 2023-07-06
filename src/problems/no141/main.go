package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type ListNode = common.ListNode

/*
*141. 环形链表
https://leetcode.cn/problems/linked-list-cycle/
*/
func hasCycle(head *ListNode) bool {

	dic := make(map[*ListNode]bool)

	cur := head

	for cur != nil {

		_, exist := dic[cur]

		if exist {
			return true
		}

		dic[cur] = true

		cur = cur.Next
	}

	return false
}

func main() {

	node1 := &ListNode{
		Val: 3,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 0,
	}
	node4 := &ListNode{
		Val: 4,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	//node4.Next = node2

	fmt.Printf("%v", hasCycle(node1))

}
