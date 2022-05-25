package problem100

import "math"

/**
169. 多数元素

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func majorityElement(nums []int) int {
	m := make(map[int]int)
	maxK := 0
	maxV := math.MinInt32
	for _, num := range nums {
		m[num] = m[num] + 1
		if m[num] > maxV {
			maxK = num
			maxV = m[num]
		}
	}
	return maxK
}
