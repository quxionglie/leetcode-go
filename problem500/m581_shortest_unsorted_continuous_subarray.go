package problem500

import "sort"

/*
*
581. 最短无序连续子数组
给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
请你找出符合题意的 最短 子数组，并输出它的长度。

示例 1：
输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。

示例 2：
输入：nums = [1,2,3,4]
输出：0

示例 3：
输入：nums = [1]
输出：0

提示：
1 <= nums.length <= 10^4
-10^5 <= nums[i] <= 10^5
*/
func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted)
	left, right := 0, len(nums)-1
	for nums[left] == numsSorted[left] {
		left++
	}
	for nums[right] == numsSorted[right] {
		right--
	}
	return right - left + 1
}
