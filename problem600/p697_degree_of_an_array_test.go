package problem600

import "testing"

func Test_findShortestSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 3, 1},
			},
			wantAns: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{1, 2, 2, 3, 1, 4, 2},
			},
			wantAns: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findShortestSubArray(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("findShortestSubArray() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
