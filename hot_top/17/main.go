package main

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {

	/*
	   回溯算法
	   将所有的值先保存在map中
	   然后逐渐根据按键添加值，加到没有新的key就结束输入，记录结果。
	   回退本次按键，再此进行按键输入
	   // 难点对递归的应用
	*/

	// 如果输入按键数为0直接退出
	if len(digits) == 0 {
		return []string{}
	}

	// 回溯算法
	backtrack(digits, 0, "")
	return combinations
}

/*
   digits 按键值
   index 当前检索位置
   combination 计算出的结果
*/
func backtrack(digits string, index int, combination string) {

	// 判断是否到达最后，如果到达直接保存结果
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {

		// 获取当前数字键值
		digit := string(digits[index])

		// 获取数字键值对应的字符串
		letters := phoneMap[digit]

		// 遍历字符串每个值进行拼接结果
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
