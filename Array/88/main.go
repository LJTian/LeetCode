package main

// 1、解题思路，倒叙
func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}

	x, y := m-1, n-1
	for i := len(nums1) - 1; i > 0; i-- {
		if nums1[x] > nums2[y] {
			nums1[i] = nums1[x]
			x--
		} else {
			nums1[i] = nums2[y]
			y--
		}
		if x == -1 {
			copy(nums1, nums2[:y+1])
			break
		}
		if y == -1 {
			break
		}
	}
}
