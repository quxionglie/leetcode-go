package problem600

import "testing"

func Test_maximumProduct(t *testing.T) {
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
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: 24,
		},
		{
			name: "",
			args: args{
				nums: []int{-10, -2, 0, 1, 2},
			},
			want: 40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
