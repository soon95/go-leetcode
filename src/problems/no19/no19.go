package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type ListNode = common.ListNode

/*
*
19. 删除链表的倒数第 N 个结点
https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	length := countNthFromEnd(head, n)

	if length == n-1 {
		return head.Next
	}

	return head
}

func countNthFromEnd(head *ListNode, n int) int {

	var length int

	if head != nil && head.Next == nil {
		length = 0
	} else {
		length = countNthFromEnd(head.Next, n) + 1
	}

	if length == n {
		head.Next = head.Next.Next
	}

	return length
}

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	ans := removeNthFromEnd(node1, 5)

	fmt.Print(ans)

}
