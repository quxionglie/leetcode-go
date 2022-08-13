package problem600

import "sort"

/*
*
628. 三个数的最大乘积
给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1：
输入：nums = [1,2,3]
输出：6

示例 2：
输入：nums = [1,2,3,4]
输出：24

示例 3：
输入：nums = [-1,-2,-3]
输出：-6
提示：
3 <= nums.length <=10^4
-1000 <= nums[i] <= 1000
*/
func maximumProduct(nums []int) int {
	//方法一：排序
	//首先将数组排序。
	//如果数组中全是非负数，则排序后最大的三个数相乘即为最大乘积；如果全是非正数，则最大的三个数相乘同样也为最大乘积。
	//如果数组中有正数有负数，则最大乘积既可能是三个最大正数的乘积，也可能是两个最小负数（即绝对值最大）与最大正数的乘积。
	//综上，我们在给数组排序后，分别求出三个最大正数的乘积，以及两个最小负数与最大正数的乘积，二者之间的最大值即为所求答案。
	sort.Ints(nums)
	n := len(nums)

	r1 := nums[0] * nums[1] * nums[n-1]
	r2 := nums[n-3] * nums[n-2] * nums[n-1]
	if r1 > r2 {
		return r1
	}
	return r2

}
