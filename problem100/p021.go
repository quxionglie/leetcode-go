package problem100

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, l3 *ListNode
	l3 = &ListNode{-1, nil}
	head = l3

	for l1 != nil || l2 != nil {
		if l1 != nil && (l2 == nil || l1.Val < l2.Val) {
			l3.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			l3.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		if l3.Next != nil {
			l3 = l3.Next
		}
	}

	return head.Next
}
