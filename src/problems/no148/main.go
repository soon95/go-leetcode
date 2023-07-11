package main

import (
	"fmt"
	"go-leetcode/src/common"
)

type ListNode = common.ListNode

/*
*
148. 排序链表
https://leetcode.cn/problems/sort-list/
因为题目中需要算法复杂度为 nlogn 因此考虑使用归并排序
*/
func sortList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	} else if head.Next.Next == nil {

		node1 := head
		node2 := head.Next

		if node1.Val < node2.Val {
			node1.Next = node2
			node2.Next = nil
			return node1
		} else {
			node2.Next = node1
			node1.Next = nil
			return node2
		}
	}

	middle := findMiddle(head)

	nextHead := middle.Next

	middle.Next = nil

	head1 := sortList(head)
	head2 := sortList(nextHead)

	return mergeNode(head1, head2)
}

func mergeNode(head1, head2 *ListNode) *ListNode {

	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var head *ListNode
	if head1.Val <= head2.Val {
		head = head1
		head.Next = mergeNode(head1.Next, head2)
	} else {
		head = head2
		head.Next = mergeNode(head1, head2.Next)
	}

	return head
}

func findMiddle(head *ListNode) *ListNode {

	slowNode, fastNode := head, head.Next

	for fastNode != nil {

		fastNode = fastNode.Next

		if fastNode == nil {
			break
		} else {
			fastNode = fastNode.Next
		}

		slowNode = slowNode.Next

	}

	return slowNode
}

func main() {

	//node1 := &ListNode{
	//	Val: 4,
	//}
	//node2 := &ListNode{
	//	Val: 2,
	//}
	//node3 := &ListNode{
	//	Val: 1,
	//}
	//node4 := &ListNode{
	//	Val: 3,
	//}
	//node1.Next = node2
	//node2.Next = node3
	//node3.Next = node4

	node1 := &ListNode{
		Val: -1,
	}
	node2 := &ListNode{
		Val: 5,
	}
	node3 := &ListNode{
		Val: 3,
	}
	node4 := &ListNode{
		Val: 4,
	}
	node5 := &ListNode{
		Val: 0,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	ans := sortList(node1)

	fmt.Printf("%v", ans)

}
