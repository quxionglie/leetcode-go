package problem500

/**
504. 七进制数
给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
示例 1:
输入: num = 100
输出: "202"

示例 2:
输入: num = -7
输出: "-10"

提示：
-10^7 <= num <= 10^7
*/
func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	negative := num < 0
	if negative {
		num = -num
	}
	s := []byte{}
	for num > 0 {
		s = append(s, '0'+byte(num%7))
		num /= 7
	}
	if negative {
		s = append(s, '-')
	}
	for i, n := 0, len(s); i < n/2; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
	return string(s)
}
