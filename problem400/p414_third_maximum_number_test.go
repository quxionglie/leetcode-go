package problem400

import "testing"

func Test_thirdMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{[]int{3, 2, 1}},
			want: 1,
		},
		{
			name: "",
			args: args{[]int{2, 1}},
			want: 2,
		},
		{
			name: "",
			args: args{[]int{2, 2, 3, 1}},
			want: 1,
		},
		{
			name: "",
			args: args{[]int{1, 2, 2, 5, 3, 5}},
			want: 2,
		},
		{
			name: "",
			args: args{[]int{1, 2, -2147483648}},
			want: -2147483648,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := thirdMax(tt.args.nums); got != tt.want {
				t.Errorf("thirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
