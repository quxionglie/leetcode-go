package problem400

import (
	"reflect"
	"testing"
)

func Test_constructRectangle(t *testing.T) {
	type args struct {
		area int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{4},
			want: []int{2, 2},
		},
		{
			name: "",
			args: args{37},
			want: []int{37, 1},
		},
		{
			name: "",
			args: args{122122},
			want: []int{427, 286},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := constructRectangle(tt.args.area)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
