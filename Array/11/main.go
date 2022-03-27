package main

// 方法1
func maxArea1(height []int) int {

	// 保存最大值
	max := 0

	// 两边同时移动
	for i, j := 0, len(height)-1; i < j; {

		// 水容积
		var area int

		// 判断哪边水位低，就移动那边
		if height[i] > height[j] {
			area = (j - i) * height[j]
			j--
		} else {
			area = (j - i) * height[i]
			i++
		}
		if area > max {
			max = area
		}
	}
	return max
}
