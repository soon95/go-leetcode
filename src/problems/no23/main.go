package main

import (
	"go-leetcode/src/common"
	"math"
)

type ListNode = common.ListNode

/*
*
23. 合并 K 个升序链表
https://leetcode.cn/problems/merge-k-sorted-lists/
*/
func mergeKLists(lists []*ListNode) *ListNode {

	min := math.MaxInt
	index := -1

	for i, list := range lists {

		if list != nil && min > list.Val {
			index = i
			min = list.Val
		}
	}

	if index == -1 {
		return nil
	} else {

		cur := lists[index]

		lists[index] = cur.Next

		cur.Next = mergeKLists(lists)

		return cur
	}
}

func main() {

	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3

	node4 := &ListNode{Val: 1}
	node5 := &ListNode{Val: 3}
	node6 := &ListNode{Val: 4}
	node4.Next = node5
	node5.Next = node6

	node7 := &ListNode{Val: 2}
	node8 := &ListNode{Val: 6}
	node7.Next = node8

	lists := []*ListNode{node1, node4, node7}

	ans := mergeKLists(lists)

	print(ans)

}
