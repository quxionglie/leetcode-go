package problem400

import "math"

/**
qxl 414. 第三大的数
给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。
示例 1：
输入：[3, 2, 1]
输出：1
解释：第三大的数是 1 。

示例 2：
输入：[1, 2]
输出：2
解释：第三大的数不存在, 所以返回最大的数 2 。

示例 3：
输入：[2, 2, 3, 1]
输出：1
解释：注意，要求返回第三大的数，是指在所有不同数字中排第三大的数。
此例中存在两个值为 2 的数，它们都排第二。在所有不同数字中排第三大的数为 1 。

提示：
1 <= nums.length <= 104
-231 <= nums[i] <= 231 - 1
*/
func thirdMax(nums []int) int {
	result := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	var isHasMinInt32 bool
	for _, num := range nums {
		if num == math.MinInt32 {
			isHasMinInt32 = true
		}
		if num > result[0] {
			result = []int{num, result[0], result[1]}
		} else if num < result[0] && num > result[1] {
			result = []int{result[0], num, result[1]}
		} else if num < result[0] && num < result[1] && num > result[2] {
			result = []int{result[0], result[1], num}
		}
	}

	if isHasMinInt32 && result[1] != math.MinInt32 {
		return math.MinInt32
	} else if result[2] == math.MinInt32 {
		return result[0]
	} else {
		return result[2]
	}
	return result[0]
}
