package problem000

import (
	"reflect"
	"testing"
)

//提示：
//每个链表中的节点数在范围 [1, 100] 内
//0 <= Node.val <= 9
//题目数据保证列表表示的数字不含前导零
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
func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"01", args{newListNode([]int{2, 4, 3}), newListNode([]int{5, 6, 4})}, newListNode([]int{7, 0, 8})},
		{"01", args{newListNode([]int{0}), newListNode([]int{0})}, newListNode([]int{0})},
		{"01", args{newListNode([]int{9, 9, 9, 9, 9, 9, 9}), newListNode([]int{9, 9, 9, 9})}, newListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
