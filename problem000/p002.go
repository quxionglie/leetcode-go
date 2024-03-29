package problem000

// https://leetcode-cn.com/problems/add-two-numbers/
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1 := 0
		n2 := 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		carry = sum / 10

		if head == nil || tail == nil {
			head = &ListNode{sum % 10, nil}
			tail = head
		} else {
			tail.Next = &ListNode{sum % 10, nil}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}
	return head
}
