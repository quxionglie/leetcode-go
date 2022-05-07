package problem100

import (
	"regexp"
	"strings"
)

/**
125. 验证回文串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true
解释："amanaplanacanalpanama" 是回文串

示例 2:
输入: "race a car"
输出: false
解释："raceacar" 不是回文串
*/
func isPalindrome(s string) bool {
	reg, err := regexp.Compile("[0-9a-zA-Z]")
	if err != nil {
		return false
	}
	left := 0

	runeStr := []rune(s)
	right := len(runeStr) - 1

	for _, c := range s {
		leftStr := string(c)
		if !reg.MatchString(leftStr) {
			continue
		}

		rightStr := string(runeStr[right])
		for !reg.MatchString(rightStr) {
			right--
			rightStr = string(runeStr[right])
		}

		if left+right >= len(runeStr) || left >= right {
			break
		}
		if !strings.EqualFold(leftStr, rightStr) {
			return false
		}
		left++
		right--
	}

	return true
}
