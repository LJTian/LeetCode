package main

func climbStairs(n int) int {

	// 注意：这题的思路，主要为
	// 上第三个台阶时，只有两种情况：
	// 1、从第二台阶上面迈一步上来
	// 2、从第一台阶上面迈两步上来
	// 故：上第三个台阶 一共有 f(1) + f(2)
	// 上第四台阶 一共有 f(2) + f(3)
	// 数学归纳一下：
	// 第 N 台阶 f(n) = f(n-1) + f(n-2)

	var num = [3]int{1, 2, 3}

	// 小于3 直接使用默认值即可
	if n < 4 {
		return num[n-1]
	}

	// 大于3，需要依次计算最新值，并保存后三位的值
	for i := 3; i < n; i++ {
		num[0] = num[1]
		num[1] = num[2]
		num[2] = num[0] + num[1] // 计算最终结果
	}
	return num[2]
}
