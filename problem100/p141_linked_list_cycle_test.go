package problem100

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{createListNode([]int{3, 2, 0, 4}, 1)},
			want: true,
		},
		{
			name: "",
			args: args{createListNode([]int{1, 2}, 0)},
			want: true,
		},
		{
			name: "",
			args: args{createListNode([]int{1}, -1)},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createListNode(v []int, pos int) *ListNode {
	if len(v) == 0 {
		return nil
	}
	var head *ListNode
	var tail *ListNode
	var cycleNode *ListNode

	for i, av := range v {
		curNode := &ListNode{
			Val:  av,
			Next: nil,
		}

		if pos != -1 && i == pos {
			cycleNode = curNode
		}
		if head == nil {
			head = curNode
			tail = curNode
		} else {
			tail.Next = curNode
			tail = curNode
		}
	}
	tail.Next = cycleNode
	return head
}
