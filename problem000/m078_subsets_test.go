package problem000

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3},
			},
			wantAns: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAns := subsets(tt.args.nums)
			var flag bool
			for _, an := range gotAns {
				for _, wantAn := range tt.wantAns {
					if !reflect.DeepEqual(an, wantAn) {
						flag = true
					}
				}
				if !flag {
					t.Errorf("subsets() = %v, want %v", gotAns, tt.wantAns)
				}
			}

		})
	}
}
