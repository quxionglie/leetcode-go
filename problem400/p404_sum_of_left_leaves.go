package problem400

/**
qxl 404. 左叶子之和
给定二叉树的根节点 root ，返回所有左叶子之和。

给定二叉树的根节点root，返回所有左叶子之和。
示例 1：
输入: root = [3,9,20,null,null,15,7]
输出: 24
解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

示例2:
输入: root = [1]
输出: 0

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return doSumOfLeftLeaves(root.Left, true) + doSumOfLeftLeaves(root.Right, false)
}

func doSumOfLeftLeaves(root *TreeNode, isLeftNode bool) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if isLeftNode {
			return root.Val
		} else {
			return 0
		}
	}
	return doSumOfLeftLeaves(root.Left, true) + doSumOfLeftLeaves(root.Right, false)
}
