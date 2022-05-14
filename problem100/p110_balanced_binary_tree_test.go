package problem100

import (
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val:   15,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   7,
					Left:  nil,
					Right: nil,
				},
			},
		}}, true}, //[3,9,20,null,null,15,7]
		{"", args{&TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:  2,
				Left: nil,
				Right: &TreeNode{
					Val:  3,
					Left: nil,
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
			},
		}}, false}, //[1,2,2,3,null,null,3,4,null,null,4]
		/**
				  1
				/   \
				2     2
			  /        \
			 3          3
			/            \
		   4              4
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
