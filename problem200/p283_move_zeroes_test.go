package problem200

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{[]int{0, 1, 0, 3, 12}},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			name: "",
			args: args{[]int{0}},
			want: []int{0},
		},
		{
			name: "",
			args: args{[]int{0, 1}},
			want: []int{1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			for i := 0; i < len(tt.want); i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("missingNumber() = %v, want %v", tt.args.nums, tt.want)
				}
			}

		})
	}
}
