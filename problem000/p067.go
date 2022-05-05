package problem000

import "strconv"

/**
67. 二进制求和
给你两个二进制字符串，返回它们的和（用二进制表示）。
输入为 非空 字符串且只包含数字1和0。
示例1:
输入: a = "11", b = "1"
输出: "100"

示例2:
输入: a = "1010", b = "1011"
输出: "10101"
链接：https://leetcode-cn.com/problems/add-binary
*/
func addBinary(a string, b string) string {
	aLen, bLen := len(a), len(b)
	ans := ""
	carry := 0

	max := bLen
	if aLen > bLen {
		max = aLen
	}

	for i := 0; i < max; i++ {
		if i < aLen {
			carry += int(a[aLen-i-1] - '0')
		}
		if i < bLen {
			carry += int(b[bLen-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans //当前位值
		carry /= 2                        //下一进位值
	}

	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
