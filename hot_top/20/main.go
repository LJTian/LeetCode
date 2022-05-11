package main

func isValid(s string) bool {

	// 非偶数直接退出
	if len(s)%2 != 0 {
		return false
	}

	// 设置映射关系
	m1 := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	// 栈空间
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		// 尾括号
		if m1[s[i]] > 0 {
			// 检验出栈的是否与之匹配
			if len(stack) == 0 || stack[len(stack)-1] != m1[s[i]] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 首括号 入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func main() {

}
