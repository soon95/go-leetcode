package main

import "go-leetcode/src/common"

type ListNode = common.ListNode

/*
24. 两两交换链表中的节点
*https://leetcode.cn/problems/swap-nodes-in-pairs/
*/
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next

	head.Next = swapPairs(next.Next)

	next.Next = head

	return next
}

func main() {

	node1 := &ListNode{Val: 1}
	//node2 := &ListNode{Val: 2}
	//node3 := &ListNode{Val: 3}
	//node4 := &ListNode{Val: 4}
	//node1.Next = node2
	//node2.Next = node3
	//node3.Next = node4

	ans := swapPairs(node1)

	print(ans)

}
