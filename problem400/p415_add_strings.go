package problem400

import "strconv"

/*
qxl 415. 字符串相加
给定两个字符串形式的非负整数num1 和num2，计算它们的和并同样以字符串形式返回。
你不能使用任何內建的用于处理大整数的库（比如 BigInteger），也不能直接将输入的字符串转换为整数形式。

示例 1：
输入：num1 = "11", num2 = "123"
输出："134"

示例 2：
输入：num1 = "456", num2 = "77"
输出："533"

示例 3：
输入：num1 = "0", num2 = "0"
输出："0"

提示：
1 <= num1.length, num2.length <= 104
num1 和num2 都只包含数字0-9
num1 和num2 都不包含任何前导零
*/
func addStrings(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	ans := ""
	carry := 0

	max := n2
	if n1 > n2 {
		max = n1
	}

	for i := 0; i < max; i++ {
		if i < n1 {
			carry += int(num1[n1-i-1] - '0')
		}
		if i < n2 {
			carry += int(num2[n2-i-1] - '0')
		}
		ans = strconv.Itoa(carry%10) + ans //当前位值
		carry /= 10                        //下一进位值
	}

	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
