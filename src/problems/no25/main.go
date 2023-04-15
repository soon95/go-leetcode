package main

import (
	"go-leetcode/src/common"
)

type ListNode = common.ListNode

/*
25. K 个一组翻转链表
*https://leetcode.cn/problems/reverse-nodes-in-k-group/
*/
func reverseKGroup(head *ListNode, k int) *ListNode {

	if k == 1 {
		return head
	}

	count := 1

	end := head
	for end != nil && end.Next != nil && count < k {
		end = end.Next
		count++
	}

	if count < k {
		// 如果剩下元素不满k个，不需要反转
		return head
	}

	tail := head
	pre := head
	cur := head.Next

	for cur != end {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	tail.Next = reverseKGroup(end.Next, k)

	end.Next = pre

	return end
}

func main() {

	k := 4

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	ans := reverseKGroup(node1, k)

	print(ans)

}
