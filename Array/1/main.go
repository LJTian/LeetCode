package main

// 1:暴力
// 2:hash 表 target-v

func twoSum(nums []int, target int) []int {

	map1 := make(map[int]int, 0)
	for k1, v := range nums {
		if k2, ok := map1[target-v]; ok {
			return []int{k1, k2}
		}
		map1[v] = k1
	}
	return nil
}
