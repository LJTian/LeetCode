package main

// 1:  当前位置 + 3 % len(nums),就是新的位置
func rotate1(nums []int, k int) {

	lenth := len(nums)
	newNums := make([]int, lenth)

	for k1, v := range nums {
		newNums[(k1+k)%lenth] = v
	}

	//for k, v := range newNums {
	//	nums[k] = v
	//}

	copy(nums, newNums)
}

// 2:O(1) 空间

// 翻转数组
func reversal(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[i]
	}
}

// 翻转数组
func rotate2(nums []int, k int) {
	realK := k % len(nums)
	reversal(nums)
	reversal(nums[:realK])
	reversal(nums[realK:])
}
