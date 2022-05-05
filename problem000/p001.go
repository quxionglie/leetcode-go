package problem000

/**
https://leetcode-cn.com/problems/two-sum/
给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
你可以按任意顺序返回答案。
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := range nums {
		v := target - nums[i]
		if val, ok := m[v]; ok {
			return []int{val, i}
		}
		m[nums[i]] = i
	}
	return nil
}
