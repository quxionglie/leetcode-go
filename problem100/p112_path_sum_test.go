package problem100

import (
	"leetcode/data_struct"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *data_struct.TreeNode
		targetSum int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				root: &data_struct.TreeNode{
					Val: 1,
					Left: &data_struct.TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &data_struct.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				root:      nil,
				targetSum: 5,
			},
			want: false,
		},
		{
			name: "",
			args: args{
				root: &data_struct.TreeNode{
					Val: 5,
					Left: &data_struct.TreeNode{
						Val: 4,
						Left: &data_struct.TreeNode{
							Val: 11,
							Left: &data_struct.TreeNode{
								Val:   7,
								Left:  nil,
								Right: nil,
							},
							Right: &data_struct.TreeNode{
								Val:   2,
								Left:  nil,
								Right: nil,
							},
						},
						Right: nil,
					},
					Right: &data_struct.TreeNode{
						Val: 8,
						Left: &data_struct.TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
						Right: &data_struct.TreeNode{
							Val:  4,
							Left: nil,
							Right: &data_struct.TreeNode{
								Val:   1,
								Left:  nil,
								Right: nil,
							},
						},
					},
				},
				targetSum: 22,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasPathSum(tt.args.root, tt.args.targetSum)
			if got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
