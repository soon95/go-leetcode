package no88

func merge(nums1 []int, m int, nums2 []int, n int) {

	tmp := make([]int, m)
	for i := 0; i < m; i++ {
		tmp[i] = nums1[i]
	}

	index, index1, index2 := 0, 0, 0

	for index1 < m && index2 < n {
		if tmp[index1] <= nums2[index2] {
			nums1[index] = tmp[index1]
			index1++
		} else {
			nums1[index] = nums2[index2]
			index2++
		}

		index++
	}

	for index1 < m {
		nums1[index] = tmp[index1]
		index1++
		index++
	}
	for index2 < n {
		nums1[index] = nums2[index2]
		index2++
		index++
	}

	return
}
