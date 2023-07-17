package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type ListNode = common.ListNode

/*
*160. 相交链表
https://leetcode.cn/problems/intersection-of-two-linked-lists/
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	dic := map[*ListNode]bool{}

	cur := headA

	for cur != nil {
		dic[cur] = true
		cur = cur.Next
	}

	cur = headB

	for cur != nil {

		_, exist := dic[cur]

		if exist {
			return cur
		}

		cur = cur.Next
	}

	return nil
}

func main() {

	nodeA1 := &ListNode{
		Val: 1,
	}
	nodeA2 := &ListNode{
		Val: 2,
	}

	nodeB1 := &ListNode{
		Val: 11,
	}
	nodeB2 := &ListNode{
		Val: 12,
	}
	nodeB3 := &ListNode{
		Val: 13,
	}

	nodeC1 := &ListNode{
		Val: 21,
	}
	nodeC2 := &ListNode{
		Val: 22,
	}
	nodeC3 := &ListNode{
		Val: 23,
	}

	nodeA1.Next = nodeA2
	nodeA2.Next = nodeC1

	nodeB1.Next = nodeB2
	nodeB2.Next = nodeB3
	nodeB3.Next = nodeC1

	nodeC1.Next = nodeC2
	nodeC2.Next = nodeC3

	node := getIntersectionNode(nodeA1, nodeB1)

	fmt.Printf("%v", node)

}
