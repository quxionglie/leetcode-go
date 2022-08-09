package problem600

/*
605. 种花问题
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
给你一个整数数组  flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，
能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false。

示例 1：
输入：flowerbed = [1,0,0,0,1], n = 1
输出：true

示例 2：
输入：flowerbed = [1,0,0,0,1], n = 2
输出：false

提示：
1 <= flowerbed.length <= 2 * 10^4
flowerbed[i] 为 0 或 1
flowerbed 中不存在相邻的两朵花
0 <= n <= flowerbed.length
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	/**
	方法一：贪心
	根据上述计算方法，计算花坛中可以种入的花的最多数量，判断是否大于或等于 n 即可。具体做法如下。
	维护 prev 表示上一朵已经种植的花的下标位置，初始时 prev=−1，表示尚未遇到任何已经种植的花。
	从左往右遍历数组 flowerbed，当遇到 flowerbed[i]=1 时根据 prev 和 i 的值计算上一个区间内可以种植花的最多数量，
	然后令 prev=i，继续遍历数组 flowerbed 剩下的元素。
	遍历数组 flowerbed 结束后，根据数组 prev 和长度 mm 的值计算最后一个区间内可以种植花的最多数量。
	判断整个花坛内可以种入的花的最多数量是否大于或等于 n。
	*/
	count := 0
	m := len(flowerbed)
	prev := -1
	for i := 0; i < m; i++ {
		if flowerbed[i] == 1 {
			if prev < 0 {
				count += i / 2
			} else {
				count += (i - prev - 2) / 2
			}
			prev = i
		}
	}
	if prev < 0 {
		count += (m + 1) / 2
	} else {
		count += (m - prev - 1) / 2
	}
	return count >= n
}
