package no4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	totalLength := len(nums1) + len(nums2)

	k1 := (totalLength + 1) / 2
	k2 := (totalLength + 2) / 2

	return float64(getKth(nums1, nums2, k1)+getKth(nums1, nums2, k2)) * 0.5
}

func getKth(nums1, nums2 []int, k int) int {

	if k == 1 {
		if len(nums1) == 0 {
			return nums2[0]
		} else if len(nums2) == 0 {
			return nums1[0]
		} else {
			return min(nums1[0], nums2[0])
		}
	}

	if len(nums1) == 0 {
		return nums2[k-1]
	} else if len(nums2) == 0 {
		return nums1[k-1]
	}

	index1 := min(k/2, len(nums1)) - 1
	index2 := min(k/2, len(nums2)) - 1

	if nums1[index1] > nums2[index2] {
		return getKth(nums1, nums2[index2+1:], k-(index2+1))
	} else {
		return getKth(nums1[index1+1:], nums2, k-(index1+1))
	}
}
