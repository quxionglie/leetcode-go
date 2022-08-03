package problem500

import (
	"reflect"
	"testing"
)

func Test_preorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						&Node{
							Val: 3,
							Children: []*Node{
								&Node{
									Val:      5,
									Children: []*Node{},
								},
								&Node{
									Val:      6,
									Children: []*Node{},
								},
							},
						},
						&Node{
							Val:      2,
							Children: []*Node{},
						},
						&Node{
							Val:      4,
							Children: []*Node{},
						},
					},
				},
			},
			want: []int{1, 3, 5, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
