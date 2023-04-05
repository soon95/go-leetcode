package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/*
*
两数相加
https://leetcode.cn/problems/add-two-numbers/
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return doAddTwoNumbers(l1, l2, 0)
}

func doAddTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {

	if l1 == nil && l2 == nil {
		if carry != 0 {
			return &ListNode{Val: carry}
		} else {
			return nil
		}
	}

	var num1 int
	var l1Next *ListNode
	if l1 == nil {
		num1 = 0
		l1Next = nil
	} else {
		num1 = l1.Val
		l1Next = l1.Next
	}

	var num2 int
	var l2Next *ListNode
	if l2 == nil {
		num2 = 0
		l2Next = nil
	} else {
		num2 = l2.Val
		l2Next = l2.Next
	}

	num := num1 + num2 + carry

	curVal := num % 10
	curCarry := num / 10

	cur := &ListNode{Val: curVal}

	cur.Next = doAddTwoNumbers(l1Next, l2Next, curCarry)

	return cur
}

func main() {

	l1 := &ListNode{
		Val: 9,
	}

	l1node2 := &ListNode{
		Val: 9,
	}

	l1node3 := &ListNode{
		Val: 9,
	}
	l1node4 := &ListNode{
		Val: 9,
	}
	l1node5 := &ListNode{
		Val: 9,
	}
	l1node6 := &ListNode{
		Val: 9,
	}
	l1node7 := &ListNode{
		Val: 9,
	}

	l1.Next = l1node2
	l1node2.Next = l1node3
	l1node3.Next = l1node4
	l1node4.Next = l1node5
	l1node5.Next = l1node6
	l1node6.Next = l1node7

	l2 := &ListNode{
		Val: 9,
	}

	l2node2 := &ListNode{
		Val: 9,
	}

	l2node3 := &ListNode{
		Val: 9,
	}
	l2node4 := &ListNode{
		Val: 9,
	}

	l2.Next = l2node2
	l2node2.Next = l2node3
	l2node3.Next = l2node4

	numbers := addTwoNumbers(l1, l2)

	fmt.Print(numbers)

}
