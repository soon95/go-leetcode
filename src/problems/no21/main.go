package main

import "go-leetcode/src/common"

type ListNode = common.ListNode

/*
21. 合并两个有序链表
*https://leetcode.cn/problems/merge-two-sorted-lists/
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}
	node4.Next = node5
	node5.Next = node6

	ans := mergeTwoLists(node1, node4)

	print(ans)

}
