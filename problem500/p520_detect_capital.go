package problem500

import "strings"

/**
qxl 520. 检测大写字母
我们定义，在以下情况时，单词的大写用法是正确的：

全部字母都是大写，比如 "USA" 。
单词中所有字母都不是大写，比如 "leetcode" 。
如果单词不只含有一个字母，只有首字母大写，比如"Google" 。
给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。
示例 1：
输入：word = "USA"
输出：true

示例 2：
输入：word = "FlaG"
输出：false

提示：
1 <= word.length <= 100
word 由小写和大写英文字母组成
*/
func detectCapitalUse(word string) bool {
	if len(word) == 0 {
		return true
	}
	c := string(word[0])
	if c == strings.ToUpper(c) {
		if word == strings.ToUpper(word) {
			//全部字母都是大写，比如 "USA" 。
			return true
		}
		//如果单词不只含有一个字母，只有首字母大写， 比如 "Google" 。
		tmp := word[1:len(word)]
		return tmp == strings.ToLower(tmp)
	} else {
		//单词中所有字母都不是大写，比如 "leetcode" 。
		return word == strings.ToLower(word)
	}

}
