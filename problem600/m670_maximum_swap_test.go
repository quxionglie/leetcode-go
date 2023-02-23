package problem600

import "testing"

func Test_maximumSwap(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				num: 2736,
			},
			want: 7236,
		},
		{
			name: "",
			args: args{
				num: 9973,
			},
			want: 9973,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSwap(tt.args.num); got != tt.want {
				t.Errorf("maximumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
