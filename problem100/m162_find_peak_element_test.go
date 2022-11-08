package problem100

import "testing"

func Test_findPeakElement(t *testing.T) {
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
			args: args{
				nums: []int{1, 2, 3, 1},
			},
			want: []int{2},
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			want: []int{1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.args.nums)
			var ok bool
			for _, cur := range tt.want {
				if got == cur {
					ok = true
				}
			}
			if !ok {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}

		})
	}
}
