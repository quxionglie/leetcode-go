package problem600

/*
*
671. 二叉树中第二小的节点
给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为2或0。
如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一个。
更正式地说，即root.val = min(root.left.val, root.right.val) 总成立。
给出这样的一个二叉树，你需要输出所有节点中的第二小的值 。
如果第二小的值不存在的话，输出 -1 。

  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/
func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	rootVal := root.Val
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil || ans != -1 && node.Val >= ans {
			return
		}
		if node.Val > rootVal {
			ans = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}