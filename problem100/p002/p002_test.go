package p002

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//提示：
//每个链表中的节点数在范围 [1, 100] 内
//0 <= Node.val <= 9
//题目数据保证列表表示的数字不含前导零
func TestAddTwoNumbers(t *testing.T) {
	var l1, l2, l3 *ListNode
	//示例 1：
	//输入：l1 = [2,4,3], l2 = [5,6,4]
	//输出：[7,0,8]
	//解释：342 + 465 = 807.
	l1 = &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 = &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l3 = addTwoNumbers(l1, l2)
	assertList(t, l3, []int{7, 0, 8})

	//示例 2：
	//输入：l1 = [0], l2 = [0]
	//输出：[0]
	l1 = &ListNode{Val: 0}
	l2 = &ListNode{Val: 0}
	l3 = addTwoNumbers(l1, l2)
	assertList(t, l3, []int{0})

	//示例 3：
	//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
	//输出：[8,9,9,9,0,0,0,1]
	l1 = newListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 = newListNode([]int{9, 9, 9, 9})
	l3 = addTwoNumbers(l1, l2)
	assertList(t, l3, []int{8, 9, 9, 9, 0, 0, 0, 1})
}

func newListNode(in []int) *ListNode {
	var head, tail *ListNode
	for _, v := range in {
		if head == nil && tail == nil {
			head = &ListNode{Val: v}
			tail = head
		} else {
			tail.Next = &ListNode{Val: v}
			tail = tail.Next
		}
	}
	return head
}
func assertList(t *testing.T, l3 *ListNode, outs []int) {
	for _, out := range outs {
		if l3 != nil {
			assert.Equal(t, out, l3.Val)
		}
		l3 = l3.Next
	}
}
