package problem000

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "",
			args: args{
				head: createLinkList([]int{1, 4, 3, 2, 5, 2}),
				x:    3,
			},
			want: createLinkList([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			name: "",
			args: args{
				head: createLinkList([]int{2, 1}),
				x:    3,
			},
			want: createLinkList([]int{1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
