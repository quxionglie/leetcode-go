package problem000

/*
*
39. 组合总和
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，
找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。

示例 1：
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。

示例 2：
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]

示例 3：
输入: candidates = [2], target = 1
输出: []
*/
func combinationSum(candidates []int, target int) (ans [][]int) {
	//方法一：搜索回溯
	//思路与算法
	//对于这类寻找所有可行解的题，我们都可以尝试用「搜索回溯」的方法来解决。
	//回到本题，我们定义递归函数 dfs(target,combine,idx)
	//表示当前在 candidates 数组的第 idx 位，还剩  target 要组合，
	//已经组合的列表为 combine。递归的终止条件为 target≤0
	//或者 candidates 数组被全部用完。那么在当前的函数中，
	//每次我们可以选择跳过不用第  idx 个数，即执行 dfs(target,combine,idx+1)。也可以选择使用第 idx 个数，
	//即执行dfs(target−candidates[idx],combine,idx)，
	//注意到每个数字可以被无限制重复选取，因此搜索的下标仍为 idx。
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}
