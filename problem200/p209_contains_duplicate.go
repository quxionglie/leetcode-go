package problem200

/**
217. 存在重复元素
给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。
示例 1：

输入：nums = [1,2,3,1]
输出：true
示例 2：

输入：nums = [1,2,3,4]
输出：false
示例3：

输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true
*/
func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		if m[num] > 0 {
			return true
		}
		m[num] = m[num] + 1
	}
	return false
}
