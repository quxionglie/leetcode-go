package problem100

import (
	"testing"
)

/**
示例 1：
输入：root = [3,9,20,null,null,15,7]
  3
 / \
 9  20
   /  \
   15  7
输出：2

示例 2：
输入：root = [2,null,3,null,4,null,5,null,6]
  2
   \
    3
     \
      4
       \
        5
         \
          6
输出：5
*/
func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
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
		}}, 2},
		{"", args{&TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
						Left: &TreeNode{
							Val:   6,
							Left:  nil,
							Right: nil,
						},
						Right: nil,
					},
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
