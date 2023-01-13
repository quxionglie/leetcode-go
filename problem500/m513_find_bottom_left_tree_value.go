package problem500

/*
*
513. 找树左下角的值
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
假设二叉树中至少有一个节点。
示例 1:
输入: root = [2,1,3]
输出: 1

示例 2:
输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/
func findBottomLeftValue(root *TreeNode) (curVal int) {
	//方法一：深度优先搜索
	//使用 height 记录遍历到的节点的高度，curVal 记录高度在 curHeight 的最左节点的值。
	//在深度优先搜索时，我们先搜索当前节点的左子节点，再搜索当前节点的右子节点，
	//然后判断当前节点的高度 height 是否大于 curHeight，如果是，那么将 curVal 设置为当前结点的值，
	//curHeight 设置为 height。

	curHeight := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		if height > curHeight {
			curHeight = height
			curVal = node.Val
		}
	}
	dfs(root, 0)
	return
}
