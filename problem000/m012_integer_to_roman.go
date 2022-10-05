package problem000

/*
*
12.整数转罗马数字

罗马数字包含以下七种字符：I，V，X，L，C，D和M。

字符数值
I1
V5
X10
L50
C100
D500
M1000
例如，罗马数字2写做II，即为两个并列的1。12写做XII，即为X+II。27写做XXVII,即为XX+V+II。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，
例如4不写做IIII，而是IV。数字1在数字5的左边，所表示的数等于大数5减小数1得到的数值4。
同样地，数字9表示为IX。这个特殊的规则只适用于以下六种情况：
I可以放在V(5)和X(10)的左边，来表示4和9。
X可以放在L(50)和C(100)的左边，来表示40和90。
C可以放在D(500)和M(1000)的左边，来表示400和900。
给你一个整数，将其转为罗马数字。

示例1:
输入:num=3
输出:"III"

示例2:
输入:num=4
输出:"IV"

示例3:
输入:num=9
输出:"IX"

示例4:
输入:num=58
输出:"LVIII"
解释:L=50,V=5,III=3.

示例5:
输入:num=1994
输出:"MCMXCIV"
解释:M=1000,CM=900,XC=90,IV=4.

提示：
1<=num<=3999
*/
func intToRoman(num int) string {
	//模拟
	//根据罗马数字的唯一表示法，
	//为了表示一个给定的整数  num，我们寻找不超过  num 的最大符号值，
	//将  num 减去该符号值，然后继续寻找不超过 num 的最大符号值，将该符号拼接在上一个找到的符号之后，
	//循环直至 num 为 00。最后得到的字符串即为 num 的罗马数字表示。
	var valueSymbols = []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}
