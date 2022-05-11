package main

func singleNumber(nums []int) int {

	rest := nums[0]
	// 这题很简单主要是要知道下面这个特性：
	// 异或运算，出现相同的值会互相抵消
	for i := 1; i < len(nums); i++ {
		rest ^= nums[i]
	}

	return rest
}
