package problem200

/**
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若s 和 t中每个字符出现的次数都相同，则称s 和 t互为字母异位词。
示例1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false
*/
func isAnagram(s string, t string) bool {
	m := make(map[uint8]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		m[s[i]] = m[s[i]] + 1
		m[t[i]] = m[t[i]] - 1
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
