package problem200

/**
226. 翻转二叉树
给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
 2
/ \
1 3
翻转后：
 2
/ \
3 1
https://leetcode.cn/problems/invert-binary-tree/
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	//交换两棵子树
	root.Left = right
	root.Right = left
	return root
}
