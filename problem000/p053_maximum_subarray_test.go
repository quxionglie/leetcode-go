package problem000

import "testing"

func Test_maxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				nums: []int{5, 4, -1, 7, 8},
			},
			want: 23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubArray(tt.args.nums); got != tt.want {
				t.Errorf("maxSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
