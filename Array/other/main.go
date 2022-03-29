package main

import (
	"fmt"
	"sort"
)

// 从 K 个数 Array中找到最大两个数 循环或者递归
// 排序，取最后两个数
func findKthLargest(nums []int) (int, int) {
	sort.Ints(nums)
	return nums[len(nums)-1], nums[len(nums)-2]
}

// 测试
func main() {
	nums := []int{3, 2, 3, 7, 8, 4, 5, 5, 6}
	fmt.Println(fun1(nums))
}

// 循环或者递归，我理解题意想考一下排序算法(快排)
// 但是题中，没有说不让直接用 sort 包
// 理解错了，不需要排序，遍历一遍就够了，看fun1
// 递归，没有想明白怎么写

func swap(num1, num2 *int) bool {
	if *num2 > *num1 {
		*num1, *num2 = *num2, *num1
		return true
	}
	return false
}

// 循环
func fun1(nums []int) (int, int) {

	max1, max2 := nums[0], nums[1]
	swap(&max1, &max2)
	for i := 2; i < len(nums); i++ {
		if swap(&max2, &nums[i]) {
			swap(&max1, &max2)
		}
	}
	return max1, max2
}
