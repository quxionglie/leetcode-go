package problem000

/*
*
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]
提示：
链表中结点的数目为 sz
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz
  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//方法一：计算链表长度
	//一种容易想到的方法是，我们首先从头节点开始对链表进行一次遍历，
	//得到链表的长度 L。随后我们再从头节点开始对链表进行一次遍历，
	//当遍历到第 L-n+1 个节点时，它就是我们需要删除的节点。

	//方法二：栈
	//遍历链表的同时将所有节点依次入栈。
	//根据栈「先进后出」的原则，我们弹出栈的第 n 个节点就是需要删除的节点，
	//并且目前栈顶的节点就是待删除节点的前驱节点。这样一来，删除操作就变得十分方便了。
	nodes := []*ListNode{}
	dummy := &ListNode{0, head}
	for node := dummy; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummy.Next
}
