package problem700

import "math"

/*
*
739. 每日温度
给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，
其中answer[i]是指对于第 i 天，下一个更高温度出现在几天后。如果气温在这之后都不会升高，请在该位置用0 来代替。

示例 1:
输入: temperatures = [73,74,75,71,69,72,76,73]
输出:[1,1,4,2,1,1,0,0]

示例 2:
输入: temperatures = [30,40,50,60]
输出:[1,1,1,0]

示例 3:
输入: temperatures = [30,60,90]
输出: [1,1,0]

提示：
1 <= temperatures.length <= 10^5
30 <= temperatures[i] <= 100
*/
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	next := make([]int, 101)
	for i := 0; i < 101; i++ {
		next[i] = math.MaxInt32
	}
	for i := length - 1; i >= 0; i-- {
		warmerIndex := math.MaxInt32
		for t := temperatures[i] + 1; t <= 100; t++ {
			if next[t] < warmerIndex {
				warmerIndex = next[t]
			}
		}
		if warmerIndex < math.MaxInt32 {
			ans[i] = warmerIndex - i
		}
		next[temperatures[i]] = i
	}
	return ans
}
