package problem000

/*
*
24. 两两交换链表中的节点
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]
提示：

链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100

  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }
*/
func swapPairs(head *ListNode) *ListNode {
	//递归
	//可以通过递归的方式实现两两交换链表中的节点。
	//递归的终止条件是链表中没有节点，或者链表中只有一个节点，此时无法进行交换。
	//如果链表中至少有两个节点，则在两两交换链表中的节点之后，原始链表的头节点变成新的链表的第二个节点，
	//原始链表的第二个节点变成新的链表的头节点。链表中的其余节点的两两交换可以递归地实现。
	//在对链表中的其余节点递归地两两交换之后，更新节点之间的指针关系，即可完成整个链表的两两交换。
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}
