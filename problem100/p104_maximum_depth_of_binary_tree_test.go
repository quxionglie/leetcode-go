package problem100

import (
	"leetcode/data_struct"
	"testing"
)

/**
    3
   / \
  9  20
    /  \
   15   7

retru 3
*/
func Test_maxDepth(t *testing.T) {
	type args struct {
		root *data_struct.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{
			root: &data_struct.TreeNode{Val: 3,
				Left: &data_struct.TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &data_struct.TreeNode{
					Val: 20,
					Left: &data_struct.TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &data_struct.TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil},
				}}}, 3},
		{"", args{root: nil}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
