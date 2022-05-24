package problem100

/**
168. Excel表列名称

给你一个整数 columnNumber ，返回它在 Excel 表中相对应的列名称。
例如：
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...

示例 1：
输入：columnNumber = 1
输出："A"

示例 2：
输入：columnNumber = 28
输出："AB"

示例 3：
输入：columnNumber = 701
输出："ZY"

示例 4：
输入：columnNumber = 2147483647
输出："FXSHRXW

官方题解 https://leetcode.cn/problems/excel-sheet-column-title/solution/excelbiao-lie-ming-cheng-by-leetcode-sol-hgj4/
*/
func convertToTitle(columnNumber int) string {
	//输入：columnNumber = 701
	//输出："ZY" 26^1*26+25 = 701
	//...+26^2*(X3-A+1)+26^1*(X2-A+1)+(X1-A+1)
	// X3 X2 X1
	ans := []byte{}
	for columnNumber > 0 {
		a0 := (columnNumber-1)%26 + 1
		ans = append(ans, 'A'+byte(a0-1))
		columnNumber = (columnNumber - a0) / 26
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}
