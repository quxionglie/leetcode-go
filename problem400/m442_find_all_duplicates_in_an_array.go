package problem400

/*
*
442. 数组中重复的数据

给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。

你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。

示例 1：
输入：nums = [4,3,2,7,8,2,3,1]
输出：[2,3]

示例 2：
输入：nums = [1,1,2]
输出：[1]

示例 3：
输入：nums = [1]
输出：[]
*/
func findDuplicates(nums []int) (ans []int) {
	//方法二：使用正负号作为标记
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		} else {
			ans = append(ans, x)
		}
	}
	return
}
