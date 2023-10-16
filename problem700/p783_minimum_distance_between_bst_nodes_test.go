package problem700

import "testing"

func Test_minDiffInBST(t *testing.T) {
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
					Val: 4,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   0,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 48,
						Left: &TreeNode{
							Val:   12,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   49,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDiffInBST(tt.args.root); got != tt.want {
				t.Errorf("minDiffInBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
