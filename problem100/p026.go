package problem100

/**
26. 删除有序数组中的重复项
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
*/
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	pos := 0
	for i := 1; i < n; i++ {
		if nums[pos] < nums[i] {
			pos++
			if pos < i {
				//有重复项时,pos必小于i
				nums[pos] = nums[i]
			}
		}
	}
	return pos + 1
}
