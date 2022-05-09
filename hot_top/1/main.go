package main

import "fmt"

func twoSum(nums []int, target int) []int {
	map1 := make(map[int]int)

	// 记住 a + b = c 反之 b = c - a
	for k, v := range nums {
		if value, ok := map1[target-v]; ok {
			return []int{value, k}
		}
		map1[v] = k
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
