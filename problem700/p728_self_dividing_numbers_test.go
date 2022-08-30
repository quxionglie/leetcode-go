package problem700

import (
	"reflect"
	"testing"
)

func Test_selfDividingNumbers(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	var tests = []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "",
			args: args{
				left:  1,
				right: 22,
			},
			wantAns: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := selfDividingNumbers(tt.args.left, tt.args.right); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("selfDividingNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
