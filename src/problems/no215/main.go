package main

import (
	"fmt"
)

/*
*215. 数组中的第K个最大元素
https://leetcode.cn/problems/kth-largest-element-in-an-array/
维护最大堆
*/
func findKthLargest(nums []int, k int) int {

	length := len(nums)

	buildMaxHeap(nums)

	for i := 1; i < k; i++ {

		exchange(nums, 0, length-1)

		length--

		adjustMaxHeap(nums, 0, length)
	}

	return nums[0]
}

func buildMaxHeap(nums []int) {

	length := len(nums)

	curIndex := length/2 - 1

	for curIndex >= 0 {

		adjustMaxHeap(nums, curIndex, length)

		curIndex--
	}
}

func adjustMaxHeap(nums []int, head, length int) {

	for head <= length/2-1 {

		leftIndex := head*2 + 1
		rightIndex := head*2 + 2

		if rightIndex < length && nums[rightIndex] > nums[leftIndex] {

			if nums[head] < nums[rightIndex] {

				exchange(nums, head, rightIndex)

				head = rightIndex
			} else {
				break
			}

		} else {

			if nums[head] < nums[leftIndex] {
				exchange(nums, head, leftIndex)

				head = leftIndex
			} else {
				break
			}
		}
	}
}

func exchange(nums []int, index1, index2 int) {

	temp := nums[index1]
	nums[index1] = nums[index2]
	nums[index2] = temp
}

func main() {

	//nums := []int{3, 2, 1, 5, 6, 4}
	//k := 2

	//nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	//k := 4

	nums := []int{7, 6, 5, 4, 3, 2, 1}
	k := 5

	fmt.Printf("%v", findKthLargest(nums, k))

}
