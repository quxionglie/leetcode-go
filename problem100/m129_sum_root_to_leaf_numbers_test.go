package problem100

import "testing"

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			//从根到叶子节点路径 1->2 代表数字 12
			//从根到叶子节点路径 1->3 代表数字 13
			//因此，数字总和 = 12 + 13 = 25
			want: 25,
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 9,
						Left: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
				},
			},
			//从根到叶子节点路径 4->9->5 代表数字 495
			//从根到叶子节点路径 4->9->1 代表数字 491
			//从根到叶子节点路径 4->0 代表数字 40
			//因此，数字总和 = 495 + 491 + 40 = 1026
			want: 1026,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
