package problem100

import (
	"testing"
)

//-2^31 <= x <= 2^31 - 1
func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"01", args{123}, 321},
		{"02", args{-123}, -321},
		{"03", args{12}, 21},
		{"04", args{2147483647}, 0},
		{"05", args{-2147483647}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
