package main

// 这题的意思是：进位计算+1
// 思路到的查询后面每个值判断是否为9，如果是9置0，并继续向前查询，直至没有大于9的数为其加1

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}
	// 特殊情况，全是9的，扩以为置1
	arry := make([]int, len(digits)+1)
	copy(arry[1:], digits)
	arry[0] = 1
	return arry
}
