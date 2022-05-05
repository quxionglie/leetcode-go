package problem000

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//链接：https://leetcode-cn.com/problems/valid-parentheses
func isValid(s string) bool {
	size := len(s)
	if size%2 == 1 {
		return false
	}

	m := map[byte]byte{
		']': '[',
		'}': '{',
		')': '('}

	stack := []byte{}

	for i := 0; i < size; i++ {
		if m[s[i]] > 0 {
			// 如果栈为空或栈顶元素不相等，返回false
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
