package problem000

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			name: "",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name: "",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.args.candidates, tt.args.target)
			var eq bool
			for _, g := range got {
				for _, w := range tt.want {
					if reflect.DeepEqual(g, w) {
						eq = true
						break
					}
				}
				if !eq {
					t.Errorf("combinationSum() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
