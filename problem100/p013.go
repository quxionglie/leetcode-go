package problem100

import "strings"

//https://leetcode-cn.com/problems/roman-to-integer/
//罗马数字包含以下七种字符:I，V，X，L，C，D和M。
//
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//例如， 罗马数字 2 写做II，即为两个并列的 1。12 写做XII，即为X+II。 27 写做XXVII, 即为XX+V+II。
//
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做IIII，而是IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为IX。这个特殊的规则只适用于以下六种情况：
//
//I可以放在V(5) 和X(10) 的左边，来表示 4 和 9。
//X可以放在L(50) 和C(100) 的左边，来表示 40 和90。
//C可以放在D(500) 和M(1000) 的左边，来表示400 和900。
//给定一个罗马数字，将其转换成整数。输入确保在 1到 3999 的范围内。
//

//解法1
func romanToInt(s string) int {
	basic := make(map[string]int)
	basic["I"] = 1
	basic["V"] = 5
	basic["X"] = 10
	basic["L"] = 50
	basic["C"] = 100
	basic["D"] = 500
	basic["M"] = 1000

	com := make(map[string]int)
	com["IV"] = 4
	com["IX"] = 9
	com["XL"] = 40
	com["XC"] = 90
	com["CD"] = 400
	com["CM"] = 900

	sum := 0
	slen := len(s)
	for slen > 0 {
		for k := range com {
			out := strings.Replace(s, k, "", 1)
			if out != s {
				sum = sum + com[k]
				s = out
			}
		}
		if slen == len(s) {
			break
		}
		slen = len(s)
	}
	for _, c := range s {
		sum = sum + basic[string(c)]
	}
	return sum
}
