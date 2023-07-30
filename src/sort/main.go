package main

import "fmt"

/*
*
冒泡排序
*/
func bubbleSort(nums []int) []int {

	length := len(nums)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
			}
		}
	}

	return nums
}

/*
*
归并排序
*/
func mergeSort(nums []int) []int {

	length := len(nums)

	if length < 2 {
		return nums
	}

	middle := length / 2

	leftNums := mergeSort(nums[:middle])
	rightNums := mergeSort(nums[middle:])

	leftIndex, rightIndex := 0, 0

	ans := []int{}

	for leftIndex < len(leftNums) && rightIndex < len(rightNums) {
		if leftNums[leftIndex] < rightNums[rightIndex] {
			ans = append(ans, leftNums[leftIndex])
			leftIndex++
		} else {
			ans = append(ans, rightNums[rightIndex])
			rightIndex++
		}
	}

	if leftIndex == len(leftNums) {
		for rightIndex < len(rightNums) {
			ans = append(ans, rightNums[rightIndex])
			rightIndex++
		}
	} else if rightIndex == len(rightNums) {
		for leftIndex < len(leftNums) {
			ans = append(ans, leftNums[leftIndex])
			leftIndex++
		}
	}

	return ans
}

/*
*
快速排序
*/
func quickSort(nums []int) []int {

	doQuickSort(nums, 0, len(nums)-1)

	return nums
}

func doQuickSort(nums []int, start, end int) {

	if start >= end {
		return
	}

	middle := partition(nums, start, end)
	doQuickSort(nums, start, middle-1)
	doQuickSort(nums, middle+1, end)
}

func partition(nums []int, start, end int) int {

	if start > end {
		return 0
	} else if start == end {
		return start
	}

	middle := nums[start]

	left, right := start+1, end

	for true {

		for left <= right && nums[left] <= middle {
			left++
		}

		for left <= right && nums[right] >= middle {
			right--
		}

		if left >= right {
			break
		}
		swap(nums, left, right)

	}

	swap(nums, start, right)

	return right
}

/*
*
交换两个元素
*/
func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {

	nums := []int{7, 2, 5, 6, 8, 3, 1, 4}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//nums := []int{8, 7, 6, 5, 4, 3, 2, 1}
	//nums := []int{3, 1, 2}

	fmt.Printf("%v\n", bubbleSort(nums))
	//fmt.Printf("%v", partition(nums, 0, 7))
	fmt.Printf("%v\n", quickSort(nums))
	fmt.Printf("%v\n", mergeSort(nums))

}
