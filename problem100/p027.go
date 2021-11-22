package problem100

/**
27. 移除元素
给你一个数组 nums和一个值 val，你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
说明:
为什么返回数值是整数，但输出的答案是数组呢?
请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
链接：https://leetcode-cn.com/problems/remove-element
*/
func removeElement(nums []int, val int) int {
	n := len(nums)
	pos := 0
	for i := 0; i < n; i++ {
		if nums[i] == val {
			continue
		} else {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}
