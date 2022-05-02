package problem100

import (
	"math"
	"testing"
)

func Test_sortedArrayToBST(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"", args{[]int{-10, -3, 0, 5, 9}}},
		{"", args{[]int{-10, -3, 0}}},
		{"", args{[]int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedArrayToBST(tt.args.nums)

			if got != nil {
				x := maxDepth(got.Left) - maxDepth(got.Right)
				if math.Abs(float64(x)) > float64(1.0) {
					t.Errorf("sortedArrayToBST() = %v", got)
				}
			}
		})
	}
}
