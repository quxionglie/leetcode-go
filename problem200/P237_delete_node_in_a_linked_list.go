package problem200

/**
237. 删除链表中的节点
请编写一个函数，用于 删除单链表中某个特定节点 。在设计函数时需要注意，你无法访问链表的头节点 head ，只能直接访问 要被删除的节点 。
题目数据保证需要删除的节点 不是末尾节点 。

 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
*/
func deleteNode(node *ListNode) {
	next := node.Next
	node.Val = next.Val
	node.Next = next.Next
}
