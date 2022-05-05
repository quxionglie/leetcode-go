package problem100

import (
	"leetcode/data_struct"
	"testing"
)

func Test_isSameTree(t *testing.T) {
	type args struct {
		p *data_struct.TreeNode
		q *data_struct.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{&data_struct.TreeNode{1, &data_struct.TreeNode{2, nil, nil}, &data_struct.TreeNode{3, nil, nil}},
			&data_struct.TreeNode{1, &data_struct.TreeNode{2, nil, nil}, &data_struct.TreeNode{3, nil, nil}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSameTree(tt.args.p, tt.args.q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
