package problem300

/**
qxl 349. 两个数组的交集
给定两个数组nums1和nums2 ，返回 它们的交集。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。

示例 1：
输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2]

示例 2：
输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[9,4]
解释：[4,9] 也是可通过的
*/
func intersection(nums1 []int, nums2 []int) []int {
	resultSet := map[int]int{}
	set1 := map[int]int{}
	for _, v := range nums1 {
		set1[v]++
	}
	for _, v := range nums2 {
		if _, ok := set1[v]; ok {
			resultSet[v]++
		}
	}

	result := []int{}
	for k, _ := range resultSet {
		result = append(result, k)
	}

	return result
}
