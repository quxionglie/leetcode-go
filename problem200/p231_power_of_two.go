package problem200

/**
231. 2 的幂
给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。
如果存在一个整数 x 使得 n == 2x ，则认为 n 是 2 的幂次方。

示例 1：
输入：n = 1
输出：true
解释：20 = 1

示例 2：
输入：n = 16
输出：true
解释：24 = 16

示例 3：
输入：n = 3
输出：false
*/
func isPowerOfTwo(n int) bool {
	//一个数 n 是 2 的幂，当且仅当 n 是正整数，并且 n 的二进制表示中仅包含 1 个 1。
	return n > 0 && n&(n-1) == 0
}
