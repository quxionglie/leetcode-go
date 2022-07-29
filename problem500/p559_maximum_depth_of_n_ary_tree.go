package problem500

/**
qxl 559. N 叉树的最大深度
给定一个 N 叉树，找到其最大深度。
最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。
N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。

 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
*/

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	left := doMaxDepth(root.Children) + 1
	return left
}
func doMaxDepth(nodes []*Node) int {
	var max int
	for _, node := range nodes {
		dep := maxDepth(node)
		if dep > max {
			max = dep
		}
	}
	return max
}
