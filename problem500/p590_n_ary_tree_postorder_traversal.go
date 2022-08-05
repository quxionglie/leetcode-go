package problem500

/**
qxl 590. N 叉树的后序遍历
给定一个 n 叉树的根节点 root ，返回 其节点值的 后序遍历 。
n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
*/
func postorder(root *Node) []int {
	result := []int{}
	if root == nil {
		return result
	}
	for _, child := range root.Children {
		result = append(result, postorder(child)...)
	}
	result = append(result, root.Val)
	return result
}
