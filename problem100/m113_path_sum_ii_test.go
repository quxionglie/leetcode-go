package problem100

import (
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *TreeNode
		targetSum int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 11,
							Left: &TreeNode{
								Val:   7,
								Left:  nil,
								Right: nil,
							},
							Right: &TreeNode{
								Val:   2,
								Left:  nil,
								Right: nil,
							},
						},
						Right: nil,
					},
					Right: &TreeNode{
						Val: 8,
						Left: &TreeNode{
							Val:   13,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val: 4,
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
					},
				},
				targetSum: 22,
			},
			wantAns: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
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
				targetSum: 5,
			},
			wantAns: [][]int{},
		},
		{
			name: "",
			args: args{
				root: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
				targetSum: 0,
			},
			wantAns: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("pathSum() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
