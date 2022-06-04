package problem200

import (
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	inputNode := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "",
			args: args{
				root: inputNode,
				p: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				q: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
			},
			want: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		{
			name: "",
			args: args{
				root: inputNode,
				p: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				q: &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
			},
			want: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); got.Val != tt.want.Val {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
