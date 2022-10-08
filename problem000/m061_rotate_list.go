package problem000

/*
*
61. 旋转链表
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]

  - Definition for singly-linked list.
  - type ListNode struct {
  - Val int
  - Next *ListNode
  - }

示例 2：
输入：head = [0,1,2], k = 4
输出：[2,0,1]

提示：

链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 10^9
*/
func rotateRight(head *ListNode, k int) *ListNode {
	//方法一：闭合为环
	//思路及算法
	//
	//记给定链表的长度为 nn，注意到当向右移动的次数 k \geq nk≥n 时，我们仅需要向右移动 k \bmod nkmodn 次即可。因为每 nn 次移动都会让链表变为原状。这样我们可以知道，新链表的最后一个节点为原链表的第 (n - 1) - (k \bmod n)(n−1)−(kmodn) 个节点（从 00 开始计数）。
	//
	//这样，我们可以先将给定的链表连接成环，然后将指定位置断开。
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n - k%n
	if add == n {
		return head
	}
	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
