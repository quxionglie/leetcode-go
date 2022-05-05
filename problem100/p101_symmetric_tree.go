package problem100

import "leetcode/data_struct"

/**
101. 对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
func isSymmetric(root *data_struct.TreeNode) bool {
	//递归解决: 引入一个队列，这是把递归程序改写成迭代程序的常用方法。
	//初始化时我们把根节点入队两次。每次提取两个结点并比较它们的值（队列中每两个连续的结点应该是相等的，而且它们的子树互为镜像），
	//然后将两个结点的左右子结点按相反的顺序插入队列中。当队列为空时，
	//或者我们检测到树不对称（即从队列中取出两个不相等的连续结点）时，该算法结束。
	queue := []*data_struct.TreeNode{}
	queue = append(queue, root)
	queue = append(queue, root)
	for len(queue) > 0 {
		curLeft := queue[0]
		curRight := queue[1]
		queue = queue[2:]

		if curLeft == nil && curRight == nil {
			continue
		}
		if curLeft == nil || curRight == nil {
			return false
		}
		if curLeft.Val != curRight.Val {
			return false
		}

		queue = append(queue, curLeft.Left)
		queue = append(queue, curRight.Right)

		queue = append(queue, curLeft.Right)
		queue = append(queue, curRight.Left)
	}

	return true
}
