package problem100

/*
*
105. 从前序与中序遍历序列构造二叉树
给定两个整数数组preorder 和 inorder，其中preorder 是二叉树的先序遍历，
inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。
示例 1:

		 3
		/ \
	   9  20
		 / \
		15 7

输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
输出: [3,9,20,null,null,15,7]
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
