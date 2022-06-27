package problem400

/*
qxl 434. 字符串中的单词数

统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
请注意，你可以假定字符串里不包括任何不可打印的字符。
示例:
输入: "Hello, my name is John"
输出: 5
解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。
*/
func countSegments(s string) int {
	str := ""
	n := 0
	for i, _ := range s {
		if string(s[i]) == " " {
			if len(str) > 0 {
				n++
				str = ""
			}
		} else {
			str += string(s[i])
		}
		i++
	}
	if len(str) > 0 {
		n++
	}
	return n
}
