package problem700

import "strings"

/*
*
796. 旋转字符串
给定两个字符串, s和goal。如果在若干次旋转操作之后，s能变成goal，那么返回true。
s的 旋转操作 就是将s 最左边的字符移动到最右边。
例如, 若s = 'abcde'，在旋转一次之后结果就是'bcdea'。

示例 1:
输入: s = "abcde", goal = "cdeab"
输出: true

示例 2:
输入: s = "abcde", goal = "abced"
输出: false

提示:
1 <= s.length, goal.length <= 100
s和goal由小写英文字母组成
*/
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
