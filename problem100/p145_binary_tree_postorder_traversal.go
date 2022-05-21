package problem100

/**
145. 二叉树的后序遍历
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func postorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = append(result, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, postorderTraversal(root.Right)...)
	}
	result = append(result, root.Val)
	return result
}
