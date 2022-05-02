package problem100

import "math"

/**
110. 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1 。

示例 1：
输入：root = [3,9,20,null,null,15,7]
输出：true

示例 2：
输入：root = [1,2,2,3,3,null,null,4,4]
输出：false

示例 3：
输入：root = []
输出：true
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, isBalanced := maxTreeDepth(root)
	return isBalanced
}

func maxTreeDepth(root *TreeNode) (int, bool) {
	//递归实现
	if root == nil {
		return 0, true
	}
	depthLeft, isBalancedLeft := maxTreeDepth(root.Left)
	depthRight, isBalancedRight := maxTreeDepth(root.Right)
	if isBalancedLeft && isBalancedRight {
		if math.Abs(float64(depthLeft-depthRight)) > 1.0 {
			return 0, false
		}
		if depthLeft > depthRight {
			return 1 + depthLeft, true
		} else {
			return 1 + depthRight, true
		}
	}
	return 0, false
}
