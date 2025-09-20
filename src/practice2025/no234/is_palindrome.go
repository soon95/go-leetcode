package no234

func isPalindrome(head *ListNode) bool {

	nums := make([]int, 0)

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	i, j := 0, len(nums)-1

	for i < j {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromeV2(head *ListNode) bool {

	lowNode, fastNode := head, head
	for fastNode.Next != nil && fastNode.Next.Next != nil {
		lowNode = lowNode.Next
		fastNode = fastNode.Next.Next
	}

	half := lowNode.Next

	newHead := reverseList(half)

	for head != nil && newHead != nil {

		if head.Val != newHead.Val {
			return false
		}

		head = head.Next
		newHead = newHead.Next
	}
	return true
}

func reverseList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	res := reverseList(head.Next)

	head.Next.Next = head
	head.Next = nil

	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}
