package no173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	index int
	list  []int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		index: 0,
		list:  parseToList(root),
	}
}

func parseToList(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	leftList := parseToList(root.Left)
	rightList := parseToList(root.Right)

	leftList = append(leftList, root.Val)

	return merge(leftList, rightList)
}

func merge(list1, list2 []int) []int {

	list := make([]int, 0)

	index1, index2 := 0, 0

	for index1 < len(list1) && index2 < len(list2) {
		if list1[index1] <= list2[index2] {
			list = append(list, list1[index1])
			index1++
		} else {
			list = append(list, list2[index2])
			index2++
		}
	}

	for index1 < len(list1) {
		list = append(list, list1[index1])
		index1++
	}
	for index2 < len(list2) {
		list = append(list, list2[index2])
		index2++
	}

	return list
}

func (this *BSTIterator) Next() int {
	num := this.list[this.index]
	this.index++
	return num
}

func (this *BSTIterator) HasNext() bool {
	return this.index < len(this.list)
}
