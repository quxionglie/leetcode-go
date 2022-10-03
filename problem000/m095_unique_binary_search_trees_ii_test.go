package problem000

import (
	"testing"
)

func Test_generateTrees(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "",
			args: args{
				n: 2,
			},
			want: []*TreeNode{
				{
					Val:  1,
					Left: nil,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				{
					Val: 2,
					Left: &TreeNode{
						Val:   1,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateTrees(tt.args.n)

			for _, gotNode := range got {
				var ok bool
				for _, wantNode := range tt.want {
					if isTreeEq(gotNode, wantNode) {
						ok = true
					}
				}
				if !ok {
					t.Errorf("generateTrees() = %v, want %v", gotNode, tt.want)
				}
			}

		})
	}
}
