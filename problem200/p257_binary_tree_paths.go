package problem200

import (
	"strconv"
	"strings"
)

/**
257. 二叉树的所有路径
给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
叶子节点 是指没有子节点的节点。
  1
 /  \
2    3
  \
   5
输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]

 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
*/
var result []string

func binaryTreePaths(root *TreeNode) []string {
	result = []string{}
	parent := []string{}
	doBinaryTreePaths(parent, root)
	return result
}

func doBinaryTreePaths(parent []string, root *TreeNode) {
	parent = append(parent, strconv.Itoa(root.Val))

	if root.Left == nil && root.Right == nil {
		result = append(result, strings.Join(parent, "->"))
	}
	if root.Left != nil {
		doBinaryTreePaths(parent, root.Left)
	}
	if root.Right != nil {
		doBinaryTreePaths(parent, root.Right)
	}
}
