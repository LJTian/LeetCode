package main

func maxProfit(prices []int) int {

	// 找出低地点，与最大差值
	// 因为是正向遍历的，所以只要找到最小值对应的最大差值就行，比如，今天是一个3元，明天是2元，后天的4元，一定是2元入 4元出值最大
	// 如果遇到了后面更低的值，但是没有特别大的值，这样结果依然会保存之前的最大差值。
	// 在遍历完成之后，虽然不能保证min是买入值，但是可以保证max可定是历史差值中最大的
	min, max := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		if prices[i]-min > max {
			max = prices[i] - min
		}
	}
	return max
}
