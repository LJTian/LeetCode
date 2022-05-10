package main

import "fmt"

func longestPalindrome(s string) string {

	// 我读懂了别人的解法了
	// 此题需要画一个二维数组方式来解析
	// x,y 分别标识起始位置和终止位置

	// 回文字符串的特点：
	// 1-单个字符肯定是
	// 2-两个字符，它俩必须相等
	// 3-首位相等
	// >3-首位相等，去掉首位依然符合上面的条件

	if len(s) < 2 {
		return s
	}
	bit := make([][]bool, len(s))
	maxLen := 1
	startIndex := 0

	// 首先是单个字符的，对角线标识的是单个字符
	for i := 0; i < len(s); i++ {
		bit[i] = make([]bool, len(s))
		bit[i][i] = true
	}

	// 固定长度 最长也就是字符长度
	for sLen := 2; sLen < len(s)+1; sLen++ {

		// 移动起始位置
		for start := 0; start < len(s)-sLen+1; start++ {
			// 判断首位字符是否相等
			if s[start] != s[start+sLen-1] {
				continue
			} else if sLen == 2 {
				bit[start][start+sLen-1] = true
			} else {
				bit[start][start+sLen-1] = bit[start+1][start+sLen-2]
			}
			// 注意：这个必须要判断这个字符串是否是回文串，然后再判断是否是最大长度
			if bit[start][start+sLen-1] && sLen > maxLen {
				maxLen = sLen
				startIndex = start
			}
		}
	}

	return s[startIndex : startIndex+maxLen]

}

func main() {
	fmt.Println(longestPalindrome("bb"))
}
