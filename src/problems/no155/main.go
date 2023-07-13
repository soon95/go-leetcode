package main

import "fmt"

/*
*155. 最小栈
https://leetcode.cn/problems/min-stack/

看了题解，做复杂了，只需要维护一个数组，维护当前栈顶元素和当前的最小元素的元组，就可以了
*/
type MinStack struct {
	Nodes   []*StackNode
	MinNode *StackNode
}

type StackNode struct {
	Val      int
	Previous *StackNode
	Next     *StackNode
}

func Constructor() MinStack {

	return MinStack{
		Nodes: []*StackNode{},
	}
}

func (this *MinStack) Push(val int) {

	node := &StackNode{
		Val: val,
	}

	this.Nodes = append([]*StackNode{node}, this.Nodes...)

	this.Insert(this.MinNode, node)

}

func (this *MinStack) Insert(head, target *StackNode) {

	if head == nil {
		this.MinNode = target
		return
	}

	var preNode *StackNode
	curNode := head

	for curNode != nil {

		if curNode.Val > target.Val {
			break
		}

		preNode = curNode
		curNode = curNode.Next
	}

	if preNode == nil {
		target.Next = head
		head.Previous = target
		head = target
	} else {

		next := preNode.Next

		preNode.Next = target
		target.Previous = preNode

		if next == nil {
			target.Next = nil
		} else {
			target.Next = next
			next.Previous = target
		}

	}
	this.MinNode = head
}

func (this *MinStack) Pop() {

	target := this.Nodes[0]

	this.Nodes = this.Nodes[1:]

	this.Remove(target)
}

func (this *MinStack) Remove(target *StackNode) {

	previous := target.Previous
	next := target.Next

	if previous == nil {

		if next != nil {
			next.Previous = nil
		}

		this.MinNode = next

	} else if next == nil {
		previous.Next = nil
	} else {
		previous.Next = next
		next.Previous = previous
	}

}

func (this *MinStack) Top() int {

	return this.Nodes[0].Val
}

func (this *MinStack) GetMin() int {

	return this.MinNode.Val

}

func main() {

	//minStack := Constructor()
	//minStack.Push(-2)
	//minStack.Push(0)
	//minStack.Push(-3)
	//fmt.Printf("%v\n", minStack.GetMin())
	//minStack.Pop()
	//fmt.Printf("%v\n", minStack.Top())
	//fmt.Printf("%v\n", minStack.GetMin())

	minStack := Constructor()
	minStack.Push(395)
	fmt.Printf("%v\n", minStack.GetMin())
	fmt.Printf("%v\n", minStack.Top())
	fmt.Printf("%v\n", minStack.GetMin())
	minStack.Push(276)
	minStack.Push(29)
	fmt.Printf("%v\n", minStack.GetMin())
	minStack.Push(-482)
	fmt.Printf("%v\n", minStack.GetMin())
	minStack.Pop()
	minStack.Push(-108)
	minStack.Push(-251)
	fmt.Printf("%v\n", minStack.GetMin())
	fmt.Printf("%v\n", minStack.Top())
	minStack.Push(-439)
	fmt.Printf("%v\n", minStack.Top())
	minStack.Push(370)
	minStack.Pop()
	minStack.Pop()
	fmt.Printf("%v\n", minStack.GetMin())
	minStack.Pop()
	fmt.Printf("%v\n", minStack.GetMin())

}
