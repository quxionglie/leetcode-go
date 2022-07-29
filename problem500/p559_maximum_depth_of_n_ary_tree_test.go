package problem500

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: &Node{
					Val: 0,
					Children: []*Node{
						&Node{
							Val: 3,
							Children: []*Node{
								&Node{
									Val:      5,
									Children: nil,
								},
								&Node{
									Val:      6,
									Children: nil,
								},
							},
						},
						&Node{
							Val:      2,
							Children: nil,
						},
						&Node{
							Val:      4,
							Children: nil,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
