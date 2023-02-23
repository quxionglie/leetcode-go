package problem600

import "strconv"

/*
*
670. 最大交换
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :
输入: 2736
输出: 7236
解释: 交换数字2和数字7。

示例 2 :
输入: 9973
输出: 9973
解释: 不需要交换。
注意:
给定数字的范围是[0, 10^8]
*/
func maximumSwap(num int) int {
	ans := num
	s := []byte(strconv.Itoa(num))
	for i := range s {
		for j := range s[:i] {
			s[i], s[j] = s[j], s[i]
			v, _ := strconv.Atoi(string(s))
			ans = max(ans, v)
			s[i], s[j] = s[j], s[i]
		}
	}
	return ans
}
