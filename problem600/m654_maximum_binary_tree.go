package problem600

/*
*
654. 最大二叉树
给定一个不重复的整数数组nums 。最大二叉树可以用下面的算法从nums 递归地构建:

创建一个根节点，其值为nums 中的最大值。
递归地在最大值左边的子数组前缀上构建左子树。
递归地在最大值 右边 的子数组后缀上构建右子树。
返回nums 构建的 最大二叉树 。
输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	best := 0
	for i, num := range nums {
		if num > nums[best] {
			best = i
		}
	}
	return &TreeNode{
		Val:   nums[best],
		Left:  constructMaximumBinaryTree(nums[:best]),
		Right: constructMaximumBinaryTree(nums[best+1:]),
	}
}
