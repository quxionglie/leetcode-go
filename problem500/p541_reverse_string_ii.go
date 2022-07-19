package problem500

/**
qxl 541. 反转字符串 II
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
示例 1：
输入：s = "abcdefg", k = 2
输出："bacdfeg"

示例 2：
输入：s = "abcd", k = 2
输出："bacd"

提示：
1 <= s.length <= 10^4
s 仅由小写英文组成
1 <= k <= 10^4
*/
func reverseStr(s string, k int) string {
	// 分成n组，每个2k个字符串
	n := len(s) / (2 * k)
	result := ""
	for i := 0; i < n; i++ {
		left := i * 2 * k
		right := (i + 1) * 2 * k
		str := s[left:right]
		result = result + reverse(str[0:k]) + str[k:]
	}

	// 不足2k长度
	if len(s)-n*2*k > 0 {
		str := s[n*2*k:]
		if len(str) > k { // k<len<2k
			result = result + reverse(str[0:k]) + str[k:]
		} else if len(str) <= k { // len<k
			result = result + reverse(str)
		}
	}

	return result
}

func reverse(s string) string {
	result := ""
	n := len(s)
	for i := 0; i < n; i++ {
		result = result + string(s[n-1-i])
	}
	return result
}
