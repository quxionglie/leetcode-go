package problem600

import (
	"fmt"
	"strconv"
)

/*
*
606. 根据二叉树创建字符串

给你二叉树的根节点 root ，请你采用前序遍历的方式，将二叉树转化为一个由括号和整数组成的字符串，返回构造出的字符串。
空节点使用一对空括号对 "()" 表示，转化后需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。
示例 1：

		  1
		 / \
		2   3
	   /
	  4

输入：root = [1,2,3,4]
输出："1(2(4))(3)"
解释：初步转化后得到 "1(2(4)())(3()())" ，但省略所有不必要的空括号对后，字符串应该是"1(2(4))(3)" 。

  - Definition for a binary tree node.
  - type TreeNode struct {
  - Val int
  - Left *TreeNode
  - Right *TreeNode
  - }
*/
func tree2str(root *TreeNode) string {
	switch {
	case root == nil:
		return ""
	case root.Left == nil && root.Right == nil:
		return strconv.Itoa(root.Val)
	case root.Right == nil:
		return fmt.Sprintf("%d(%s)", root.Val, tree2str(root.Left))
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.Val, tree2str(root.Left), tree2str(root.Right))
	}
}
