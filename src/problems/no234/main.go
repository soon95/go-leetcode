package main

import "go-leetcode/src/common"

type ListNode = common.ListNode

/*
*234. 回文链表
https://leetcode.cn/problems/palindrome-linked-list/
*/
func isPalindrome(head *ListNode) bool {

	nums := []int{}

	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}

	left, right := 0, len(nums)-1

	for left < right {

		if nums[left] != nums[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {

	node1 := &ListNode{
		Val: 1,
	}
	node2 := &ListNode{
		Val: 2,
	}
	node3 := &ListNode{
		Val: 2,
	}
	node4 := &ListNode{
		Val: 1,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4

	println(isPalindrome(node1))

}
