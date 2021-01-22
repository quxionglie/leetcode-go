package problem100

import (
	"math"
)

// https://leetcode-cn.com/problems/reverse-integer/
// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
//注意：
//假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为[−2^31,2^31−1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
// 需要考虑的情况：
// 1.如何反转
// pop := x % 10
// x = x / 10
// rev = rev*10 + pop

// 2.正数和负数什么时候会溢出
// int32  : -2147483648 to 2147483647
// - 当 rev>INT_MAX/10,不可能再rev*10 或 rev==INT_MAX/10且pop>7
// - 当 rev<INT_MAX/10,不可能再rev*10 或 rev==INT_MIN/10且pop<-8
func reverse(x int) int {
	// 321 -> 123
	rev := 0
	for x != 0 {
		pop := x % 10
		x = x / 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}

	return rev
}
