package problem100

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"01", args{newListNode([]int{1, 2, 4}), newListNode([]int{1, 3, 4})}, newListNode([]int{1, 1, 2, 3, 4, 4})},
		{"02", args{newListNode([]int{}), newListNode([]int{})}, newListNode([]int{})},
		{"03", args{newListNode([]int{0}), newListNode([]int{})}, newListNode([]int{0})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !assertList(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func assertList(l3 *ListNode, l4 *ListNode) bool {
	for l3 != nil || l4 != nil {
		if l3 == nil || l4 == nil || l3.Val != l4.Val {
			return false
		}
		l3 = l3.Next
		l4 = l4.Next
	}
	return true
}
