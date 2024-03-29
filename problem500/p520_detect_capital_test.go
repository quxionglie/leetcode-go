package problem500

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "",
			args: args{"USA"},
			want: true,
		},
		{
			name: "",
			args: args{"FlaG"},
			want: false,
		},
		{
			name: "",
			args: args{"leetcode"},
			want: true,
		},
		{
			name: "",
			args: args{"Leetcode"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCapitalUse(tt.args.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
