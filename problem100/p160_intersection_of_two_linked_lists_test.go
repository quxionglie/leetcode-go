package problem100

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	same := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				headA: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  0,
						Next: same,
					},
				},
				headB: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: same,
					},
				},
			},
			want: same,
		},
		{
			name: "",
			args: args{
				headA: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  0,
						Next: nil,
					},
				},
				headB: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: nil,
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
