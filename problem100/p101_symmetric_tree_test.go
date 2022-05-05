package problem100

import (
	"leetcode/data_struct"
	"testing"
)

/**
    1
  2    2
3  4  4  3
*/
func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *data_struct.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{
			root: &data_struct.TreeNode{Val: 1,
				Left: &data_struct.TreeNode{Val: 2,
					Left: &data_struct.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &data_struct.TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil}},
				Right: &data_struct.TreeNode{
					Val: 2,
					Left: &data_struct.TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &data_struct.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil},
				}}}, true}, //[1,2,2,3,4,4,3] true
		{"", args{
			root: &data_struct.TreeNode{
				Val: 1,
				Left: &data_struct.TreeNode{Val: 2,
					Left: nil,
					Right: &data_struct.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil}},
				Right: &data_struct.TreeNode{
					Val:  2,
					Left: nil,
					Right: &data_struct.TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil},
				}}}, false}, //[1,2,2,null,3,null,3] false
		{"", args{&data_struct.TreeNode{}}, true}, //[1,2,2,null,3,null,3] false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
