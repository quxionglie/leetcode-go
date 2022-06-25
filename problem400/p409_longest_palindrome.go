package problem400

/**
qxl 409. 最长回文串
给定一个包含大写字母和小写字母的字符串s，返回通过这些字母构造成的 最长的回文串。

在构造过程中，请注意 区分大小写 。比如"Aa"不能当做一个回文字符串。

示例 1:
输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。

示例 2:
输入:s = "a"
输入:1

示例 3:
输入:s = "bb"
输入: 2
*/
func longestPalindrome(s string) int {
	m := map[int32]int{}
	for _, ch := range s {
		m[ch]++
	}
	n := 0
	var singleCh bool
	for _, v := range m {
		if v%2 == 0 {
			n = n + v
		} else {
			n = n + v/2*2
			singleCh = true
		}
	}
	if singleCh {
		n++
	}
	return n
}
