package problem300

/**
392. 判断子序列
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。

字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
进阶：

如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？
示例 1：
输入：s = "abc", t = "ahbgdc"
输出：true

示例 2：
输入：s = "axc", t = "ahbgdc"
输出：false

提示：

0 <= s.length <= 100
0 <= t.length <= 10^4
两个字符串都只由小写字符组成。
*/
func isSubsequence(s string, t string) bool {
	//双指针
	//这样，我们初始化两个指针 ii 和 jj，分别指向 ss 和 tt 的初始位置。每次贪心地匹配，匹配成功则 ii 和 jj 同时右移，
	//匹配 ss 的下一个位置，匹配失败则 jj 右移，ii 不变，尝试用 tt 的下一个字符匹配 ss。
	//最终如果 ii 移动到 ss 的末尾，就说明 ss 是 tt 的子序列。
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}
