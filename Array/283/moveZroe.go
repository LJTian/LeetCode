package main

// 移动0的位置
func moveZeroes1(nums []int) {
	// 遍历数组
	for i, j := 0, 0; i < len(nums); i++ {
		// 如果是非0值，将值移至到0的位置
		if nums[i] != 0 {
			nums[j] = nums[i]

			// 排除[1] 这种情况 --- 这个还是不太懂--(懂了：如果 i一直没有遇到0 则，i=j 没有进行交换，如果遇到0 则 j 为0 的位置，i 为非0位置，进行交换)
			if i != j {
				nums[i] = 0
			}
			// 记录0的位置
			j++
		}
	}
}

// 方法2
func moveZeroes2(nums []int) {

	// nums
	for i, j := 0, 0; i < len(nums); i++ {

		// 如果 i一直没有遇到0 则，i=j 没有进行交换，如果遇到0 则 j 为0 的位置，i 为非0位置，进行交换
		if nums[i] != 0 {
			// 交换方式
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}
