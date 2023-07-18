package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type ListNode = common.ListNode

/*
206. 反转链表
*https://leetcode.cn/problems/reverse-linked-list/

迭代和递归两种方法
*/
func reverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	next := head.Next
	head.Next = nil

	ans := reverseList(next)

	next.Next = head

	return ans
}

func reverseListV2(head *ListNode) *ListNode {

	cur := head
	var pre *ListNode
	for cur != nil {

		next := cur.Next

		cur.Next = pre

		pre = cur
		cur = next
	}

	return pre
}

func main() {

	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 3,
	}
	node4 := &ListNode{
		Val: 4,
	}
	node5 := &ListNode{
		Val: 5,
	}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	ans := reverseListV2(node1)

	fmt.Printf("%v", ans)

}
