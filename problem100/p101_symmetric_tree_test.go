package problem100

import (
	"testing"
)

/**
    1
  2    2
3  4  4  3
*/
func Test_isSymmetric(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil}},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil},
				}}}, true}, //[1,2,2,3,4,4,3] true
		{"", args{
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil}},
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil},
				}}}, false}, //[1,2,2,null,3,null,3] false
		{"", args{&TreeNode{}}, true}, //[1,2,2,null,3,null,3] false
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
