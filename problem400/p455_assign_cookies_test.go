package problem400

import "testing"

func Test_findContentChildren(t *testing.T) {
	type args struct {
		g []int
		s []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "",
			args: args{
				g: []int{1, 2, 3},
				s: []int{1, 1},
			},
			wantAns: 1,
		},
		{
			name: "",
			args: args{
				g: []int{1, 2},
				s: []int{1, 2, 3},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findContentChildren(tt.args.g, tt.args.s); gotAns != tt.wantAns {
				t.Errorf("findContentChildren() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
