package problem100

import "container/list"

/*
*
199. 二叉树的右视图
给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
示例 1:
输入:[1,2,3,null,5,null,4]
输出:[1,3,4]

示例 2:
输入:[1,null,3]
输出:[1,3]

示例 3:
输入:[]
输出:[]

提示:
二叉树的节点个数的范围是 [0,100]
-100<= Node.val <= 100
  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/
/**
199. 二叉树的右视图
*/
func rightSideView(root *TreeNode) []int {
	queue := list.New()
	res := [][]int{}
	var finaRes []int
	if root == nil {
		return finaRes
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		tmp := []int{}
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		res = append(res, tmp)
	}
	//取每一层的最后一个元素
	for i := 0; i < len(res); i++ {
		finaRes = append(finaRes, res[i][len(res[i])-1])
	}
	return finaRes
}
