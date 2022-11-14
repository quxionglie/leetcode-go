package problem100

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	n3 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n0 := &ListNode{
		Val:  0,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  4,
		Next: nil,
	}
	n3.Next = n2
	n2.Next = n0
	n0.Next = n4
	n4.Next = n2
	head := n3
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: head,
			},
			want: n2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := detectCycle(tt.args.head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
