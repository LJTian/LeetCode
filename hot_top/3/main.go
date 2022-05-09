package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	map1 := make(map[byte]int)
	maxLen, Len, index := 0, 0, 0

	for i := 0; i < len(s); i++ {
		// 判断是否有重复值
		if v, ok := map1[s[i]]; ok {

			//  没想明白（也许是想把位移之前的数据不再进行统计，后面出现了重复的，之前的字段就没必要再计算了）
			if v+1 > index { // abba
				index = v + 1
			}
		}

		Len = i - index + 1
		if Len > maxLen {
			maxLen = Len
		}
		map1[s[i]] = i
	}

	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}
