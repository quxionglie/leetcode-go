package problem100

/**
144. 二叉树的前序遍历
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func preorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	if root.Left != nil {
		result = append(result, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, preorderTraversal(root.Right)...)
	}
	return result
}
