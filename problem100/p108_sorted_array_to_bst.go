package problem100

/**
108. 将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func sortedArrayToBST(nums []int) *TreeNode {

	return doBuildBst(nums, 0, len(nums)-1)
}

func doBuildBst(nums []int, left, right int) *TreeNode {

	if left > right {
		return nil
	}
	//中序遍历，总是选择中间位置左边的数字作为根节点
	mid := (left + right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = doBuildBst(nums, left, mid-1)
	root.Right = doBuildBst(nums, mid+1, right)
	return root
}
