package problem200

import (
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: createLinkList([]int{1, 2, 6, 3, 4, 5, 6}),
				val:  6,
			},
			want: createLinkList([]int{1, 2, 3, 4, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !eqLinkList(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func eqLinkList(n1, n2 *ListNode) bool {
	for n1 != nil || n2 != nil {
		if (n1 == nil && n2 != nil) || (n1 != nil && n2 == nil) {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		n1 = n1.Next
		n2 = n2.Next
	}
	return true
}

func createLinkList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	var head, tail *ListNode
	for _, num := range nums {
		if head == nil {
			head = &ListNode{
				Val:  num,
				Next: nil,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val:  num,
				Next: nil,
			}
			tail = tail.Next
		}
	}
	return head
}
