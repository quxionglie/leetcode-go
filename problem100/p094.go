package problem100

/**
给定一个二叉树的根节点 root ，返回它的 中序 遍历。
https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}
